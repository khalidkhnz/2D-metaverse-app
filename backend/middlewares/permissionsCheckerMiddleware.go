package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	permissionService "github.com/khalidkhnz/2D-metaverse-app/backend/services/permission"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
)

func PermissionsCheckerMiddleware(next http.Handler, roles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fetchedPermissions, err := permissionService.GetPermissionsByNames(context.TODO(), roles)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Invalid Permission",
			})
			return
		}

		user, ok := r.Context().Value("user").(types.FullProfile)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Invalid User",
			})
			return
		}

		userPermissions := user.User.Permissions

		// Check if all fetched permissions exist in userPermissions
		allPermissionsExist := true
		for _, fetchedPermission := range *fetchedPermissions {
			found := false
			for _, userPermission := range userPermissions {
				if fetchedPermission.ID == userPermission.ID {
					found = true
					break
				}
			}
			if !found {
				allPermissionsExist = false
				break
			}
		}

		if allPermissionsExist {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid Permission",
		})

		return
	})
}
