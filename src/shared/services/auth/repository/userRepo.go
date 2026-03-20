package repository

import (
	"api-monitoring/src/shared/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
