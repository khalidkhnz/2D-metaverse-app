package authService

import (
	"context"
	"fmt"
	"strings"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)
func CreateAccount(authBody schema.AuthSchema) (*schema.AuthSchema, error) {
	// GETTING ALL COLLECTIONS
	authCollection := lib.Collections("auths")
	profileCollection := lib.Collections("profiles")

	// HASING PASSWORD BEFORE SAVE
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authBody.Password), 12)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %v", err)
	}
	// CHANGED WITH HASHED PASSWORD
	authBody.Password = string(hashedPassword)
	authBody.Email = strings.ToLower(string(authBody.Email))

	// SAVING TO DB
	result, err := authCollection.InsertOne(context.TODO(), authBody)
	if err != nil {
		return nil, fmt.Errorf("could not create auth: %v", err)
	}

	// GETTING AUTH ID
	authID := result.InsertedID.(primitive.ObjectID)

	// CREATING PROFILE
	profile := schema.CreateProfile(authID, authBody.FullName, "", "", "", "", "", "")

	// INSERTING PROFILE
	_, err = profileCollection.InsertOne(context.TODO(), profile)
	if err != nil {
		return nil, fmt.Errorf("could not create profile: %v", err)
	}

	return &authBody, nil
}


func IsAccountAlreadyExist(authBody schema.AuthSchema) (bool,error) {
	collection := lib.Collections("auths")
	var auth schema.AuthSchema
	err := collection.FindOne(context.TODO(), bson.M{"email": strings.ToLower(string(authBody.Email))}).Decode(&auth)
	if err!=nil {
		return false,fmt.Errorf("error %v",err.Error())
	}
	if(auth.ID != primitive.NilObjectID){
		return true, nil
	}
	return false, nil
}


func Login(authBody types.LoginBody) (bool,schema.AuthSchema,error) {
	
	// GET COLLECTIONS
	collection := lib.Collections("auths")
	var auth schema.AuthSchema

	// FIND DOC
	err := collection.FindOne(context.TODO(), bson.M{"email": strings.ToLower(string(authBody.Email))}).Decode(&auth)
	if err!=nil {
		return false,schema.AuthSchema{},fmt.Errorf("error %v",err.Error())
	}
	
	// COMPARE PASSWORD
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(authBody.Password))
	if err != nil {
		return false,schema.AuthSchema{}, fmt.Errorf("invalid password: %v", err)
	}

	return true,auth, nil
}


func GetUserByUserId(userID primitive.ObjectID) (schema.AuthSchema, error) {
	// Get the "auths" collection
	collection := lib.Collections("auths")
	var auth schema.AuthSchema

	// Query the database for a user with the matching userID
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&auth)
	if err != nil {
		return schema.AuthSchema{}, fmt.Errorf("could not find user: %v", err)
	}

	return auth, nil
}