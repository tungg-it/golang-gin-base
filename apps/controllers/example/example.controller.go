package example

import "github.com/gin-gonic/gin"

type IExampleController interface {
	Login(c *gin.Context)
}

type exampleController struct {
}

func NewController() IExampleController {
	return &exampleController{}
}
