package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jangozw/go-quick-api/erron"

	"github.com/gin-gonic/gin"
	"github.com/jangozw/go-quick-api/param"
	"github.com/jangozw/go-quick-api/pkg/auth"
)

const (
	CtxKeyResponse = "ctx-response"
)

// 登陆用户信息, 可根据业务需要扩充字段
type LoginUser struct {
	ID int64
}

type TokenPayload struct {
	UserID int64 `json:"uid"`
}

// 根据token解析登陆用户
func ParseUserByToken(token string) (TokenPayload, error) {
	user := TokenPayload{}
	if token == "" {
		return user, errors.New("empty token")
	}
	if jwtPayload, err := auth.ParseJwtToken(token, Cfg.General.JwtSecret); err != nil {
		return user, err
	} else if err := jwtPayload.ParseUser(&user); err != nil {
		return user, err
	}
	if user.UserID == 0 {
		return user, errors.New("invalid login user")
	}
	return user, nil
}

func GetLoginUser(c *gin.Context) (user LoginUser, err error) {
	info, err := ParseUserByToken(c.GetHeader(param.TokenHeaderKey))
	if err != nil {
		return user, err
	}
	return LoginUser{ID: info.UserID}, nil
}

func MustGetLoginUser(c *gin.Context) LoginUser {
	user, err := GetLoginUser(c)
	if err != nil || user.ID == 0 {
		panic(err)
	}
	return user
}

// 绑定输入参数
func BindInput(c *gin.Context, input interface{}) erron.E {
	if err := checkInput(input); err == nil {
		if err := c.ShouldBind(input); err != nil {
			return erron.New(erron.ErrRequestParam, err.Error())
		}
		// 如果实现了 params 接口就验证参数
		if obj, ok := input.(CheckIF); ok {
			if err := obj.Check(); err != nil {
				return erron.New(erron.ErrRequestParam, err.Error())
			}
		}
	}
	return nil
}

// 展示的分页
func GetPager(c *gin.Context) *Pager {
	pager := &Pager{}
	c.ShouldBind(pager)
	pager.Secure()
	return pager
}

func setResponse(c *gin.Context, resp *response) {
	c.Set(CtxKeyResponse, resp)
}

// 正常输出
func OutputJSON(c *gin.Context, resp *response) {
	setResponse(c, resp)
	c.JSON(http.StatusOK, resp)
}

// 中断并输出
func AbortJSON(c *gin.Context, resp *response) {
	setResponse(c, resp)
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func checkInput(input interface{}) error {
	inputTypeErr := errors.New("input 必须是一个结构体变量的地址")
	rv := reflect.ValueOf(input)
	if rv.Kind() != reflect.Ptr || rv.IsNil() || !rv.IsValid() {
		return inputTypeErr
	}
	rt := reflect.TypeOf(input)
	if rt.Kind() == reflect.Ptr {
		if rt.Elem().Kind() != reflect.Struct {
			return inputTypeErr
		}
	} else {
		if rt.Kind() != reflect.Struct {
			return inputTypeErr
		}
	}
	return nil
}

// api 请求发生了panic 记入日志
func LogApiPanic(c *gin.Context, panicMsg interface{}) {
	user, _ := GetLoginUser(c)
	start := c.GetTime(CtxStartTime)
	// 执行时间
	latency := time.Since(start)
	resp, ok := c.Get(CtxKeyResponse)
	if !ok {
		resp = struct{}{}
	}
	var query interface{}
	if c.Request.Method == "GET" {
		query = c.Request.URL.Query()
	} else {
		postData, _ := c.GetRawData()
		query := make(map[string]interface{})
		json.Unmarshal(postData, &query)
	}

	// log 里有json.Marshal() 导致url转义字符
	Logger.WithFields(logrus.Fields{
		"uid":      user.ID,
		"query":    query,
		"response": resp,
		"method":   c.Request.Method,
		"uri":      c.Request.URL.RequestURI(),
		"latency":  fmt.Sprintf("%3v", latency),
		"ip":       c.ClientIP(),
	}).Infof("--panic: %s | %s %s", panicMsg, c.Request.Method, c.Request.URL.RequestURI())
}
