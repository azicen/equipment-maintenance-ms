package service

import (
	"management-system-server/core"
	"management-system-server/serializer"
)

func (s *Service) Ping(c *core.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Ping",
	})
}
