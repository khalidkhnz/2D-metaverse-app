package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	roleService "github.com/khalidkhnz/2D-metaverse-app/backend/services/role"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
)


func RoleCheckerMiddleware(next http.Handler, roles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fetchedRoles, err := roleService.GetRolesByNames(context.TODO(), roles)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Invalid Role",
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

		for _, element := range *fetchedRoles {
			if element.ID.String() == user.User.Role.ID.String() {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid Role",
		})

		return
	})
}
