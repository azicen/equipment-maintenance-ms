package middleware

import (
	"management-system-server/model"
	"management-system-server/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetEmployee(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

//AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if employee, _ := c.Get("user"); employee != nil {
			if _, ok := employee.(*model.Employee); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
