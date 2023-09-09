package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"github.com/hobbyfarm/gargantua/pkg/util"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	Id      string   `bson:"_id" json:"id"`
	History []string `bson:"history" json:"history"`
}

var (
	collection *mongo.Collection

	URI            string
	dbName         string
	collectionName string
	servicePort    int
)

func init() {
	flag.StringVar(&URI, "mongoURI", "mongodb://mongo-db-service.hobbyfarm.svc.cluster.local:27017", "URI of the mongodb")
	flag.StringVar(&dbName, "dbName", "hobbyfarmDB", "Name of the mongodb")
	flag.StringVar(&collectionName, "collection", "sessionData", "Name of the collection in the database")
	flag.IntVar(&servicePort, "servicePort", 8080, "Port to run service on")

	// Use env if available
	if envURI := os.Getenv("DB_URI"); envURI != "" {
		URI = envURI
	}
	if envName := os.Getenv("DB_NAME"); envName != "" {
		dbName = envName
	}
	if envCollection := os.Getenv("DB_COLLECTION"); envCollection != "" {
		collectionName = envCollection
	}
}

func main() {
	flag.Parse()
	r := mux.NewRouter()
	SetupRoutes(r)

	credentials := options.Credential{
		AuthSource:    "admin",
		AuthMechanism: "SCRAM-SHA-256",
		Username:      "adminuser",
		Password:      "password123",
	}

	// Create MongoDB client
	clientOptions := options.Client().ApplyURI(URI).SetAuth(credentials)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		glog.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Select MongoDB collection
	collection = client.Database(dbName).Collection(collectionName)

	// Start the server
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(servicePort),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	glog.Info("Starting server on port ", servicePort)

	if err := srv.ListenAndServe(); err != nil {
		glog.Fatal(err)
	}
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/items", getItems).Methods("GET")
	r.HandleFunc("/api/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/api/items", createItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", deleteItem).Methods("DELETE")
	glog.V(2).Infof("API Routes are set up")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received getItems")
	// Fetch all items from the MongoDB collection
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	defer cur.Close(context.Background())

	var items []Item
	for cur.Next(context.Background()) {
		var item Item
		if err := cur.Decode(&item); err != nil {
			util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		items = append(items, item)
	}
	glog.Info(items)
	content, err := json.Marshal(items)
	glog.Info(content)
	if err != nil {
		glog.Error(err)
	}
	util.ReturnHTTPContent(w, r, http.StatusOK, "content", content)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received getItem")
	// Extract the item ID from the URL path parameter
	id := mux.Vars(r)["id"]

	// Find the item with the given ID in the MongoDB collection
	filter := bson.M{"_id": id}
	result := collection.FindOne(context.Background(), filter)

	var item Item
	if err := result.Decode(&item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	content, err := json.Marshal(item)
	if err != nil {
		glog.Error(err)
	}
	util.ReturnHTTPContent(w, r, http.StatusOK, "content", content)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received createItem")
	// Parse the request body into an Item struct
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	glog.Info(item)
	// Insert the item into the MongoDB collection
	if _, err := collection.InsertOne(context.Background(), item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	content, err := json.Marshal(item)
	if err != nil {
		glog.Error(err)
	}
	glog.Info(content)
	util.ReturnHTTPContent(w, r, http.StatusOK, "content", content)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into an Item struct
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update the item in the MongoDB collection
	filter := bson.M{"_id": item.Id}
	update := bson.M{"history": item.History}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if result.MatchedCount == 0 {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	content, err := json.Marshal(item)
	if err != nil {
		glog.Error(err)
	}
	util.ReturnHTTPContent(w, r, http.StatusOK, "content", content)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path parameter
	id := mux.Vars(r)["id"]

	// Delete the item from the MongoDB collection
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if result.DeletedCount == 0 {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	//todo check messageType
	util.ReturnHTTPMessage(w, r, http.StatusOK, "success", "Item deleted successfully")
}
