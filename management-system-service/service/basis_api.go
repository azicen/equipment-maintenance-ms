package service

import (
	"management-system-server/core"

	"github.com/gin-gonic/gin"
)

func (s *Service) Ping(c *core.Context) {
	c.JSON(200, gin.H{
		"msg": "Ping",
	})
}
