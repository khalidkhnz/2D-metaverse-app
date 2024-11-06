package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	// MUX ROUTER
	router := mux.NewRouter()

	// CUSTOM LOGGING MIDDLEWARE
	router.Use(middlewares.LoggingMiddleware)
		
	// API VER. PREFIX
	apiRouter := router.PathPrefix(lib.ApiPrefix).Subrouter()
	
	// WEBSOCKET CONN "/api/v1"
	router.HandleFunc("/ws", Handler).Methods("GET")
	
	// API ROUTERS
	s.PublicRouter(apiRouter)
	s.AuthRouter(apiRouter)
	s.OrganizationRouter(apiRouter)
	
	// PROXY SERVER "/"
	s.ProxyServer(lib.FrontEndProxyURL,router)

	// FILE SERVER "/{api-prefix}/file-server"
	s.FileServer("./views","/file-server",apiRouter)
	
	// NOT FOUND HANDLE "*"
	router.HandleFunc("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		lib.WriteJSON(w, http.StatusNotFound, map[string]any{
			"success": false,
			"message": "Endpoint does not exist on the server",
		})
	})

	log.Println("API SERVER RUNNING ON PORT", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)
	if err!= nil {
		log.Fatal(err)
	}
}


