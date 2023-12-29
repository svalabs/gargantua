package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/hobbyfarm/gargantua/v3/pkg/util"
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

const (
	NOTFOUND       = "Item not found"
	INVALIDPAYLOAD = "Invalid Request Payload"
)

// get parameters from flags for the mongoDB connection
func init() {
	flag.StringVar(&URI, "mongoURI", "", "URI of the mongodb")
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

	// Iterate over each document and put it into an Item struct
	var items []Item
	for cur.Next(context.Background()) {
		var item Item
		if err := cur.Decode(&item); err != nil {
			util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		items = append(items, item)
	}
	if err != nil {
		glog.Error(err)
	}
	util.ReturnJSONResponse(w, r, http.StatusOK, items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received getItem")

	// Extract ID from path parameter and find it in the db
	id := mux.Vars(r)["id"]
	filter := bson.M{"_id": id}
	result := collection.FindOne(context.Background(), filter)

	var item Item
	if err := result.Decode(&item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusNotFound, NOTFOUND)
		return
	}

	util.ReturnJSONResponse(w, r, http.StatusOK, item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received createItem")

	// Parse body into an Item struct
	var item Item
	if err := util.DecodeJSONRequest(r, &item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, INVALIDPAYLOAD)
		return
	}

	// Ensure collection exists to avoid transaction error
	if err := ensureCollectionExists(collection); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	session, err := collection.Database().Client().StartSession()
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}

		// Insert the item into the DB
		if _, err := collection.InsertOne(sc, item); err != nil {
			session.AbortTransaction(sc)
			return err
		}

		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	util.ReturnJSONResponse(w, r, http.StatusOK, item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received updateItem")

	id := mux.Vars(r)["id"]
	var item Item
	if err := util.DecodeJSONRequest(r, &item); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, INVALIDPAYLOAD)
		return
	}

	session, err := collection.Database().Client().StartSession()
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err
		}

		// Update the item in the MongoDB collection
		filter := bson.M{"_id": id}
		update := bson.M{"$set": bson.M{"data": item.Data}}
		result, err := collection.UpdateOne(sc, filter, update)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		if result.MatchedCount == 0 {
			session.AbortTransaction(sc)
			return fmt.Errorf(NOTFOUND)
		}

		if err := session.CommitTransaction(sc); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	util.ReturnJSONResponse(w, r, http.StatusOK, item)
}

func appendDataToItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received appendDataToItem")
	id := mux.Vars(r)["id"]

	var requestData map[string]interface{}
	var updatedData map[string]interface{}

	// Parse the request body into a map[string]interface{} to get the data to append
	if err := util.DecodeJSONRequest(r, &requestData); err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusBadRequest, INVALIDPAYLOAD)
		return
	}

	session, err := collection.Database().Client().StartSession()
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err
		}

		// Retrieve the current value of the "data" field
		var currentItem Item
		filter := bson.M{"_id": id}
		if err := collection.FindOne(sc, filter).Decode(&currentItem); err != nil {
			session.AbortTransaction(sc)
			return fmt.Errorf(NOTFOUND)
		}

		// Ensure "Data" is a map[string]interface{}
		if currentItem.Data == nil {
			currentItem.Data = make(map[string]interface{})
		}

		// Merge the new data into the current data object
		if err := mergeJSON(currentItem.Data, requestData); err != nil {
			session.AbortTransaction(sc)
			return err
		}

		// Update the "data" field with the merged value
		update := bson.M{"$set": bson.M{"data": currentItem.Data}}
		_, err = collection.UpdateOne(sc, filter, update)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		if err := session.CommitTransaction(sc); err != nil {
			return err
		}
		updatedData = currentItem.Data
		return nil
	})

	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	util.ReturnJSONResponse(w, r, http.StatusOK, updatedData)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	glog.Info("Received deleteItem")

	id := mux.Vars(r)["id"]
	session, err := collection.Database().Client().StartSession()
	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err
		}

		filter := bson.M{"_id": id}
		result, err := collection.DeleteOne(sc, filter)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		if result.DeletedCount == 0 {
			session.AbortTransaction(sc)
			return fmt.Errorf(NOTFOUND)
		}

		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		util.ReturnHTTPErrorMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	util.ReturnHTTPMessage(w, r, http.StatusOK, "success", "Item deleted successfully")
}

// Merge two JSON objects
func mergeJSON(currentData map[string]interface{}, newData map[string]interface{}) error {
	for key, newValue := range newData {
		if existingValue, exists := currentData[key]; exists {
			switch existingValTyped := existingValue.(type) {
			case map[string]interface{}:
				// Both values are maps, merge them
				if newMap, ok := newValue.(map[string]interface{}); ok {
					if err := mergeJSON(existingValTyped, newMap); err != nil {
						return err
					}
				} else {
					return fmt.Errorf("type mismatch for key '%s': existing value is map, new value is not", key)
				}
			default:
				// Check if both are slices
				if reflect.TypeOf(existingValue).Kind() == reflect.Slice && reflect.TypeOf(newValue).Kind() == reflect.Slice {
					// Convert both values to []interface{} and append
					existingValSlice := normalizeSlice(existingValue)
					newValSlice := normalizeSlice(newValue)
					currentData[key] = append(existingValSlice, newValSlice...)
				} else if reflect.TypeOf(existingValue) != reflect.TypeOf(newValue) {
					// Different types and not both slices
					return fmt.Errorf("type mismatch for key '%s'", key)
				} else {
					// Same types but not slices or maps, cannot merge
					return fmt.Errorf("cannot merge non-slice, non-map types for key '%s'", key)
				}
			}
		} else {
			// Key does not exist in current data, add it
			currentData[key] = newValue
		}
	}
	return nil
}

// Convert slice into a slice of the empty interface type
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

func ensureCollectionExists(collection *mongo.Collection) error {
	// Attempt to create an index, which will create the collection if it does not exist.
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{"dummyField", 1}}, // Create an index on a dummy field.
		},
	)
	return err
}
