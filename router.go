package ginger

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func CRON(engine IEngine, spec string, job func()) {
	engine.GetCron().AddFunc(spec, job)
}

func GET[T any](engine IEngine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().GET(route, append(middleware, ApiServiceToGinHandler[T](engine, handler))...)
}

func GETCached[T any](engine IEngine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().GET(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, ApiServiceToGinHandler[T](engine, handler)))...)
}

func POST[T any](engine IEngine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().POST(route, append(middleware, ApiServiceToGinHandler[T](engine, handler))...)
}

func POSTCached[T any](engine IEngine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().POST(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, ApiServiceToGinHandler[T](engine, handler)))...)
}

func PUT[T any](engine IEngine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().PUT(route, append(middleware, ApiServiceToGinHandler[T](engine, handler))...)
}

func PUTCached[T any](engine IEngine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().PUT(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, ApiServiceToGinHandler[T](engine, handler)))...)
}

func DELETE[T any](engine IEngine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().DELETE(route, append(middleware, ApiServiceToGinHandler[T](engine, handler))...)
}

func DELETECached[T any](engine IEngine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().DELETE(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, ApiServiceToGinHandler[T](engine, handler)))...)
}

func WS[T any](engine IEngine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().GET(route, append(middleware, WsServiceToGinHandler[T](engine, handler))...)
}
