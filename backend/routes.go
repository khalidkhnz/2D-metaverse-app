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
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(middlewares.AuthMiddleware)

	router.HandleFunc("/auth/signup", makeHTTPHandleFunc(authControllers.HandleCreateAccount)).Methods("POST")
	router.HandleFunc("/auth/get/{id}", makeHTTPHandleFunc(authControllers.HandleGetAccount)).Methods("GET")
	router.HandleFunc("/auth/login", makeHTTPHandleFunc(authControllers.HandleLogin)).Methods("POST")
	// REQUIRES TOKEN
	authRouter.Handle("/current-user", makeHTTPHandleFunc(authControllers.HandleCurrentUser)).Methods("GET")
	authRouter.Handle("/ws-token", makeHTTPHandleFunc(authControllers.HandleGenerateShortLivedJwtTokenForSocket)).Methods("GET")
}

func (s *APIServer) RoleRouter(router *mux.Router){
	authRouter := router.PathPrefix("/role").Subrouter()
	authRouter.Use(middlewares.AuthMiddleware)
	authRouter.Use(func(handler http.Handler) http.Handler {
		return middlewares.RoleCheckerMiddleware(handler, []string{"SUPER_ADMIN"})
	})

	authRouter.Handle("/create-role", makeHTTPHandleFunc(roleController.HandleCreateRole)).Methods("POST")
	authRouter.Handle("/delete-role", makeHTTPHandleFunc(roleController.HandleDeleteRole)).Methods("DELETE")
	authRouter.Handle("/get-role", makeHTTPHandleFunc(roleController.HandleGetRole)).Methods("GET")
	authRouter.Handle("/update-role", makeHTTPHandleFunc(roleController.HandleUpdateRole)).Methods("PUT")
	router.Handle("/role/get-all-role", makeHTTPHandleFunc(roleController.HandleGetAllRoles)).Methods("GET")
}

func (s *APIServer) PermissionRouter(router *mux.Router){
	authRouter := router.PathPrefix("/permission").Subrouter()
	authRouter.Use(middlewares.AuthMiddleware)
	authRouter.Use(func(handler http.Handler) http.Handler {
		return middlewares.RoleCheckerMiddleware(handler, []string{"SUPER_ADMIN"})
	})

	authRouter.Handle("/create-permission", makeHTTPHandleFunc(permissionController.HandleCreatePermissions)).Methods("POST")
	authRouter.Handle("/delete-permission", makeHTTPHandleFunc(permissionController.HandleDeletePermissions)).Methods("DELETE")
	authRouter.Handle("/get-permission", makeHTTPHandleFunc(permissionController.HandleGetPermissions)).Methods("GET")
	authRouter.Handle("/update-permission", makeHTTPHandleFunc(permissionController.HandleUpdatePermissions)).Methods("PUT")
	router.Handle("/permission/get-all-permission", makeHTTPHandleFunc(permissionController.HandleGetAllPermissions)).Methods("GET")
}


func (s *APIServer) OrganizationRouter(router *mux.Router) {
	// router.HandleFunc("/org/signup", makeHTTPHandleFunc(authControllers.HandleCreateAccount)).Methods("POST")
	// router.Handle("/org/current-user", middlewares.AuthMiddleware(makeHTTPHandleFunc(authControllers.HandleCurrentUser))).Methods("GET")
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

