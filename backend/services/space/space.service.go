package spaceService

import (
	"context"
	"fmt"
	"time"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateSpace(ctx context.Context, space *schema.SpaceSchema) (*schema.SpaceSchema, error) {
	currentTime := time.Now()

	if err := space.Validate(); err != nil {
		return nil, err
	}

	space.CreatedAt = currentTime
	space.UpdatedAt = currentTime

	result, err := lib.Collections("spaces").InsertOne(ctx, space)
	if err != nil {
		return nil, err
	}

	var createdSpace schema.SpaceSchema
	err = lib.Collections("spaces").FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&createdSpace)
	if err != nil {
		return nil, err
	}

	return &createdSpace, nil
}


// GetAllSpaces retrieves all space documents in the collection.
func GetAllSpaces(ctx context.Context) ([]*schema.SpaceSchema, error) {
	var spaces []*schema.SpaceSchema

	cursor, err := lib.Collections("spaces").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var space schema.SpaceSchema
		if err := cursor.Decode(&space); err != nil {
			return nil, err
		}
		spaces = append(spaces, &space)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return spaces, nil
}

// GetSpaceById retrieves a space document by its ID.
func GetSpaceById(ctx context.Context, spaceID interface{}) (*schema.SpaceSchema, error) {
	var space schema.SpaceSchema

	err := lib.Collections("spaces").FindOne(ctx, bson.M{"_id": spaceID}).Decode(&space)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("space not found")
		}
		return nil, err
	}

	return &space, nil
}

// GetAllMySpaces retrieves all spaces created by a specific user.
func GetAllMySpaces(ctx context.Context, userID interface{}) ([]*schema.SpaceSchema, error) {
	var spaces []*schema.SpaceSchema

	cursor, err := lib.Collections("spaces").Find(ctx, bson.M{"creatorId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var space schema.SpaceSchema
		if err := cursor.Decode(&space); err != nil {
			return nil, err
		}
		spaces = append(spaces, &space)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return spaces, nil
}

// SearchSpaces retrieves spaces that match a search query, e.g., by name or description.
func SearchSpaces(ctx context.Context, query string) ([]*schema.SpaceSchema, error) {
	var spaces []*schema.SpaceSchema

	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	cursor, err := lib.Collections("spaces").Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var space schema.SpaceSchema
		if err := cursor.Decode(&space); err != nil {
			return nil, err
		}
		spaces = append(spaces, &space)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return spaces, nil
}


