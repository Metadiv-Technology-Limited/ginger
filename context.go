package ginger

import (
	"time"

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
