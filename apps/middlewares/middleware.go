package middleware

import "github.com/gin-gonic/gin"

type IMiddleware interface {
	ErrorHandler(c *gin.Context)
}

type middleware struct {
}

func NewMiddleware() IMiddleware {
	return &middleware{}
}
