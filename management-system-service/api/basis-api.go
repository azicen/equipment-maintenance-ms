package api

import (
	"github.com/gin-gonic/gin"
	"management-system-server/serializer"
)

func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}
