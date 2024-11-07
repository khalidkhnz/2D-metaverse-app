package roleService

import (
	"context"
	"fmt"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreateRole(ctx context.Context, role *schema.RoleSchema) (*mongo.InsertOneResult, error) {
	if err := role.Validate(); err != nil {
		return nil, err
	}

	// Check if role name already exists
	var existingRoles schema.RoleSchema
	err := lib.Collections("roles").FindOne(ctx, bson.M{"name": role.Name}).Decode(&existingRoles)
	if err == nil {
		return nil, fmt.Errorf("role with name %s already exists", role.Name)
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	return lib.Collections("roles").InsertOne(ctx, role)
}

func GetRoleByID(ctx context.Context, id primitive.ObjectID) (*schema.RoleSchema, error) {
	var role schema.RoleSchema
	err := lib.Collections("roles").FindOne(ctx, bson.M{"_id": id}).Decode(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}


func GetAllRoles(ctx context.Context) ([]schema.RoleSchema, error) {
	var roles []schema.RoleSchema
	cursor, err := lib.Collections("roles").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}



func UpdateRole(ctx context.Context, id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error) {
	if _, ok := update["name"]; ok && update["name"] == "" {
		return nil, fmt.Errorf("missing required field: name")
	}
	if _, ok := update["description"]; ok && update["description"] == "" {
		return nil, fmt.Errorf("missing required field: description")
	}
	return lib.Collections("roles").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
}

func DeleteRole(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return lib.Collections("roles").DeleteOne(ctx, bson.M{"_id": id})
}
