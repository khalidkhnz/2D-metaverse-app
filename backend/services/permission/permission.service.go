package permissionService

import (
	"context"
	"fmt"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreatePermission(ctx context.Context, permission *schema.PermissionSchema) (*mongo.InsertOneResult, error) {
	if err := permission.Validate(); err != nil {
		return nil, err
	}

	// Check if permission name already exists
	var existingPermission schema.PermissionSchema
	err := lib.Collections("permissions").FindOne(ctx, bson.M{"name": permission.Name}).Decode(&existingPermission)
	if err == nil {
		return nil, fmt.Errorf("permission with name %s already exists", permission.Name)
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	return lib.Collections("permissions").InsertOne(ctx, permission)
}

func GetPermissionByID(ctx context.Context, id primitive.ObjectID) (*schema.PermissionSchema, error) {
	var permission schema.PermissionSchema
	err := lib.Collections("permissions").FindOne(ctx, bson.M{"_id": id}).Decode(&permission)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}


func GetAllPermissions(ctx context.Context) ([]schema.PermissionSchema, error) {
	var permissions []schema.PermissionSchema
	cursor, err := lib.Collections("permissions").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &permissions); err != nil {
		return nil, err
	}
	return permissions, nil
}



func UpdatePermission(ctx context.Context, id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error) {
	if _, ok := update["name"]; ok && update["name"] == "" {
		return nil, fmt.Errorf("missing required field: name")
	}
	if _, ok := update["description"]; ok && update["description"] == "" {
		return nil, fmt.Errorf("missing required field: description")
	}
	return lib.Collections("permissions").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
}

func DeletePermission(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return lib.Collections("permissions").DeleteOne(ctx, bson.M{"_id": id})
}
