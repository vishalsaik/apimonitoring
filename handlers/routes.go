package handlers

import (
	"net/http"

	"api-monitoring/app"
	"api-monitoring/src/shared/middleware"
	"api-monitoring/src/shared/models"
	"api-monitoring/src/shared/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *app.App) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{app.Config.CorsConfig.AllowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	router.Use(middleware.LoggerHandler(app.Log))
	router.Use(middleware.ErrorHandler(app.Log))

	router.GET("/health", HealthCheckHandler)
	router.GET("/", RootHandler)
	auth := router.Group("/auth")
	{
		auth.POST("/onboard-super-admin", app.AuthController.OnboardSuperAdmin)
		auth.POST("/login", app.AuthController.Login)
		protected := auth.Group("/")
		protected.Use(middleware.Authenticate(app.Config.JwtConfig.SecretKey))
		{
			protected.GET("/profile", app.AuthController.GetProfile)
			protected.POST("/logout", app.AuthController.Logout)
			protected.POST("/register", middleware.Authorize(models.RoleSuperAdmin), app.AuthController.Register)

		}
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Endpoint not found", http.StatusNotFound, nil))
	})
	return router
}
