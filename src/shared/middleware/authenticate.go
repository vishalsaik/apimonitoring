package middleware

import (
	"errors"
	"net/http"

	"api-monitoring/src/shared/models"
	"api-monitoring/src/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("authToken")
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Authentication token is required", http.StatusUnauthorized, nil))
			return
		}

		claims := &models.JWTClaims{}
		_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			msg := "Invalid token"
			if errors.Is(err, jwt.ErrTokenExpired) {
				msg = "Token expired"
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(msg, http.StatusUnauthorized, nil))
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("clientId", claims.ClientID)
		c.Next()
	}
}
