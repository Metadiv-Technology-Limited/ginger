package ginger

import (
	"ginger/util"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Metadiv-Technology-Limited/nanoid"
	"github.com/Metadiv-Technology-Limited/sql/models"

	"github.com/gin-gonic/gin"
)

type IContext[T any] interface {
	GetEngine() IEngine
	SetEngine(IEngine)
	GetGinCtx() *gin.Context
	SetGinCtx(*gin.Context)
	GetPage() models.IPagination
	SetPage(models.IPagination)
	GetSort() models.ISorting
	SetSort(models.ISorting)
	GetRequest() *T
	SetRequest(*T)
	GetResponse() IResponse
	SetResponse(IResponse)
	GetStartAt() time.Time
	SetStartAt(time.Time)
	GetHasResp() bool
	SetHasResp(bool)
	GetIsFile() bool
	SetIsFile(bool)
}

func NewContext[T any](engine IEngine, ginCtx *gin.Context) IContext[T] {
	var page models.IPagination
	var sort models.ISorting
	if ginCtx.Request.Method == "GET" {
		page = new(models.Pagination)
		sort = new(models.Sorting)
		ginCtx.ShouldBindQuery(page)
		ginCtx.ShouldBindQuery(sort)
	}
	return &Context[T]{
		Engine:  engine,
		GinCtx:  ginCtx,
		Page:    page,
		Sort:    sort,
		Request: util.GinRequest[T](ginCtx),
		StartAt: time.Now(),
	}
}

type Context[T any] struct {
	Engine   IEngine
	GinCtx   *gin.Context
	Page     models.IPagination
	Sort     models.ISorting
	Request  *T
	Response IResponse
	StartAt  time.Time
	HasResp  bool
	IsFile   bool
}

func (c *Context[T]) GetEngine() IEngine {
	return c.Engine
}

func (c *Context[T]) SetEngine(engine IEngine) {
	c.Engine = engine
}

func (c *Context[T]) GetGinCtx() *gin.Context {
	return c.GinCtx
}

func (c *Context[T]) SetGinCtx(ginCtx *gin.Context) {
	c.GinCtx = ginCtx
}

func (c *Context[T]) GetPage() models.IPagination {
	return c.Page
}

func (c *Context[T]) SetPage(page models.IPagination) {
	c.Page = page
}

func (c *Context[T]) GetSort() models.ISorting {
	return c.Sort
}

func (c *Context[T]) SetSort(sort models.ISorting) {
	c.Sort = sort
}

func (c *Context[T]) GetRequest() *T {
	return c.Request
}

func (c *Context[T]) SetRequest(request *T) {
	c.Request = request
}

func (c *Context[T]) GetResponse() IResponse {
	return c.Response
}

func (c *Context[T]) SetResponse(response IResponse) {
	c.Response = response
}

func (c *Context[T]) GetStartAt() time.Time {
	return c.StartAt
}

func (c *Context[T]) SetStartAt(startAt time.Time) {
	c.StartAt = startAt
}

func (c *Context[T]) GetHasResp() bool {
	return c.HasResp
}

func (c *Context[T]) SetHasResp(hasResp bool) {
	c.HasResp = hasResp
}

func (c *Context[T]) GetIsFile() bool {
	return c.IsFile
}

func (c *Context[T]) SetIsFile(isFile bool) {
	c.IsFile = isFile
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.GinCtx.ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.GinCtx.Request.UserAgent()
}

func (ctx *Context[T]) BearerToken() string {
	token := ctx.GinCtx.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")
	token = strings.ReplaceAll(token, "BEARER ", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}

func (ctx *Context[T]) OK(data any, page ...models.IPagination) {
	if ctx.GetHasResp() {
		log.Println("Warning: context already responded")
		return
	}
	// pagination
	var pageResponse models.IPagination
	if len(page) > 0 {
		pageResponse = page[0]
	}
	ctx.SetResponse(&Response{
		Success:    true,
		Pagination: pageResponse,
		Data:       data,
		Duration:   time.Since(ctx.GetStartAt()).Milliseconds(),
	})
	ctx.SetHasResp(true)
}

func (ctx *Context[T]) OKFile(bytes []byte, filename ...string) {
	if ctx.GetHasResp() {
		log.Println("Warning: context already responded")
		return
	}

	var name string
	if len(filename) == 0 || filename[0] == "" {
		name = nanoid.NewSafe()
	} else {
		name = filename[0]
	}

	ctx.GinCtx.Header("Content-Disposition", "filename="+name)
	ctx.GinCtx.Data(http.StatusOK, util.DetermineFileType(name), bytes)
	ctx.SetHasResp(true)
	ctx.SetIsFile(true)
}

func (ctx *Context[T]) OKDownload(bytes []byte, filename ...string) {
	if ctx.GetHasResp() {
		log.Println("Warning: context already responded")
		return
	}

	var name string
	if len(filename) == 0 || filename[0] == "" {
		name = nanoid.NewSafe()
	} else {
		name = filename[0]
	}

	ctx.GinCtx.Header("Content-Disposition", "filename="+name)
	ctx.GinCtx.Data(http.StatusOK, "application/octet-stream", bytes)
	ctx.SetHasResp(true)
	ctx.SetIsFile(true)
}

func (ctx *Context[T]) Err(code string, locale ...string) {
	if ctx.GetHasResp() {
		log.Println("Warning: context already responded")
		return
	}
	ctx.SetResponse(&Response{
		Success:  false,
		Duration: time.Since(ctx.GetStartAt()).Milliseconds(),
		Error:    errMapObj.GetError(code, locale...),
	})
	ctx.SetHasResp(true)
}
