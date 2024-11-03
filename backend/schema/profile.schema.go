package schema

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileSchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AuthID    primitive.ObjectID `bson:"authId" json:"authId"`
	FullName  string             `bson:"firstName" json:"fullName"`
	Bio       string             `bson:"bio" json:"bio"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Username  string             `bson:"username" json:"username"`
	Role      string             `bson:"role" json:"role"`
	Status    string             `bson:"status" json:"status"`
	SocketID  string             `bson:"socketId" json:"socketId"`
}

func (p *ProfileSchema) Validate() error {
	if p.AuthID == primitive.NilObjectID {
		return errors.New("missing required field: authId")
	}
	if p.FullName == "" {
		return errors.New("missing required field: fullName")
	}
	if p.Username == "" {
		return errors.New("missing required field: username")
	}
	if p.Role == "" {
		return errors.New("missing required field: role")
	}
	return nil
}


func CreateProfile(authID primitive.ObjectID, fullName, bio, avatar, username, role, status, socketID string) *ProfileSchema {
	return &ProfileSchema{
		ID:        primitive.NewObjectID(),
		AuthID:    authID,
		FullName:  fullName,
		Bio:       bio,
		Avatar:    avatar,
		Username:  username,
		Role:      role,
		Status:    status,
		SocketID:  socketID,
	}
}
