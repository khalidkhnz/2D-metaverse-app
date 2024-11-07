package types

import (
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginBody struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}


type AuthSchemaPopulated struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	FullName    string               `bson:"fullName" json:"fullName"`
	Email       string               `bson:"email" json:"email"`
	Password    string               `bson:"password" json:"password"`
	Role      schema.RoleSchema             `bson:"role" json:"role"`
	Permissions []schema.PermissionSchema   `bson:"permissions" json:"permissions"`
	Profile     schema.ProfileSchema        `bson:"profile" json:"profile"`
}

type FullProfile struct {
	User *AuthSchemaPopulated	`bson:"user" json:"user"`
}
