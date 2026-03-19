package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"api-monitoring/src/shared/config/logger"
	"api-monitoring/src/shared/utils"
)

func ErrorHandler(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		statusCode := http.StatusInternalServerError
		message := "Internal server error"
		var errors interface{}

		log.Error("Error occurred",
			zap.String("message", err.Error()),
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
		)

		switch e := err.(type) {
		case *utils.AppError:
			statusCode = e.StatusCode
			message = e.Message
			errors = e.Errors
		default:
			errStr := err.Error()
			if strings.Contains(errStr, "E11000") || strings.Contains(errStr, "duplicate key") {
				statusCode = http.StatusConflict
				message = "Duplicate key error"
			} else if strings.Contains(errStr, "token is expired") {
				statusCode = http.StatusUnauthorized
				message = "Token expired"
			} else if strings.Contains(errStr, "signature is invalid") || strings.Contains(errStr, "token contains an invalid number of segments") {
				statusCode = http.StatusUnauthorized
				message = "Invalid token"
			}
			_ = e
		}

		c.AbortWithStatusJSON(statusCode, utils.ErrorResponse(message, statusCode, errors))
	}
}
