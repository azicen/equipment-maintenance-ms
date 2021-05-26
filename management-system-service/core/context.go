package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			Context: c,
			code:    http.StatusOK,
		}
		h(ctx)
	}
}

type Context struct {
	*gin.Context

	code int
	data interface{}
	Errors []error
}

//GetData 获取HTTP返回数据
func (c *Context) GetData() interface{} {
	return c.data
}

//SetData 设置HTTP返回数据
func (c *Context) SetData(data interface{}) {
	c.data = data
}

//GetCode 获取http code
func (c *Context) GetCode() int {
	return c.code
}

//SetCode 设置http code
func (c *Context) SetCode(code int) {
	c.code = code
}

//Error 添加错误
func (c *Context) Error(code int, errs ...error) {
	c.SetCode(code)
	for _, err := range errs {
		if err == nil {
			panic("err is nil")
		}
		c.Errors = append(c.Errors, err)
	}
}

//IsError 判断Context上下文是否有汇报错误
func (c *Context) IsError() bool {
	return len(c.Errors) > 0
}
