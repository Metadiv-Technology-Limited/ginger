package ginger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WsServiceToGinHandler[T any](engine IEngine, service func(ctx IContext[T])) gin.HandlerFunc {
	wsUpGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return func(ctx *gin.Context) {
		ws, err := wsUpGrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer ws.Close()

		c := NewContext[T](engine, ctx)
		service(c)

		if c.GetResponse().GetError() != nil {
			ctx.JSON(400, c.GetResponse())
		}
	}
}
