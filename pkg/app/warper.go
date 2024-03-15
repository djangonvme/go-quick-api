package app

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/qubic-pool/config"
	"gitlab.com/qubic-pool/pkg/logger"
	"io"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/qubic-pool/erron"
)

// GetRequestBody 修复 gin 获取body只能获取一次，调用GetRequestBody可以无限次获取
func GetRequestBody(ctx *gin.Context) []byte {
	data, _ := ctx.GetRawData()
	if data != nil {
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 关键点
		ctx.Set("requestBody", string(data))
	}
	return data
}

type ApiHandlerFunc func(c *gin.Context) (data interface{}, err error)

type RegisterRouteFunc func(engine *Engine)

func NewGin(registerRoutes RegisterRouteFunc) *Engine {
	// 注册路由
	if !config.IsEnvProduction() {
		gin.SetMode(gin.DebugMode)
	}
	eng := &Engine{gin.Default()}
	registerRoutes(eng)
	return eng
}

type RouteGroup struct {
	rg *gin.RouterGroup
}

func (r *RouteGroup) GinRouterGroup() *gin.RouterGroup {
	return r.rg
}

func (r *RouteGroup) GET(relativePath string, handler ApiHandlerFunc) {
	r.rg.GET(relativePath, apiWrapper(handler))
}

func (r *RouteGroup) POST(relativePath string, handler ApiHandlerFunc) {
	r.rg.POST(relativePath, apiWrapper(handler))
}

func (r *RouteGroup) DELETE(relativePath string, handler ApiHandlerFunc) {
	r.rg.DELETE(relativePath, apiWrapper(handler))
}

func (r *RouteGroup) Any(relativePath string, handler ApiHandlerFunc) {
	r.rg.Any(relativePath, apiWrapper(handler))
}

func (r *RouteGroup) PATCH(relativePath string, handler ApiHandlerFunc) {
	r.rg.PATCH(relativePath, apiWrapper(handler))
}

func (r *RouteGroup) OPTIONS(relativePath string, handler ApiHandlerFunc) {
	r.rg.OPTIONS(relativePath, apiWrapper(handler))
}

func (r *RouteGroup) PUT(relativePath string, handler ApiHandlerFunc) {
	r.rg.PUT(relativePath, apiWrapper(handler))
}

func (r *RouteGroup) HEAD(relativePath string, handler ApiHandlerFunc) {
	r.rg.HEAD(relativePath, apiWrapper(handler))
}

// 需要用其他的再加

type Engine struct {
	engine *gin.Engine
}

func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouteGroup {
	for i, h := range handlers {
		handlers[i] = middlewareWrapper(h)
	}
	group := e.engine.Group(relativePath, handlers...)
	return &RouteGroup{rg: group}
}

func (e *Engine) Run(listen string) error {
	return e.engine.Run(listen)
}

// 需要用其他的再加

func logApiResponse(c *gin.Context, response *Response) {
	if config.IsEnvDebug() {
		logger.Instance.Debugf("ApiRequest %s %s Body: %s Response: %s", c.Request.Method, c.Request.URL, c.GetString("requestBody"), response.ToJson())
	}
}

func apiWrapper(handler ApiHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				logger.Instance.Errorf("panic info: %s", string(debug.Stack()))
				logApiPanic(c, msg)
				resp := Response{Code: erron.PanicCode, Msg: fmt.Sprintf("%v", msg), Timestamp: time.Now().Unix(), Data: nil}
				logApiResponse(c, &resp)
				c.AbortWithStatusJSON(http.StatusOK, resp)
				return
			}
		}()
		result, err := handler(c)
		var errInfo erron.ErrorIF
		if err != nil {
			if v, ok := err.(erron.ErrorIF); ok {
				errInfo = v
			} else {
				errInfo = erron.New(erron.Failed, err.Error())
			}
		} else {
			errInfo = erron.New(erron.Success, "")
		}
		resp := ResponseApi(errInfo, result)
		logApiResponse(c, resp)
		c.JSON(http.StatusOK, resp)
		return
	}
}

func middlewareWrapper(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				logApiPanic(c, msg)
				c.AbortWithStatusJSON(http.StatusOK, Response{Code: erron.PanicCode, Msg: fmt.Sprintf("%v", msg), Timestamp: time.Now().Unix(), Data: nil})
				return
			}
		}()
		handler(c)
	}
}

func logApiPanic(c *gin.Context, panicMsg interface{}) {
	var errMsg string
	if len(c.Errors) > 0 {
		for i, es := range c.Errors {
			errMsg += fmt.Sprintf(" #%d: %+v", i, es)
		}
	}
	// log 里有json.Marshal() 导致url转义字符
	logger.Instance.WithFields(logrus.Fields{
		"query":  c.Request.URL.Query(),
		"uri":    c.Request.URL.RequestURI(),
		"ip":     c.ClientIP(),
		"header": c.Request.Header,
		"method": c.Request.Method,
		"errors": errMsg,
	}).Infof("[api_panic] %s | %s %s", panicMsg, c.Request.Method, c.Request.URL.RequestURI())
}
