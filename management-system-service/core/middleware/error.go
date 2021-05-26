package middleware

import (
	"management-system-server/core"
	log "management-system-server/util/logger"
)

func Error(c *core.Context) {
	c.Next()
	log.Debug("func Error(c *core.Context)")
	if c.IsError() {
		for _, err := range c.Errors {
			log.Warning(err.Error())
		}
	}
}
