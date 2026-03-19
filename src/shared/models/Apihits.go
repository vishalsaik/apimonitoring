package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HTTPMethod string

const APIHitCollection = "api_hits"
const (
	MethodGet     HTTPMethod = "GET"
	MethodPost    HTTPMethod = "POST"
	MethodPut     HTTPMethod = "PUT"
	MethodDelete  HTTPMethod = "DELETE"
	MethodPatch   HTTPMethod = "PATCH"
	MethodOptions HTTPMethod = "OPTIONS"
	MethodHead    HTTPMethod = "HEAD"
)

type APIHit struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	EventID     string             `json:"event_id" bson:"event_id"`
	TimeStamp   time.Time          `json:"timestamp" bson:"timestamp"`
	ServiceName string             `json:"service_name" bson:"service_name"`
	Endpoint    string             `json:"endpoint" bson:"endpoint"`
	Method      HTTPMethod         `json:"method" bson:"method"`
	StatusCode  int                `json:"status_code" bson:"status_code"`
	LatencyMs   float64            `json:"latency_ms" bson:"latency_ms"`
	ClientID    primitive.ObjectID `json:"client_id" bson:"client_id"`
	ApiKeyID    primitive.ObjectID `json:"api_key_id" bson:"api_key_id"`
	IP          string             `json:"ip" bson:"ip"`
	UserAgent   string             `json:"user_agent" bson:"user_agent"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
