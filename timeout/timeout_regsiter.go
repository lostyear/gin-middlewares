package timeout

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultTimeoutMsg = "timeout"

var msg string

func init() {
	msg = defaultTimeoutMsg
}

func ChangeTimeoutMsg(s string) {
	msg = s
}

type ITimeoutRoutes interface {
	Use(...gin.HandlerFunc) ITimeoutRouter

	Handle(string, string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	Any(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	GET(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	POST(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	PUT(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	DELETE(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	PATCH(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	OPTIONS(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter
	HEAD(string, time.Duration, ...gin.HandlerFunc) ITimeoutRouter

	StaticFile(string, string) ITimeoutRouter
	Static(string, string) ITimeoutRouter
	StaticFS(string, http.FileSystem) ITimeoutRouter
}
type ITimeoutRouter interface {
	ITimeoutRoutes
	Group(string, ...gin.HandlerFunc) ITimeoutRouter
}

type TimeoutRegsiter struct {
	ITimeoutRouter
	gin.IRouter
}

func NewRegsiter(r gin.IRouter) ITimeoutRouter {
	return &TimeoutRegsiter{
		IRouter: r,
	}
}

func (tr *TimeoutRegsiter) Use(handlers ...gin.HandlerFunc) ITimeoutRouter {
	tr.IRouter.Use(handlers...)
	return tr
}

func (tr *TimeoutRegsiter) Group(relativePath string, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return NewRegsiter(tr.IRouter.Group(relativePath, handlers...))
}

func (tr *TimeoutRegsiter) Handle(httpMethod, relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	hs := make([]gin.HandlerFunc, 1, len(handlers)+1)
	hs[0] = TimeoutMiddleware(timeout, msg)
	hs = append(hs, handlers...)
	tr.IRouter.Handle(httpMethod, relativePath, hs...)
	return tr
}

func (tr *TimeoutRegsiter) GET(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return tr.Handle(http.MethodGet, relativePath, timeout, handlers...)
}

func (tr *TimeoutRegsiter) POST(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return tr.Handle(http.MethodPost, relativePath, timeout, handlers...)
}

func (tr *TimeoutRegsiter) PUT(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return tr.Handle(http.MethodPut, relativePath, timeout, handlers...)
}

func (tr *TimeoutRegsiter) DELETE(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return tr.Handle(http.MethodDelete, relativePath, timeout, handlers...)
}

func (tr *TimeoutRegsiter) PATCH(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return tr.Handle(http.MethodPatch, relativePath, timeout, handlers...)
}

func (tr *TimeoutRegsiter) OPTIONS(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return tr.Handle(http.MethodOptions, relativePath, timeout, handlers...)
}

func (tr *TimeoutRegsiter) HEAD(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	return tr.Handle(http.MethodHead, relativePath, timeout, handlers...)
}

func (tr *TimeoutRegsiter) Any(relativePath string, timeout time.Duration, handlers ...gin.HandlerFunc) ITimeoutRouter {
	tr.Handle(http.MethodGet, relativePath, timeout, handlers...)
	tr.Handle(http.MethodPost, relativePath, timeout, handlers...)
	tr.Handle(http.MethodPut, relativePath, timeout, handlers...)
	tr.Handle(http.MethodPatch, relativePath, timeout, handlers...)
	tr.Handle(http.MethodHead, relativePath, timeout, handlers...)
	tr.Handle(http.MethodOptions, relativePath, timeout, handlers...)
	tr.Handle(http.MethodDelete, relativePath, timeout, handlers...)
	tr.Handle(http.MethodConnect, relativePath, timeout, handlers...)
	tr.Handle(http.MethodTrace, relativePath, timeout, handlers...)
	return tr
}

func (tr *TimeoutRegsiter) StaticFile(relativePath, filepath string) ITimeoutRouter {
	tr.IRouter.StaticFile(relativePath, filepath)
	return tr
}

func (tr *TimeoutRegsiter) Static(relativePath, root string) ITimeoutRouter {
	tr.IRouter.Static(relativePath, root)
	return tr
}

func (tr *TimeoutRegsiter) StaticFS(relativePath string, fs http.FileSystem) ITimeoutRouter {
	tr.IRouter.StaticFS(relativePath, fs)
	return tr
}
