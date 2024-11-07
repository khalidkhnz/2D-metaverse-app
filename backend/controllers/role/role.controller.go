package roleController

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/schema"
	roleService "github.com/khalidkhnz/2D-metaverse-app/backend/services/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func HandleCreateRole(w http.ResponseWriter, r *http.Request) error {
	var createRoleBody *schema.RoleSchema
	if err := json.NewDecoder(r.Body).Decode(&createRoleBody); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	err := createRoleBody.Validate()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	roleDoc, err := roleService.CreateRole(context.TODO(), createRoleBody)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Role Created",
		"data":    roleDoc,
	})
}

func HandleUpdateRole(w http.ResponseWriter, r *http.Request) error {
	var updateRoleBody *schema.RoleSchema
	if err := json.NewDecoder(r.Body).Decode(&updateRoleBody); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	err := updateRoleBody.Validate()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	roleDoc, err := roleService.UpdateRole(context.TODO(), updateRoleBody.ID,bson.M{
		"name":updateRoleBody.Name,
		"description":updateRoleBody.Description,
	})
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Role Updated",
		"data":    roleDoc,
	})
}

func HandleGetRole(w http.ResponseWriter, r *http.Request) error {
	var roleIds struct {
		RoleIds []primitive.ObjectID `bson:"roleIds" json:"roleIds"`
	}
	if err := json.NewDecoder(r.Body).Decode(&roleIds); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	var errorArray []string
	var docArray []*schema.RoleSchema

	for _, roleId := range roleIds.RoleIds {
		roleDoc, err := roleService.GetRoleByID(context.TODO(), roleId)
		if err != nil {
			errorArray = append(errorArray, err.Error())
			continue
		}
		docArray = append(docArray, roleDoc)
	}
	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data": map[string]any{
			"validDocs": docArray,
			"errors":    errorArray,
		},
	})
}

func HandleGetAllRoles(w http.ResponseWriter, r *http.Request) error {
	roleDoc, err := roleService.GetAllRoles(context.TODO())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    roleDoc,
	})
}

func HandleDeleteRole(w http.ResponseWriter, r *http.Request) error {
	var deleteRoleBody struct {
		RoleID primitive.ObjectID `bson:"roleId" json:"roleId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&deleteRoleBody); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	_,err := roleService.DeleteRole(context.TODO(), deleteRoleBody.RoleID)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return lib.WriteJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Role Deleted",
	})
}
