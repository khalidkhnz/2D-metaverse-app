package profileService

import (
	"context"
	"fmt"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfileByAuthId(authId primitive.ObjectID) (*schema.ProfileSchema, error) {
	collection := lib.Collections("profiles")

	var profile schema.ProfileSchema

	err := collection.FindOne(context.TODO(), bson.M{"authId": authId}).Decode(&profile)
	if err != nil {
		return nil, fmt.Errorf("error retrieving profile: %v", err)
	}

	return &profile, nil
}