package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type ClientSettings struct {
	DataRetentionDays int `json:"data_retention_days" bson:"data_retention_days"`
	AlertingEnabled  bool `json:"alerting_enabled" bson:"alerting_enabled"`
	Timezone          string `json:"timezone" bson:"timezone"`
}

type Client struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Slug        string             `json:"slug" bson:"slug"`
	Email       string             `json:"email" bson:"email"`
	Description string             `json:"description" bson:"description"`
	Website     string             `json:"website" bson:"website"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy   primitive.ObjectID `json:"created_by" bson:"created_by"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	Settings    ClientSettings     `json:"settings" bson:"settings"`
}
