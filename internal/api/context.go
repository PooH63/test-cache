package api

import (
	"github.com/gin-gonic/gin"
)

// Context структура контекста.
type Context struct {
	*gin.Context
}

// WrapContext обёртка контекста.
func WrapContext(c *gin.Context) *Context {
	return &Context{c}
}

// HandlerFunc func.
type HandlerFunc func(*Context) error

// Handler обёртка над gin.Context.
func Handler(f HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(WrapContext(c)); err != nil {
			c.Error(err)
		}
	}
}