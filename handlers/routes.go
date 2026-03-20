package handlers

import (
	"net/http"

	"api-monitoring/src/shared/config/logger"
	"api-monitoring/src/shared/middleware"
	"api-monitoring/src/shared/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(log *logger.Logger) *gin.Engine {
	router := gin.New()
	router.Use(middleware.LoggerHandler(log))
	router.Use(middleware.ErrorHandler(log))

	router.GET("/health", HealthCheckHandler)
	router.GET("/", RootHandler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Endpoint not found", http.StatusNotFound, nil))
	})

	// api := router.Group("/api")
	// {
	// 	api.POST("/users", userCtrl.Create)  ← add as controllers are built
	// }

	return router
}
