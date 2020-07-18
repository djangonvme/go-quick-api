package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jangozw/go-quick-api/erron"
)

// api 处理函数类型
type ApiHandlerFunc func(c *gin.Context) (data interface{}, err error)

// 注册路由函数类型
type RegisterRouteFunc func(engine *Engine)

// 定义路由 gin.Engine 统一加warp
func NewGin(registerRoutes RegisterRouteFunc) *Engine {
	// 注册路由
	if IsEnvLocal() || IsEnvDev() {
		gin.SetMode(gin.DebugMode)
	}
	eng := &Engine{gin.New()}
	registerRoutes(eng)
	return eng
}

type routeGroup struct {
	rg *gin.RouterGroup
}

// 路由组中间件
/*
// use 没必要
func (r *routeGroup) Use(middleware ...gin.HandlerFunc) *routeGroup {
	for i, h := range middleware {
		middleware[i] = WarpMiddleware(h)
	}
	r.rg.Use(middleware...)
	return r
}*/

// 返回gin原生的routerGroup
func (r *routeGroup) GinRouterGroup() *gin.RouterGroup {
	return r.rg
}

// 需要什么方法自由搬运 gin.routeGroup
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
	return e.engine.Run(HttpServeAddr())
}

// 需要用其他的再加

// api 捕获异常
func WarpApi(handler ApiHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				defer func() {
					if msg := recover(); msg != nil {
						c.AbortWithStatusJSON(http.StatusOK, response{Code: erron.ErrInternal, Msg: "bad request", Timestamp: time.Now().Unix(), Data: nil})
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

// 中间件 捕获异常
func WarpMiddleware(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				defer func() {
					if msg := recover(); msg != nil {
						c.AbortWithStatusJSON(http.StatusOK, response{Code: erron.ErrInternal, Msg: "bad request", Timestamp: time.Now().Unix(), Data: nil})
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


