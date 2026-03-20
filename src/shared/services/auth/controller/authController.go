package controller

import (
	"net/http"

	"api-monitoring/src/shared/models"
	"api-monitoring/src/shared/config"
	"api-monitoring/src/shared/services/auth/service"
	"api-monitoring/src/shared/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OnboardRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	Username string      `json:"username" binding:"required"`
	Email    string      `json:"email"    binding:"required,email"`
	Password string      `json:"password" binding:"required,min=8"`
	Role     models.Role `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthController struct {
	authService *service.AuthService
	cookie      config.CookieConfig
}

func NewAuthController(authService *service.AuthService, cookie config.CookieConfig) *AuthController {
	return &AuthController{
		authService: authService,
		cookie:      cookie,
	}
}

func (ctrl *AuthController) setAuthCookie(c *gin.Context, token string) {
	c.SetCookie("authToken", token, ctrl.cookie.MaxAge, "/", "", ctrl.cookie.Secure, ctrl.cookie.HttpOnly)
}

func (ctrl *AuthController) OnboardSuperAdmin(c *gin.Context) {
	var req OnboardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewAppError("Invalid request body", http.StatusBadRequest, nil))
		return
	}

	user := &models.User{
		UserName: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     models.RoleSuperAdmin,
	}

	resp, token, err := ctrl.authService.OnboardSuperAdmin(c.Request.Context(), user)
	if err != nil {
		c.Error(err)
		return
	}

	ctrl.setAuthCookie(c, token)
	c.JSON(http.StatusCreated, utils.Success(resp, "Super admin created successfully", http.StatusCreated))
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewAppError("Invalid request body", http.StatusBadRequest, nil))
		return
	}

	if req.Role == "" {
		req.Role = models.RoleClientViewer
	}

	user := &models.User{
		UserName: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	resp, token, err := ctrl.authService.Register(c.Request.Context(), user)
	if err != nil {
		c.Error(err)
		return
	}

	ctrl.setAuthCookie(c, token)
	c.JSON(http.StatusCreated, utils.Success(resp, "User registered successfully", http.StatusCreated))
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewAppError("Invalid request body", http.StatusBadRequest, nil))
		return
	}

	resp, token, err := ctrl.authService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	ctrl.setAuthCookie(c, token)
	c.JSON(http.StatusOK, utils.Success(resp, "Logged in successfully", http.StatusOK))
}

func (ctrl *AuthController) GetProfile(c *gin.Context) {
	userIDVal, exists := c.Get("userId")
	if !exists {
		c.Error(utils.NewUnauthorizedError("Unauthorized"))
		return
	}

	userID, ok := userIDVal.(primitive.ObjectID)
	if !ok {
		c.Error(utils.NewUnauthorizedError("Invalid user ID"))
		return
	}

	resp, err := ctrl.authService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, utils.Success(resp, "Profile fetched successfully", http.StatusOK))
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	c.SetCookie("authToken", "", -1, "/", "", ctrl.cookie.Secure, ctrl.cookie.HttpOnly)
	c.JSON(http.StatusOK, utils.Success(nil, "Logged out successfully", http.StatusOK))
}
