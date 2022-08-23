package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/terrylin13/go-api-docker/pkg/handler/v1/index"
	"github.com/terrylin13/go-api-docker/pkg/handler/v1/ws"
)

func RegisterRouter(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", index.Handle)
		v1.GET("/ws", ws.Handle)
	}
}
