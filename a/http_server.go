package a

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Ctx interface {
	Ok(obj ...any) Ctx
	Status(status int) Ctx
	JsonBody(obj any)
	GetQuery(key string) string
	GetQueryArray(key string) []string
	PathVar(key string) string
	JsonArgs(obj any)
	QueryArgs(obj any)
	Headers(obj any)
	Header(key, value string)
	GetHeader(key string) string
	HeaderMap() http.Header
	BindArgs(obj any)
	Redirect(location string)
}

type ctxImpl struct {
	raw    *gin.Context
	status int
}

func (receiver *ctxImpl) Ok(obj ...any) Ctx {
	if len(obj) > 0 {
		receiver.raw.JSON(200, obj[0])
		return nil
	}
	receiver.status = 200
	return receiver
}

func (receiver *ctxImpl) Status(status int) Ctx {
	receiver.status = status
	return receiver
}
func (receiver *ctxImpl) GetQuery(key string) string {
	return receiver.raw.Query(key)
}
func (receiver *ctxImpl) GetQueryArray(key string) []string {
	values, _ := receiver.raw.GetQueryArray(key)
	return values
}
func (receiver *ctxImpl) JsonArgs(obj any) {
	err := receiver.raw.ShouldBindJSON(obj)
	if err != nil {
		panic(err)
	}
}
func (receiver *ctxImpl) QueryArgs(obj any) {
	err := receiver.raw.ShouldBindQuery(obj)
	if err != nil {
		panic(err)
	}
}
func (receiver *ctxImpl) PathVar(key string) string {
	return receiver.raw.Param(key)
}
func (receiver *ctxImpl) HeaderMap() http.Header {
	return receiver.raw.Request.Header
}
func (receiver *ctxImpl) Header(key, value string) {
	receiver.raw.Header(key, value)
}
func (receiver *ctxImpl) GetHeader(key string) string {
	return receiver.raw.GetHeader(key)
}
func (receiver *ctxImpl) Headers(obj any) {
	err := receiver.raw.ShouldBindHeader(obj)
	if err != nil {
		panic(err)
	}
}
func (receiver *ctxImpl) BindArgs(obj any) {
	err := receiver.raw.ShouldBind(obj)
	if err != nil {
		panic(err)
	}
}
func (receiver *ctxImpl) JsonBody(obj any) {
	receiver.raw.JSON(receiver.status, obj)
}
func (receiver *ctxImpl) Redirect(location string) {
	receiver.raw.Redirect(http.StatusFound, location)
}

type HttpHandler func(Ctx)

type HttpServer struct {
	port int
	raw  *gin.Engine
}

func HttpServerDefault() *HttpServer {
	return &HttpServer{port: 8989, raw: gin.Default()}
}
func HttpServerNew(port int) *HttpServer {
	return &HttpServer{port: port, raw: gin.Default()}
}

func generateHandler(handler HttpHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &ctxImpl{raw: c}
		handler(ctx)
	}
}

func (receiver *HttpServer) Static(relativePath, root string) gin.IRoutes {
	return receiver.raw.Static(relativePath, root)
}
func (receiver *HttpServer) Post(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.POST(relativePath, generateHandler(handler))
}

func (receiver *HttpServer) Get(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.GET(relativePath, generateHandler(handler))
}

func (receiver *HttpServer) Delete(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.DELETE(relativePath, generateHandler(handler))
}

func (receiver *HttpServer) Patch(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.PATCH(relativePath, generateHandler(handler))
}

func (receiver *HttpServer) Put(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.PUT(relativePath, generateHandler(handler))
}

func (receiver *HttpServer) Options(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.OPTIONS(relativePath, generateHandler(handler))
}

func (receiver *HttpServer) Head(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.HEAD(relativePath, generateHandler(handler))
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (receiver *HttpServer) Any(relativePath string, handler HttpHandler) gin.IRoutes {
	return receiver.raw.Any(relativePath, generateHandler(handler))
}

func (receiver *HttpServer) Start() {
	err := receiver.raw.Run(fmt.Sprintf(":%v", receiver.port))
	if err != nil {
		log.Println(err)
	}
}
