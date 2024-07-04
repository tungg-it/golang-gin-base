package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Schemes
// @Tags Example
// @Success 200
// @Failure 400 {object} errorException.BadRequestResponse
// @Failure 404 {object} errorException.NotFoundResponse
// @Failure 500 {object} errorException.InternalServerResponse
// @Router /login [get]
func (con *exampleController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
