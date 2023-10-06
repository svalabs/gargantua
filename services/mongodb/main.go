package main

import (
	"context"
	"encoding/json"
	"github.com/hobbyfarm/gargantua/pkg/util"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI            string
	dbName         string
	collectionName string
	servicePort    int
	username       string
	password       string
}

type Server struct {
	config     Config
	collection *mongo.Collection
}

func NewServer(config Config) *Server {
	return &Server{config: config}
}

func (s *Server) initMongoDB() error {
	credentials := options.Credential{
		AuthSource:    "admin",
		AuthMechanism: "SCRAM-SHA-256",
		Username:      s.config.username,
		Password:      s.config.password,
	}

	clientOptions := options.Client().ApplyURI(s.config.URI).SetAuth(credentials)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	s.collection = client.Database(s.config.dbName).Collection(s.config.collectionName)
	return nil
}

func (s *Server) start() error {
	r := mux.NewRouter()
	s.setupRoutes(r)

	// Start the server
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(s.config.servicePort),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	glog.Info("Starting server on port ", s.config.servicePort)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) setupRoutes(r *mux.Router) {
	r.HandleFunc("/api/items", s.getItems).Methods("GET")
	r.HandleFunc("/api/items/{id}", s.getItem).Methods("GET")
	r.HandleFunc("/api/items", s.createItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", s.updateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", s.deleteItem).Methods("DELETE")
	r.HandleFunc("/api/items/{id}", s.appendDataToItem).Methods("PATCH")
	glog.V(2).Infof("API Routes are set up")
}

type Item struct {
	Id   string                 `bson:"_id" json:"id"`
	Data map[string]interface{} `bson:"data" json:"data"`
}

func (s *Server) getItems(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received getItems")
	// Fetch all items from the MongoDB collection
	cur, err := s.collection.Find(context.Background(), bson.D{})
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
	returnJSONResponse(w, r, http.StatusOK, items)
}

func (s *Server) getItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received getItem")
	// Extract the item ID from the URL path parameter
	id := mux.Vars(r)["id"]

	// Find the item with the given ID in the MongoDB collection
	filter := bson.M{"_id": id}
	result := s.collection.FindOne(context.Background(), filter)

	var item Item
	if err := result.Decode(&item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	returnJSONResponse(w, r, http.StatusOK, item)
}

func (s *Server) createItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received createItem")

	// Parse the request body into an Item struct
	var item Item
	if err := decodeJSONRequest(r, &item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Insert the item into the MongoDB collection
	if _, err := s.collection.InsertOne(context.Background(), item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	returnJSONResponse(w, r, http.StatusOK, item)
}

func (s *Server) updateItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var item Item

	if err := decodeJSONRequest(r, &item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update the item in the MongoDB collection
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"data": item.Data}}
	result, err := s.collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if result.MatchedCount == 0 {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	returnJSONResponse(w, r, http.StatusOK, item)
}

func (s *Server) appendDataToItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var requestData map[string]interface{}

	if err := decodeJSONRequest(r, &requestData); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Use $addToSet to append data to the existing "data" field
	update := bson.M{"$addToSet": bson.M{"data": requestData}}
	filter := bson.M{"_id": id}
	result, err := s.collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if result.MatchedCount == 0 {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	returnJSONResponse(w, r, http.StatusOK, requestData)
}

func (s *Server) deleteItem(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path parameter
	id := mux.Vars(r)["id"]

	// Delete the item from the MongoDB collection
	filter := bson.M{"_id": id}
	result, err := s.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if result.DeletedCount == 0 {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	util.ReturnHTTPMessage(w, r, http.StatusOK, "success", "Item deleted successfully")
}

// Parse the request body into an Item struct
func decodeJSONRequest(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// Return a JSON response

func returnJSONResponse(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		glog.Error(err)
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, "Failed to marshal response")
	}
}
