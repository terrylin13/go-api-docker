package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/terrylin13/go-api-docker/pkg/router/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	setUp(r)
	return r
}

func setUp(e *gin.Engine) {
	r := e.Group("")
	v1.RegisterRouter(r)
}
