package route

import (
	"net/http"
	"tungit/apps/controllers/example"
	middleware "tungit/apps/middlewares"
	config "tungit/configs"
	"tungit/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter(config *config.ConfigAppEnv) {
	router := gin.Default()
	port := config.Port
	host := config.Host
	prefix := config.Prefix

	middleware := middleware.NewMiddleware()

	trustedProxy := "trusted_proxy_ip"
	router.SetTrustedProxies([]string{trustedProxy})
	router.Use(middleware.ErrorHandler)

	// Health check
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	// API V1 GROUP
	apiV1 := router.Group("/" + prefix + "/v1")
	{
		// Example
		eController := example.NewController()
		apiV1.GET("/login", eController.Login)
	}

	if config.Environment == "development" {
		docs.SwaggerInfo.Title = "Swagger Example API"
		docs.SwaggerInfo.Description = "This is a sample API."
		docs.SwaggerInfo.Version = "1.0.0"
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.Run(host + ":" + port)
}
