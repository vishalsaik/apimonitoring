package handlers

import (
	"api-monitoring/src/shared/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success(gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}, "Service is healthy", http.StatusOK))
}
