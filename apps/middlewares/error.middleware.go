package middleware

import (
	"net/http"
	errorMessage "tungit/pkg/constants/error-messages"
	errorException "tungit/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (con *middleware) ErrorHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		switch e := err.Err.(type) {
		case errorException.Http:
			c.AbortWithStatusJSON(e.StatusCode, e)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorException.Http{
				Message:    errorMessage.InternalServer,
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	}
}
