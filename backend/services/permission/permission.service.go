package permissionService

import (
	"context"
	"fmt"
	"time"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreatePermission(ctx context.Context, permission *schema.PermissionSchema) (*mongo.InsertOneResult, error) {
	currentTime := time.Now()

	
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

	permission.CreatedAt = currentTime
	permission.UpdatedAt = currentTime

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

func GetPermissionByName(ctx context.Context, name string) (*schema.PermissionSchema, error) {
	var permission schema.PermissionSchema
	err := lib.Collections("permissions").FindOne(ctx, bson.M{"name": name}).Decode(&permission)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}


func GetPermissionsByNames(ctx context.Context, names []string) (*[]schema.PermissionSchema,error) {
	var permissions []schema.PermissionSchema
	permissionsCollection := lib.Collections("permissions")
	cursor, err := permissionsCollection.Find(context.TODO(), bson.M{"name": bson.M{"$in": names}})
	if err != nil {
		return &[]schema.PermissionSchema{}, fmt.Errorf("could not find permissions: %v", err)
	}
	if err = cursor.All(context.TODO(), &permissions); err != nil {
		return &[]schema.PermissionSchema{}, fmt.Errorf("could not decode permissions: %v", err)
	}
	return &permissions, nil
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
	currentTime := time.Now()

	if name, ok := update["name"]; ok {
		if name == "" {
			return nil, fmt.Errorf("missing required field: name")
		}
	}
	if description, ok := update["description"]; ok {
		if description == "" {
			return nil, fmt.Errorf("missing required field: description")
		}
	}
	update["updatedAt"] = currentTime
	return lib.Collections("permissions").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
}

func DeletePermission(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return lib.Collections("permissions").DeleteOne(ctx, bson.M{"_id": id})
}
