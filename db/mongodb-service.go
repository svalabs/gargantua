package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

func main() {
	// MongoDB connection settings
	mongoURI := "mongodb://mongodb-service:27017"
	dbName := "mydatabase"
	collectionName := "mycollection"

	// Create MongoDB client
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Select MongoDB collection
	collection = client.Database(dbName).Collection(collectionName)

	// Create router and routes
	router := mux.NewRouter()
	router.HandleFunc("/api/items", getItems).Methods("GET")
	router.HandleFunc("/api/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/api/items", createItem).Methods("POST")
	router.HandleFunc("/api/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/api/items/{id}", deleteItem).Methods("DELETE")

	// Start the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Starting server on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server gracefully stopped")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	// Implement your logic to fetch all items from the collection
	// and return the response in the desired format
}

func getItem(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path parameter
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Implement your logic to fetch the item with the given ID
	// from the collection and return the response in the desired format
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Implement your logic to create a new item based on the request body
	// and store it in the collection
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path parameter
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Implement your logic to update the item with the given ID
	// based on the request body
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path parameter
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Implement your logic to delete the item with the given ID
	// from the collection
}
