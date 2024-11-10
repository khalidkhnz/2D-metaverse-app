package authService

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	permissionService "github.com/khalidkhnz/2D-metaverse-app/backend/services/permission"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


func CreateAccount(authBody schema.AuthSchema,spaceCreator bool) (*schema.AuthSchema, error) {
	currentTime := time.Now()

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
	if authBody.Permissions == nil {
		authBody.Permissions = []primitive.ObjectID{}
	}
	authBody.CreatedAt = currentTime
	authBody.UpdatedAt = currentTime

	
	// GET PERMISSIONS 
	if spaceCreator {
		creatorPermission,err := permissionService.GetPermissionByName(context.TODO(),"SPACE_CREATOR")
		if err!=nil {
			return &authBody ,fmt.Errorf(err.Error())
		}
		authBody.Permissions = []primitive.ObjectID{
			creatorPermission.ID,	
		}
	}

	// SAVING TO DB
	result, err := authCollection.InsertOne(context.TODO(), authBody)
	if err != nil {
		return nil, fmt.Errorf("could not create auth: %v", err)
	}

	// GETTING AUTH ID
	authID := result.InsertedID.(primitive.ObjectID)

	// CREATING PROFILE
	profile := schema.CreateProfileDoc(authID, authBody.FullName, "", "", "", "", "", "")

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


func GetUserByUserId(userID primitive.ObjectID) (*schema.AuthSchema, error) {
	// Get the "auths" collection
	collection := lib.Collections("auths")
	var auth schema.AuthSchema

	// Query the database for a user with the matching userID
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&auth)
	if err != nil {
		return &schema.AuthSchema{}, fmt.Errorf("could not find user: %v", err)
	}

	return &auth, nil
}


func GetPopulatedUserByUserId(userID primitive.ObjectID) (*types.AuthSchemaPopulated, error) {
	// Get the "auths" collection
	collection := lib.Collections("auths")
	var auth schema.AuthSchema

	// Query the database for a user with the matching userID
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&auth)
	if err != nil {
		return &types.AuthSchemaPopulated{}, fmt.Errorf("could not find user: %v", err)
	}

	// Populate Role
	var role schema.RoleSchema
	roleCollection := lib.Collections("roles")
	err = roleCollection.FindOne(context.TODO(), bson.M{"_id": auth.RoleId}).Decode(&role)
	if err != nil {
		return &types.AuthSchemaPopulated{}, fmt.Errorf("could not find role: %v", err)
	}
	
	// Populate Profile
	var profile schema.ProfileSchema
	profileCollection := lib.Collections("profiles")
	err = profileCollection.FindOne(context.TODO(), bson.M{"authId": auth.ID}).Decode(&profile)
	if err != nil {
		return &types.AuthSchemaPopulated{}, fmt.Errorf("could not find profile: %v", err)
	}

	// Populate Permissions
	var permissions []schema.PermissionSchema
	permissionCollection := lib.Collections("permissions")
	cursor, err := permissionCollection.Find(context.TODO(), bson.M{"_id": bson.M{"$in": auth.Permissions}})
	if err != nil {
		return &types.AuthSchemaPopulated{}, fmt.Errorf("could not find permissions: %v", err)
	}
	if err = cursor.All(context.TODO(), &permissions); err != nil {
		return &types.AuthSchemaPopulated{}, fmt.Errorf("could not decode permissions: %v", err)
	}

	// Create populated user
	populatedAuth := types.AuthSchemaPopulated{
		ID:          auth.ID,
		FullName:    auth.FullName,
		Email:       auth.Email,
		Password:    auth.Password,
		Role:      role,
		Permissions: permissions,
		Profile: 	 profile,
	}

	return &populatedAuth, nil
}


func GetUserFromToken(tokenStr string, checkWsToken bool) (*types.FullProfile, error) {
	// Parse and validate token
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return lib.GetJWTSecret(), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	// Extract userID from token claims
	userIDStr, ok := (*claims)["userID"].(string)
	if !ok {
		return nil, fmt.Errorf("userID not found in token claims")
	}

	if checkWsToken {
		isWS, ok := (*claims)["isWS"].(bool)
		if !ok || !isWS {
			return nil, fmt.Errorf("its not a Websocket token")
		}
	}

	// Convert userID string to ObjectID
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid userID format")
	}

	// Fetch user data from database
	populatedUser, err := GetPopulatedUserByUserId(userID)
	if err != nil {
		return nil, err
	}

	return &types.FullProfile{User: populatedUser}, nil
}
