package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/hobbyfarm/gargantua/pkg/util"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	Id   string                 `bson:"_id" json:"id"`
	Data map[string]interface{} `bson:"data" json:"data"`
}

var (
	collection *mongo.Collection

	URI            string
	dbName         string
	collectionName string
	servicePort    int

	username string
	password string
)

func init() {
	flag.StringVar(&URI, "mongoURI", "mongodb://mongo-db-service.hobbyfarm.svc.cluster.local:27017", "URI of the mongodb")
	flag.StringVar(&dbName, "dbName", "hobbyfarmDB", "Name of the mongodb")
	flag.StringVar(&collectionName, "collection", "sessionData", "Name of the collection in the database")
	flag.IntVar(&servicePort, "servicePort", 8080, "Port to run service on")

	flag.StringVar(&username, "username", "", "Username for MongoDB authentication.")
	flag.StringVar(&password, "password", "", "Password for MongoDB authentication.")

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

	if envUsername := os.Getenv("DB_USERNAME"); envUsername != "" {
		username = envUsername
	}
	if envPassword := os.Getenv("DB_PASSWORD"); envPassword != "" {
		password = envPassword
	}
}

func main() {
	flag.Parse()
	r := mux.NewRouter()
	SetupRoutes(r)

	credentials := options.Credential{
		AuthSource:    "admin",
		AuthMechanism: "SCRAM-SHA-256",
		Username:      username,
		Password:      password,
	}

	// Create MongoDB client
	clientOptions := options.Client().ApplyURI(URI).SetAuth(credentials)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		glog.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			glog.Error(err)
		}
	}(client, context.Background())

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
	path := "/api/items"
	r.HandleFunc(path, getItems).Methods("GET")
	r.HandleFunc(path+"/{id}", getItem).Methods("GET")
	r.HandleFunc(path, createItem).Methods("POST")
	r.HandleFunc(path+"/{id}", updateItem).Methods("PUT")
	r.HandleFunc(path+"/{id}", deleteItem).Methods("DELETE")
	r.HandleFunc(path+"/{id}", appendDataToItem).Methods("PATCH")
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
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			glog.Error(err)
		}
	}(cur, context.Background())

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
	util.ReturnJSONResponse(w, r, http.StatusOK, items)
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

	util.ReturnJSONResponse(w, r, http.StatusOK, item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received createItem")

	// Parse the request body into an Item struct
	var item Item
	if err := util.DecodeJSONRequest(r, &item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Insert the item into the MongoDB collection
	if _, err := collection.InsertOne(context.Background(), item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	util.ReturnJSONResponse(w, r, http.StatusOK, item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// Parse the request body into an Item struct
	var item Item
	if err := util.DecodeJSONRequest(r, &item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update the item in the MongoDB collection
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"data": item.Data}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if result.MatchedCount == 0 {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	util.ReturnJSONResponse(w, r, http.StatusOK, item)
}

func appendDataToItem(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path parameter
	id := mux.Vars(r)["id"]

	// Parse the request body into a map[string]interface{} to get the data to append
	var requestData map[string]interface{}
	if err := util.DecodeJSONRequest(r, &requestData); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Retrieve the current value of the 'data' field
	var currentItem Item
	filter := bson.M{"_id": id}
	if err := collection.FindOne(context.Background(), filter).Decode(&currentItem); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, "Item not found")
		return
	}

	// Ensure 'Data' is a map[string]interface{}
	if currentItem.Data == nil {
		currentItem.Data = make(map[string]interface{})
	}

	// Merge the new data into the current data object
	if err := mergeJSON(currentItem.Data, requestData); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Update the 'data' field with the merged value
	update := bson.M{"$set": bson.M{"data": currentItem.Data}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	util.ReturnJSONResponse(w, r, http.StatusOK, currentItem.Data)
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
	util.ReturnHTTPMessage(w, r, http.StatusOK, "success", "Item deleted successfully")
}

// Merge two JSON objects
func mergeJSON(currentData map[string]interface{}, newData map[string]interface{}) error {
	for key, value := range newData {
		if existingValue, exists := currentData[key]; exists {
			// Normalize array types
			existingValArr, existingIsArray := existingValue.([]interface{})
			newValArr, newIsArray := value.([]interface{})
			if !existingIsArray && reflect.TypeOf(existingValue).Kind() == reflect.Slice {
				existingValArr = normalizeSlice(existingValue)
				existingIsArray = true
			}
			if !newIsArray && reflect.TypeOf(value).Kind() == reflect.Slice {
				newValArr = normalizeSlice(value)
				newIsArray = true
			}

			// Check if both are maps
			if existingMap, isMap := existingValue.(map[string]interface{}); isMap {
				if newMap, isNewMap := value.(map[string]interface{}); isNewMap {
					if err := mergeJSON(existingMap, newMap); err != nil {
						return err
					}
				} else {
					return fmt.Errorf("Key '%s' already exists and is an object.", key)
				}
			} else if existingIsArray && newIsArray { // Check if both are arrays
				existingValArr = append(existingValArr, newValArr...)
				currentData[key] = existingValArr
			} else {
				// Check if the types are compatible
				if reflect.TypeOf(existingValue) != reflect.TypeOf(value) {
					return fmt.Errorf("Type mismatch for key '%s'.", key)
				}
				// Throw an error if attempting to overwrite an object
				return fmt.Errorf("Key '%s' already exists and is not an array or object.", key)
			}
		} else {
			// If the key doesn't exist in the current data, add it
			currentData[key] = value
		}
	}
	return nil
}

// convert slice into a slice of the empty interface type
func normalizeSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}
	result := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}
	return result
}
