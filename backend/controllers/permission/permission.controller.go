package permissionsController

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	permissionsService "github.com/khalidkhnz/2D-metaverse-app/backend/services/permission"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func HandleCreatePermissions(w http.ResponseWriter, r *http.Request) error {
	var createPermissionsBody *schema.PermissionSchema
	if err := json.NewDecoder(r.Body).Decode(&createPermissionsBody); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	err := createPermissionsBody.Validate()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	permissionsDoc, err := permissionsService.CreatePermission(context.TODO(), createPermissionsBody)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Permissions Created",
		"data":    permissionsDoc,
	})
}

func HandleUpdatePermissions(w http.ResponseWriter, r *http.Request) error {
	var updatePermissionsBody *schema.PermissionSchema
	if err := json.NewDecoder(r.Body).Decode(&updatePermissionsBody); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	err := updatePermissionsBody.Validate()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	permissionsDoc, err := permissionsService.UpdatePermission(context.TODO(), updatePermissionsBody.ID,bson.M{
		"name":updatePermissionsBody.Name,
		"description":updatePermissionsBody.Description,
	})
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Permissions Updated",
		"data":    permissionsDoc,
	})
}

func HandleGetPermissions(w http.ResponseWriter, r *http.Request) error {
	var permissionsIds struct {
		PermissionsIds []primitive.ObjectID `bson:"permissionsIds" json:"permissionsIds"`
	}
	if err := json.NewDecoder(r.Body).Decode(&permissionsIds); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	var errorArray []string
	var docArray []*schema.PermissionSchema

	for _, permissionsId := range permissionsIds.PermissionsIds {
		permissionsDoc, err := permissionsService.GetPermissionByID(context.TODO(), permissionsId)
		if err != nil {
			errorArray = append(errorArray, err.Error())
			continue
		}
		docArray = append(docArray, permissionsDoc)
	}
	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data": map[string]any{
			"validDocs": docArray,
			"errors":    errorArray,
		},
	})
}

func HandleGetAllPermissions(w http.ResponseWriter, r *http.Request) error {
	permissionsDoc, err := permissionsService.GetAllPermissions(context.TODO())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    permissionsDoc,
	})
}

func HandleDeletePermissions(w http.ResponseWriter, r *http.Request) error {
	var deletePermissionsBody struct {
		PermissionsID primitive.ObjectID `bson:"permissionsId" json:"permissionsId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&deletePermissionsBody); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	_,err := permissionsService.DeletePermission(context.TODO(), deletePermissionsBody.PermissionsID)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Permissions Deleted",
	})
}
