package schema

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermissionsSchema struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}

func (p *PermissionsSchema) Validate() error {
	if p.Name == "" {
		return errors.New("missing required field: name")
	}
	if p.Description == "" {
		return errors.New("missing required field: description")
	}
	return nil
}

func CreatePermission(name, description string) *PermissionsSchema {
	return &PermissionsSchema{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: description,
	}
}