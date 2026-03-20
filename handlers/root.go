package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-monitoring/src/shared/utils"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success(gin.H{
		"service": "API Hit Monitoring System",
		"version": "1.0.0",
		"endpoints": gin.H{
			"health":    "/health",
			"ingest":    "/api/hit",
			"analytics": "/api/analytics",
		},
	}, "API Hit Monitoring Service", http.StatusOK))
}
