package middleware

import (
	"management-system-server/core"

	"github.com/gin-gonic/gin"
)

func Return(c *core.Context) {
	c.Next()
	c.JSON(c.GetCode(), gin.H{
		"msg":  "success",
		"data": c.GetData(),
	})
}
