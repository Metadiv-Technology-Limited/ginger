package mock

import (
	"net/http/httptest"

	"github.com/Metadiv-Technology-Limited/ginger"

	"github.com/Metadiv-Technology-Limited/sql/models"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type MockContextParams[T any] struct {
	Request   *T
	Page      models.IPagination
	Sort      models.ISorting
	Method    string
	Path      string
	ClientIP  string
	UserAgent string
	Headers   map[string]string
}

func MockContext[T any](params MockContextParams[T]) ginger.IContext[T] {
	w := httptest.NewRecorder()
	ctx, e := gin.CreateTestContext(w)

	if params.Method == "" {
		params.Method = "GET"
	}
	if params.Path == "" {
		params.Path = "/"
	}

	ctx.Request = httptest.NewRequest(params.Method, params.Path, nil)

	if params.ClientIP != "" {
		ctx.Request.RemoteAddr = params.ClientIP
	} else {
		ctx.Request.RemoteAddr = "127.0.0.1"
	}

	if params.UserAgent != "" {
		ctx.Request.Header.Set("User-Agent", params.UserAgent)
	} else {
		ctx.Request.Header.Set("User-Agent", "Mock Context")
	}

	if params.Headers != nil {
		for k, v := range params.Headers {
			ctx.Request.Header.Set(k, v)
		}
	}

	mockEngine := &ginger.Engine{
		Gin:  e,
		Cron: cron.New(),
	}

	mockCtx := ginger.NewContext[T](mockEngine, ctx)
	mockCtx.SetRequest(params.Request)
	mockCtx.SetPage(params.Page)
	mockCtx.SetSort(params.Sort)
	return mockCtx
}
