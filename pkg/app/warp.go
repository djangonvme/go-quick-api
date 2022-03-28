package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-quick-api/erron"
)

type ApiHandlerFunc func(c *gin.Context) (data interface{}, err error)

type RegisterRouteFunc func(engine *Engine)

func NewGin(registerRoutes RegisterRouteFunc) *Engine {
	// 注册路由
	if IsEnvLocal() || IsEnvDev() {
		gin.SetMode(gin.DebugMode)
	}
	eng := &Engine{gin.Default()}
	registerRoutes(eng)
	return eng
}

type routeGroup struct {
	rg *gin.RouterGroup
}

func (r *routeGroup) GinRouterGroup() *gin.RouterGroup {
	return r.rg
}

func (r *routeGroup) GET(relativePath string, handler ApiHandlerFunc) {
	r.rg.GET(relativePath, WarpApi(handler))
}

func (r *routeGroup) POST(relativePath string, handler ApiHandlerFunc) {
	r.rg.POST(relativePath, WarpApi(handler))
}

func (r *routeGroup) DELETE(relativePath string, handler ApiHandlerFunc) {
	r.rg.DELETE(relativePath, WarpApi(handler))
}

func (r *routeGroup) Any(relativePath string, handler ApiHandlerFunc) {
	r.rg.Any(relativePath, WarpApi(handler))
}

func (r *routeGroup) PATCH(relativePath string, handler ApiHandlerFunc) {
	r.rg.PATCH(relativePath, WarpApi(handler))
}

func (r *routeGroup) OPTIONS(relativePath string, handler ApiHandlerFunc) {
	r.rg.OPTIONS(relativePath, WarpApi(handler))
}

func (r *routeGroup) PUT(relativePath string, handler ApiHandlerFunc) {
	r.rg.PUT(relativePath, WarpApi(handler))
}

func (r *routeGroup) HEAD(relativePath string, handler ApiHandlerFunc) {
	r.rg.HEAD(relativePath, WarpApi(handler))
}

// 需要用其他的再加

//
type Engine struct {
	engine *gin.Engine
}

func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *routeGroup {
	for i, h := range handlers {
		handlers[i] = WarpMiddleware(h)
	}
	group := e.engine.Group(relativePath, handlers...)
	return &routeGroup{rg: group}
}

func (e *Engine) Run() error {
	return e.engine.Run(Cfg().Server.Addr)
}

// 需要用其他的再加

func WarpApi(handler ApiHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				defer func() {
					if msg = recover(); msg != nil {
						c.AbortWithStatusJSON(http.StatusOK, response{Code: erron.ErrInternal, Msg: fmt.Sprintf("api request panic: %v", msg), Timestamp: time.Now().Unix(), Data: nil})
					}
				}()
				err := erron.Inner(fmt.Sprintf("%v", msg))
				LogApiPanic(c, msg)
				AbortJSON(c, ResponseFail(err))
			}
		}()
		data, err := handler(c)
		var errInfo erron.E
		if err != nil {
			c.Error(err)
			if v, ok := err.(erron.E); ok {
				errInfo = v
			} else {
				errInfo = erron.New(erron.Failed, err.Error())
			}
		} else {
			errInfo = erron.New(erron.Success)
		}
		OutputJSON(c, Response(errInfo, data))
	}
}

func WarpMiddleware(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				defer func() {
					if msg = recover(); msg != nil {
						c.AbortWithStatusJSON(http.StatusOK, response{Code: erron.ErrInternal, Msg: fmt.Sprintf("middleware panic: %v", msg), Timestamp: time.Now().Unix(), Data: nil})
					}
				}()
				err := erron.Inner(fmt.Sprintf("%v", msg))
				LogApiPanic(c, msg)
				AbortJSON(c, ResponseFail(err))
			}
		}()
		if _, ok := c.Get(CtxStartTime); !ok {
			c.Set(CtxStartTime, time.Now())
		}
		handler(c)
	}
}
