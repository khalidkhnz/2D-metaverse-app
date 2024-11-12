package spaceController

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	spaceService "github.com/khalidkhnz/2D-metaverse-app/backend/services/space"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleCreateSpace(w http.ResponseWriter, r *http.Request) error {
	var createSpaceBody *schema.SpaceSchema
	if err := json.NewDecoder(r.Body).Decode(&createSpaceBody); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	user := lib.UserInContext(w, r)
	createSpaceBody.CreatorId = user.User.ID

	err := createSpaceBody.Validate()

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	createSpaceBody.MemberIds = append(createSpaceBody.MemberIds, user.User.ID)

	spaceDoc, err := spaceService.CreateSpace(context.TODO(), createSpaceBody)
	if err != nil {
		return fmt.Errorf(err.Error())
	}


	// ADD SPACE ID TO CREATOR'S AUTH DOC
	update := bson.M{
		"$push": bson.M{
			"spaceIds": spaceDoc.ID,
		},
	}
	_ , err = lib.Collections("auths").UpdateByID(context.TODO(), user.User.ID, update)
	if err != nil {
		return err
	}


	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Space Created",
		"data":    spaceDoc,
	})
}



// HandleGetAllSpaces handles fetching all spaces
func HandleGetAllSpaces(w http.ResponseWriter, r *http.Request) error {
	spaces, err := spaceService.GetAllSpaces(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to fetch spaces: %v", err)
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    spaces,
		"message": "All Spaces",
	})
}



// HandleGetSpaceById handles fetching a space by ID
func HandleGetSpaceById(w http.ResponseWriter, r *http.Request) error {
	spaceIDParam := r.URL.Query().Get("id")
	spaceID, err := primitive.ObjectIDFromHex(spaceIDParam)
	if err != nil {
		return fmt.Errorf("invalid space ID: %v", err)
	}

	space, err := spaceService.GetSpaceById(context.TODO(), spaceID)
	if err != nil {
		return fmt.Errorf("failed to fetch space: %v", err)
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    space,
		"message": "Spaces found",
	})
}



// HandleGetAllMySpaces handles fetching all spaces created by the current user
func HandleGetAllMySpaces(w http.ResponseWriter, r *http.Request) error {
	user := lib.UserInContext(w, r)

	spaces, err := spaceService.GetAllMySpaces(context.TODO(), user.User.ID)
	if err != nil {
		return fmt.Errorf("failed to fetch user's spaces: %v", err)
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    spaces,
		"message": "All My Spaces",
	})
}



// HandleSearchSpaces handles searching for spaces based on a query parameter
func HandleSearchSpaces(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query().Get("q")

	spaces, err := spaceService.SearchSpaces(context.TODO(), query)
	if err != nil {
		return fmt.Errorf("failed to search spaces: %v", err)
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    spaces,
		"message": "Spaces Search Result",
	})
}