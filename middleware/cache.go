package middleware

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

/*
Cache setup gin cache middleware
*/
func Cache(duration time.Duration, handler gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler)
}
