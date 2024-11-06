package main

import (
	"net/http"

	"github.com/gorilla/mux"
	authControllers "github.com/khalidkhnz/2D-metaverse-app/backend/controllers/auth"
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

