package middleware

import (
	"management-system-server/core"
	"management-system-server/dao"
	"management-system-server/model"
	"management-system-server/serializer"

	"github.com/gin-contrib/sessions"
)

//CurrentUser 获取登录用户
func CurrentUser(c *core.Context) {
	session := sessions.Default(c.Context)
	uid := session.Get("user_id")
	if uid != nil {
		user, err := dao.GetDao().GetUser(c, uid)
		if err == nil {
			c.Set("user", &user)
		}
	}
	c.Next()
}

//AuthRequired 需要登录
func AuthRequired(c *core.Context) {
	if user, _ := c.Get("user"); user != nil {
		if _, ok := user.(*model.User); ok {
			c.Next()
			return
		}
	}

	c.JSON(200, serializer.CheckLogin())
	c.Abort()
}
