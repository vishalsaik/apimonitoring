package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApiKeyPermissions struct {
	CanIngestData    bool     `json:"can_ingest_data" bson:"can_ingest_data"`
	CanViewAnalytics bool     `json:"can_view_analytics" bson:"can_view_analytics"`
	AllowedServices  []string `json:"allowed_services" bson:"allowed_services"`
}
type ApiKey struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	KeyId       string              `json:"key_id" bson:"key_id"`
	KeyValue    string              `json:"key_value" bson:"key_value"`
	ClientID    primitive.ObjectID `json:"client_id" bson:"client_id"`
	Name        string              `json:"name" bson:"name"`
	Description string              `json:"description" bson:"description"`
	Environment string              `json:"environment" bson:"environment"`
	IsActive    bool                `json:"is_active" bson:"is_active"`
	ExpiresAt           time.Time          `json:"expires_at" bson:"expires_at"`
	Permissions ApiKeyPermissions         `json:"permissions" bson:"permissions"`
	Security	Security              `json:"security" bson:"security"`
	MetaData	MetaData              `json:"metadata" bson:"metadata"`
	CreatedBy   primitive.ObjectID  `json:"created_by" bson:"created_by"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}
type Security struct {
	AllowedIps          []string           `json:"allowed_ips" bson:"allowed_ips"`
	AllowedOrigins      []string           `json:"allowed_origins" bson:"allowed_origins"`
	LastRotatedAt       time.Time          `json:"last_rotated_at" bson:"last_rotated_at"`
	RotationWarningDays int                `json:"rotation_warning_days" bson:"rotation_warning_days"`
}
type MetaData struct {
	CreatedBy primitive.ObjectID `json:"created_by" bson:"created_by"`
	Purpose   string             `json:"purpose" bson:"purpose"`
	Tags      []string           `json:"tags" bson:"tags"`
}
