package schema

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleSchema struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (r *RoleSchema) Validate() error {
	if r.Name == "" {
		return errors.New("missing required field: name")
	}
	if r.Description == "" {
		return errors.New("missing required field: description")
	}
	return nil
}

func CreateRoleDoc(name, description string) *RoleSchema {
	return &RoleSchema{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: description,
	}
}