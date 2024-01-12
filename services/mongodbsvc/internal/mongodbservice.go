package mongodbsvc

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/hobbyfarm/gargantua/v3/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"reflect"
)

type Item struct {
	Id   string                 `bson:"_id" json:"id"`
	Data map[string]interface{} `bson:"data" json:"data"`
}

var (
	collection *mongo.Collection
)

const (
	NOTFOUND       = "Item not found"
	INVALIDPAYLOAD = "Invalid Request Payload"
)

func (a MongoDBServer) getItems(w http.ResponseWriter, r *http.Request) {
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

func (a MongoDBServer) getItem(w http.ResponseWriter, r *http.Request) {
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

func (a MongoDBServer) createItem(w http.ResponseWriter, r *http.Request) {
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

func (a MongoDBServer) updateItem(w http.ResponseWriter, r *http.Request) {
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

func (a MongoDBServer) appendDataToItem(w http.ResponseWriter, r *http.Request) {
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

func (a MongoDBServer) deleteItem(w http.ResponseWriter, r *http.Request) {
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
					return fmt.Errorf("type mismatch for key '%s': existing value is a map, new value is of type %T", key, newValue)
				}
			default:
				if reflect.TypeOf(existingValue).Kind() == reflect.Slice && reflect.TypeOf(newValue).Kind() == reflect.Slice {
					// Convert both values to []interface{} and append
					existingValSlice := normalizeSlice(existingValue)
					newValSlice := normalizeSlice(newValue)
					currentData[key] = append(existingValSlice, newValSlice...)
				} else if reflect.TypeOf(existingValue) != reflect.TypeOf(newValue) {
					return fmt.Errorf("type mismatch for key '%s': existing value is of type %T, new value is of type %T", key, existingValue, newValue)
				} else {
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
