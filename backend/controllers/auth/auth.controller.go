package authControllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	authService "github.com/khalidkhnz/2D-metaverse-app/backend/services/auth"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	
	authID := mux.Vars(r)["id"]
	if authID == "" {
		return fmt.Errorf("missing required parameter: id")
	}

	objectID, err := primitive.ObjectIDFromHex(authID)
	if err != nil {
		return fmt.Errorf("invalid id format: %v", err)
	}

	collection := lib.Collections("auths")
	var auth schema.AuthSchema
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&auth)
	
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("auth not found")
		}
		return fmt.Errorf("could not retrieve auth: %v", err)
	}

	return lib.WriteJSON(w, http.StatusOK, auth)
}


func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	var loginCread types.LoginBody
	if err := json.NewDecoder(r.Body).Decode(&loginCread); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	loginCread.Email = strings.ToLower(string(loginCread.Email))

	isMatched,data,err := authService.Login(loginCread)

	if isMatched {

		//GENERATE JWT TOKEN
		token, err := lib.GenerateJWTToken(string(data.ID.Hex())) 
		if err != nil {
			return fmt.Errorf("could not generate JWT token: %v", err)
		}

		return lib.WriteJSON(w, http.StatusOK, map[string]any{
			"success":true,
			"data": map[string]any{"auth":map[string]any{"_id":data.ID,"fullname":data.FullName,"email":data.Email}, "token":token},
			"message": "Logged In",
		})
	}

	if err!=nil {
		return lib.WriteJSON(w, http.StatusOK, map[string]any{
			"success":false,
			"message": err.Error(),
		})
	}

	return fmt.Errorf("incorrect password")
}


func HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	var auth schema.AuthSchema
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	if err := auth.Validate(); err != nil {
		return err
	}

	isExist, _ := authService.IsAccountAlreadyExist(auth);
	
	if isExist {
		return fmt.Errorf("account already exist")
	}

	_, err := authService.CreateAccount(auth)
	if err!=nil {
		return err
	}

	return lib.WriteJSON(w, http.StatusCreated, map[string]any{
		"success":true,
		"message": "Account Created",
		"data":auth,
	})
}


func HandleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	authID := r.URL.Query().Get("id")
	if authID == "" {
		return fmt.Errorf("missing required parameter: id")
	}

	objectID, err := primitive.ObjectIDFromHex(authID)
	if err != nil {
		return fmt.Errorf("invalid id format: %v", err)
	}

	collection := lib.Collections("auths")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("could not delete auth: %v", err)
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "auth deleted",
		"success":true,
	})
}


func HandleCurrentUser(w http.ResponseWriter, r *http.Request) error {
	// Retrieve the user data from the context
	user, ok := r.Context().Value("user").(types.FullProfile)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return nil
	}

	// Respond with the user data
	return lib.WriteJSON(w,http.StatusOK,map[string]any{
		"success": true,
		"data":    user,
	})
}