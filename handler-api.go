package ginger

import "github.com/gin-gonic/gin"

type Service[T any] func(ctx IContext[T])

type Handler[T any] Service[T]

func ApiServiceToGinHandler[T any](engine IEngine, service func(ctx IContext[T])) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := NewContext[T](engine, ctx)
		service(c)

		// if file is served, no need to respond
		if c.GetHasResp() && c.GetIsFile() {
			return
		}

		// unexpected, service did not respond
		if !c.GetHasResp() || c.GetResponse() == nil {
			ctx.JSON(500, gin.H{
				"message": "service did not respond",
			})
			return
		}

		// success
		if c.GetResponse().GetSuccess() {
			ctx.JSON(200, c.GetResponse())
			return
		}

		// error, but no error object
		if c.GetResponse().GetError() == nil {
			ctx.JSON(500, gin.H{
				"message": "service did not respond with error",
			})
			return
		}

		// error
		ctx.JSON(400, c.GetResponse())
	}
}
