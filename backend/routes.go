package main

import (
	"net/http"

	"github.com/gorilla/mux"
	authControllers "github.com/khalidkhnz/2D-metaverse-app/backend/controllers/auth"
	permissionController "github.com/khalidkhnz/2D-metaverse-app/backend/controllers/permission"
	roleController "github.com/khalidkhnz/2D-metaverse-app/backend/controllers/role"
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/middlewares"
)


func (s *APIServer) AuthRouter(router *mux.Router) {
	router.HandleFunc("/auth/signup", makeHTTPHandleFunc(authControllers.HandleCreateAccount)).Methods("POST")
	router.HandleFunc("/auth/get/{id}", makeHTTPHandleFunc(authControllers.HandleGetAccount)).Methods("GET")
	router.HandleFunc("/auth/login", makeHTTPHandleFunc(authControllers.HandleLogin)).Methods("POST")
	
	// REQUIRES TOKEN
	router.Handle("/auth/current-user", middlewares.AuthMiddleware(makeHTTPHandleFunc(authControllers.HandleCurrentUser))).Methods("GET")
}

func (s *APIServer) RoleRouter(router *mux.Router){
	router.Handle("/role/create-role", middlewares.AuthMiddleware(makeHTTPHandleFunc(roleController.HandleCreateRole))).Methods("POST")
	router.Handle("/role/delete-role", middlewares.AuthMiddleware(makeHTTPHandleFunc(roleController.HandleDeleteRole))).Methods("DELETE")
	router.Handle("/role/get-role", middlewares.AuthMiddleware(makeHTTPHandleFunc(roleController.HandleGetRole))).Methods("GET")
	router.Handle("/role/update-role", middlewares.AuthMiddleware(makeHTTPHandleFunc(roleController.HandleUpdateRole))).Methods("PUT")
	router.Handle("/role/get-all-role", makeHTTPHandleFunc(roleController.HandleGetAllRoles)).Methods("GET")
}

func (s *APIServer) PermissionRouter(router *mux.Router){
	router.Handle("/permission/create-permission", middlewares.AuthMiddleware(makeHTTPHandleFunc(permissionController.HandleCreatePermissions))).Methods("POST")
	router.Handle("/permission/delete-permission", middlewares.AuthMiddleware(makeHTTPHandleFunc(permissionController.HandleDeletePermissions))).Methods("DELETE")
	router.Handle("/permission/get-permission", middlewares.AuthMiddleware(makeHTTPHandleFunc(permissionController.HandleGetPermissions))).Methods("GET")
	router.Handle("/permission/update-permission", middlewares.AuthMiddleware(makeHTTPHandleFunc(permissionController.HandleUpdatePermissions))).Methods("PUT")
	router.Handle("/permission/get-all-permission", makeHTTPHandleFunc(permissionController.HandleGetAllPermissions)).Methods("GET")
}


func (s *APIServer) OrganizationRouter(router *mux.Router) {
	router.HandleFunc("/org/signup", makeHTTPHandleFunc(authControllers.HandleCreateAccount)).Methods("POST")
	// router.HandleFunc("/org/get/{id}", makeHTTPHandleFunc(authControllers.HandleGetAccount)).Methods("GET")
	// router.HandleFunc("/org/login", makeHTTPHandleFunc(authControllers.HandleLogin)).Methods("POST")
	router.Handle("/org/current-user", middlewares.AuthMiddleware(makeHTTPHandleFunc(authControllers.HandleCurrentUser))).Methods("GET")
}

func (s *APIServer) PublicRouter(router *mux.Router) {

	router.HandleFunc("/endpoints", func(w http.ResponseWriter, r *http.Request) {
		endpoints := map[string]map[string][]string{
			"auth": {
				"POST": {"signup", "login"},
				"GET":  {"current-user", "get/{id}"},
			},
			"org": {
				"GET": {"current-user"},
			},
			"ws": {
				"GET": {"ws"},
			},
		}
		lib.WriteJSON(w, http.StatusOK, endpoints)
	}).Methods("GET")
}

