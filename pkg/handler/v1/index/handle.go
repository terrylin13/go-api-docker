package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terrylin13/go-api-docker/pkg/handler/v1/resp"
)

func Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, resp.Resp{Msg: "Success!"})
}
