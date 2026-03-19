package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	RoleSuperAdmin Role = "super_admin"
	RoleClientAdmin Role = "client_admin"
	RoleClientViewer Role = "client_viewer"
)

type UserPermissions struct {
	CanCreateApiKeys bool `json:"can_create_api_keys" bson:"can_create_api_keys"`
	CanManageUsers   bool `json:"can_manage_users" bson:"can_manage_users"`
	CanViewAnalytics bool `json:"can_view_analytics" bson:"can_view_analytics"`
	CanExportData    bool `json:"can_export_data" bson:"can_export_data"`
}
type User struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserName    string              `json:"username" bson:"username"`
	Email       string              `json:"email" bson:"email"`
	Password    string              `json:"password" bson:"password"`
	Role        Role                `json:"role" bson:"role"`
	ClientID    *primitive.ObjectID `json:"client_id" bson:"client_id"`
	IsActive    bool                `json:"is_active" bson:"is_active"`
	Permissions UserPermissions     `json:"permissions" bson:"permissions"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}
