package repository

import (
	"api-monitoring/src/shared/config/logger"
	"api-monitoring/src/shared/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	// MongoDB client and collection details would go here
	collection *mongo.Collection
	log        *logger.Logger
}

func NewMongoUserRepository(col *mongo.Collection, log *logger.Logger) *MongoUserRepository {
	return &MongoUserRepository{
		collection: col,
		log:        log,
	}
}
func (r *MongoUserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	result := r.collection.FindOne(ctx, bson.M{"_id": id})
	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *MongoUserRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, cursor.Err()
}

func (r *MongoUserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return r.FindByID(ctx, id)
}
func (r *MongoUserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	result := r.collection.FindOne(ctx, bson.M{"username": username})
	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	result := r.collection.FindOne(ctx, bson.M{"email": email})
	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
