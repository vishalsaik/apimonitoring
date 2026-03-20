package service

import (
	"context"
	"time"

	"api-monitoring/src/shared/config/logger"
	"api-monitoring/src/shared/models"
	"api-monitoring/src/shared/services/auth/repository"
	"api-monitoring/src/shared/utils"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// JWTClaims — payload inside the token

// UserResponse — user shape returned to client (no password)
type UserResponse struct {
	ID          primitive.ObjectID     `json:"id"`
	Username    string                 `json:"username"`
	Email       string                 `json:"email"`
	Role        models.Role            `json:"role"`
	ClientID    *primitive.ObjectID    `json:"clientId,omitempty"`
	IsActive    bool                   `json:"isActive"`
	Permissions models.UserPermissions `json:"userPermissions"`
	CreatedAt   time.Time              `json:"createdAt"`
}

type AuthService struct {
	repo      repository.UserRepository
	jwtSecret string
	jwtExpiry int
	log       *logger.Logger
}

func NewAuthService(repo repository.UserRepository, jwtSecret string, jwtExpiry int, log *logger.Logger) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
		log:       log,
	}
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	expiry := time.Duration(s.jwtExpiry) * time.Second
	claims := models.JWTClaims{
		UserID:   user.ID,
		Username: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
		ClientID: user.ClientID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) formatUserForResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Username:    user.UserName,
		Email:       user.Email,
		Role:        user.Role,
		ClientID:    user.ClientID,
		IsActive:    user.IsActive,
		Permissions: user.Permissions,
		CreatedAt:   user.CreatedAt,
	}
}

func (s *AuthService) comparePassword(plain, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func (s *AuthService) OnboardSuperAdmin(ctx context.Context, data *models.User) (UserResponse, string, error) {
	results, err := s.repo.FindAll(ctx)
	if err != nil {
		return UserResponse{}, "", err
	}
	if len(results) > 0 {
		return UserResponse{}, "", utils.NewAppError("Super admin onboarding is disabled", 403, nil)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, "", utils.NewAppError("Failed to hash password", 500, nil)
	}
	data.Password = string(hashed)
	data.IsActive = true
	data.Role = models.RoleSuperAdmin
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	users, err := s.repo.Create(ctx, data)
	if err != nil {
		return UserResponse{}, "", err
	}
	token, err := s.generateToken(users)
	if err != nil {
		return UserResponse{}, "", err
	}
	return s.formatUserForResponse(users), token, nil
}

func (s *AuthService) Register(ctx context.Context, data *models.User) (UserResponse, string, error) {
	user, err := s.repo.FindByUsername(ctx, data.UserName)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); !ok || appErr.StatusCode != 404 {
			return UserResponse{}, "", err
		}
	}
	if user != nil {
		return UserResponse{}, "", utils.NewAppError("Username already exists", 409, nil)
	}
	emailUser, err := s.repo.FindByEmail(ctx, data.Email)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); !ok || appErr.StatusCode != 404 {
			return UserResponse{}, "", err
		}
	}
	if emailUser != nil {
		return UserResponse{}, "", utils.NewAppError("Email already exists", 409, nil)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, "", utils.NewAppError("Failed to hash password", 500, nil)
	}
	data.Password = string(hashed)
	data.IsActive = true
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	users, err := s.repo.Create(ctx, data)
	if err != nil {
		return UserResponse{}, "", err
	}
	token, err := s.generateToken(users)
	if err != nil {
		return UserResponse{}, "", err
	}
	return s.formatUserForResponse(users), token, nil
}

func (s *AuthService) Login(ctx context.Context, username, password string) (UserResponse, string, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok && appErr.StatusCode == 404 {
			return UserResponse{}, "", utils.NewAppError("Invalid credentials", 401, nil)
		}
		return UserResponse{}, "", err
	}
	if !user.IsActive {
		return UserResponse{}, "", utils.NewAppError("Account is deactivated", 403, nil)
	}
	err = s.comparePassword(password, user.Password)
	if err != nil {
		return UserResponse{}, "", utils.NewAppError("Invalid credentials", 401, nil)
	}
	token, err := s.generateToken(user)
	if err != nil {
		return UserResponse{}, "", err
	}
	return s.formatUserForResponse(user), token, nil
}

func (s *AuthService) GetProfile(ctx context.Context, userID primitive.ObjectID) (UserResponse, error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return UserResponse{}, err
	}
	if user == nil {
		return UserResponse{}, utils.NewAppError("User not found", 404, nil)
	}
	return s.formatUserForResponse(user), nil
}
