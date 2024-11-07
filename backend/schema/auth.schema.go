package schema

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthSchema struct {
	ID        	primitive.ObjectID 		`bson:"_id,omitempty" json:"_id"`
	FullName 	string              	`bson:"fullName" json:"fullName"`
	Email     	string             		`bson:"email" json:"email"`
	Password  	string             		`bson:"password" json:"password"`
	RoleId  	primitive.ObjectID      `bson:"roleId" json:"roleId"`
	Permissions []primitive.ObjectID	`bson:"permissions" json:"permissions"`
	AvatarId  	primitive.ObjectID      `bson:"avatarId" json:"avatarId"`
}

func (a *AuthSchema) Validate() error {
	if a.FullName == "" {
		return errors.New("missing required field: fullName")
	}
	if a.Email == "" {
		return errors.New("missing required field: email")
	}
	if a.Password == "" {
		return errors.New("missing required field: password")
	}
	if a.RoleId.String() == "" {
		return errors.New("missing required field: roleId")
	}
	return nil
}

func CreateAuth(fullName, email, password string) *AuthSchema {
	return &AuthSchema{
		ID:        primitive.NewObjectID(),
		FullName:  fullName,
		Email:     email,
		Password:  password,
	}
}
