package app

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/task-dispatcher/erron"

	"github.com/gin-gonic/gin"
	"gitlab.com/task-dispatcher/pkg/auth"
	"gitlab.com/task-dispatcher/types"
)

const (
	CtxKeyResponse = "ctx-response"
	CtxStartTime   = "ctx-start-time"
	CtxRequestBody = "ctx-request-body"
)

// LoginUser 登陆用户信息, 可根据需要扩充字段
type LoginUser struct {
	ID int64
}

type TokenPayload struct {
	UserID int64 `json:"uid"`
}

// ParseUserByToken 根据token解析登陆用户
func ParseUserByToken(token string) (TokenPayload, error) {
	user := TokenPayload{}
	if token == "" {
		return user, errors.New("empty token")
	}
	if jwtPayload, err := auth.ParseJwtToken(token, Cfg().General.JwtSecret); err != nil {
		return user, err
	} else if err := jwtPayload.ParseUser(&user); err != nil {
		return user, err
	}
	if user.UserID == 0 {
		return user, errors.New("invalid login user")
	}
	return user, nil
}

/*// SetCtxRequestBody gin自带的bind系列方法只能用一次，所以此处将c.Request.Body存起来,就可以实现一次请求多次验证了
func SetCtxRequestBody(c *gin.Context) {
	bodyStr := c.GetString(CtxRequestBody)
	if bodyStr == "" {
		body, _ := c.GetRawData()
		bodyStr = string(body)
		c.Set(CtxRequestBody, bodyStr)
	}
	if bodyStr != "" {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(bodyStr)))
	}
}

// CheckRequestBody 解决gin 的一个坑吧, 需要用到xxxBind()地方之前确保c.request.body 存在
func CheckRequestBody(c *gin.Context) {
	if str := c.GetString(CtxRequestBody); str != "" {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(str)))
	}
}

func ShouldBind(c *gin.Context, input interface{}) error {
	CheckRequestBody(c)
	if err := c.ShouldBind(input); err != nil {
		return err
	}
	return nil
}*/

// GetSaveRawData save the request body
func GetSaveRawData(c *gin.Context) ([]byte, error) {
	bodyStr := c.GetString(CtxRequestBody)
	if bodyStr != "" {
		// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(bodyStr)))
		return []byte(bodyStr), nil
	}
	body, err := c.GetRawData()
	if err != nil {
		if LoggerInstance != nil {
			LoggerInstance.Errorf("GetRawData failed: %v\n", err.Error())
		} else {
			log.Printf("GetRawData failed: %v\n", err.Error())
		}
		return nil, err
	}
	c.Set(CtxRequestBody, string(body))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return []byte(c.GetString(CtxRequestBody)), nil
}

func GetLoginUser(c *gin.Context) (user LoginUser, err error) {
	info, err := ParseUserByToken(c.GetHeader(types.TokenHeaderKey))
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

// BindInput 绑定输入参数
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

func GetPager(c *gin.Context) Pager {
	pager := Pager{}
	c.ShouldBind(&pager)
	pager.Secure()
	return pager
}

func setResponse(c *gin.Context, resp *response) {
	c.Set(CtxKeyResponse, resp)
}

func OutputJSON(c *gin.Context, resp *response) {
	setResponse(c, resp)
	c.JSON(http.StatusOK, resp)
}

func AbortJSON(c *gin.Context, resp *response) {
	setResponse(c, resp)
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func checkInput(input interface{}) error {
	inputTypeErr := errors.New("input must a struct")
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

// LogApiPanic api 请求发生了panic 记入日志
func LogApiPanic(c *gin.Context, panicMsg interface{}) {

	if LoggerInstance == nil {
		return
	}
	user, _ := GetLoginUser(c)
	start := c.GetTime(CtxStartTime)
	// 执行时间
	resp, ok := c.Get(CtxKeyResponse)
	if !ok {
		resp = struct{}{}
	}
	body, _ := GetSaveRawData(c)
	var errMsg string
	if len(c.Errors) > 0 {
		for i, es := range c.Errors {
			errMsg += fmt.Sprintf(" #%d: %+v", i, es)
		}
	}

	// log 里有json.Marshal() 导致url转义字符
	LoggerInstance.WithFields(logrus.Fields{
		"uid":      user.ID,
		"body":     body,
		"query":    c.Request.URL.Query(),
		"response": resp,
		"uri":      c.Request.URL.RequestURI(),
		"cost":     fmt.Sprintf("%s", time.Since(start)),
		"ip":       c.ClientIP(),
		"header":   c.Request.Header,
		"method":   c.Request.Method,
		"errors":   errMsg,
		"type":     "api_panic",
	}).Infof("%s | %s %s", panicMsg, c.Request.Method, c.Request.URL.RequestURI())
}

// ApiLog api 接口日志记录请求和返回
func ApiLog(c *gin.Context) {

	if LoggerInstance == nil {
		return
	}
	user, _ := GetLoginUser(c)
	start := c.GetTime(CtxStartTime)
	// 执行时间
	resp, ok := c.Get(CtxKeyResponse)
	if !ok {
		resp = struct{}{}
	}
	body, _ := GetSaveRawData(c)

	var errMsg string
	if len(c.Errors) > 0 {
		for i, es := range c.Errors {
			errMsg += fmt.Sprintf(" #%d: %+v", i, es)
		}
	}

	// log 里有json.Marshal() 导致url转义字符
	LoggerInstance.WithFields(logrus.Fields{
		"uid":      user.ID,
		"body":     string(body),
		"query":    c.Request.URL.Query(),
		"response": resp,
		"type":     "api_request",
		"header":   c.Request.Header,
		"method":   c.Request.Method,
		"uri":      c.Request.URL.RequestURI(),
		"cost":     fmt.Sprintf("%s", time.Since(start)),
		"errors":   errMsg,
		"ip":       c.ClientIP(),
	}).Infof(" %s  %s | %3v", c.Request.Method, c.Request.URL.RequestURI(), time.Since(start))
}
