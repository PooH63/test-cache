package api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// ErrorHandler обработчик ошибок.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := WrapContext(c)
		ctx.Next()

		ginError := ctx.Errors.Last()
		if ginError == nil {
			return
		}

		ctx.NewResponseError(errors.New(ginError.Error()))
	}
}
