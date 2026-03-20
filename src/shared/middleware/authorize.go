package middleware

import (
	"net/http"

	"api-monitoring/src/shared/models"
	"api-monitoring/src/shared/utils"

	"github.com/gin-gonic/gin"
)

func Authorize(roles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil))
			return
		}

		userRole, ok := role.(models.Role)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil))
			return
		}
		if len(roles) == 0 {
			c.Next()
			return
		}

		for _, r := range roles {
			if userRole == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, utils.ErrorResponse("Insufficient permissions", http.StatusForbidden, nil))
	}
}
