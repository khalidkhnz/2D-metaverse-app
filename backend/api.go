package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	authControllers "github.com/khalidkhnz/2D-metaverse-app/backend/controllers/auth"
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	"github.com/khalidkhnz/2D-metaverse-app/backend/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}


func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            log.Printf("Error handling request %s: %v", r.URL.Path, err)
            lib.WriteJSON(w, http.StatusBadRequest, ApiError{Success: false, Message: err.Error()})
        }
    }
}


type APIServer struct {
	listenAddr string
	dbClient   *mongo.Client
}

func NewAPIServer(listenAddr string, mongoURI string) *APIServer {
	client, err := ConnectToMongo(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	lib.SetDBClient(client)

	return &APIServer{
		listenAddr: listenAddr,
		dbClient:   client,
	}
}

func ConnectToMongo(mongoURI string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %v", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.Use(middlewares.LoggingMiddleware)
	
	router.Handle("/",http.FileServer(http.Dir("./views")))

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	
	router.HandleFunc("/ws", Handler).Methods("GET")

	// AUTH ROUTER
	s.publicRouter(apiRouter)
	s.authRouter(apiRouter)
	s.organizationRouter(apiRouter)

	// NOT FOUND HANDLE
	router.HandleFunc("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		lib.WriteJSON(w, http.StatusNotFound, map[string]any{
			"success": false,
			"message": "Endpoint does not exist on the server",
		})
	})

	log.Println("API SERVER RUNNING ON PORT : ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}


func (s *APIServer) publicRouter(router *mux.Router) {
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


func (s *APIServer) authRouter(router *mux.Router) {
	router.HandleFunc("/auth/signup", makeHTTPHandleFunc(authControllers.HandleCreateAccount)).Methods("POST")
	router.HandleFunc("/auth/get/{id}", makeHTTPHandleFunc(authControllers.HandleGetAccount)).Methods("GET")
	router.HandleFunc("/auth/login", makeHTTPHandleFunc(authControllers.HandleLogin)).Methods("POST")
	
	// REQUIRES TOKEN
	router.Handle("/auth/current-user", middlewares.AuthMiddleware(makeHTTPHandleFunc(authControllers.HandleCurrentUser))).Methods("GET")
}


func (s *APIServer) organizationRouter(router *mux.Router) {
	router.Handle("/org/current-user", middlewares.AuthMiddleware(makeHTTPHandleFunc(authControllers.HandleCurrentUser))).Methods("GET")
}




