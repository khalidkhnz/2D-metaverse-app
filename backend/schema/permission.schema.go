package schema

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermissionSchema struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (p *PermissionSchema) Validate() error {
	if p.Name == "" {
		return errors.New("missing required field: name")
	}
	if p.Description == "" {
		return errors.New("missing required field: description")
	}
	return nil
}

func CreatePermissionDoc(name, description string) *PermissionSchema {
	return &PermissionSchema{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: description,
	}
}