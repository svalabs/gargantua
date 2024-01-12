package mongodbsvc

import (
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

type MongoDBServer struct {
}

func NewMongoDBServer() (MongoDBServer, error) {
	a := MongoDBServer{}
	return a, nil
}

func (a MongoDBServer) SetupRoutes(r *mux.Router) {
	path := "/api/items" //todo set path
	r.HandleFunc(path, a.getItems).Methods("GET")
	r.HandleFunc(path+"/{id}", a.getItem).Methods("GET")
	r.HandleFunc(path, a.createItem).Methods("POST")
	r.HandleFunc(path+"/{id}", a.updateItem).Methods("PUT")
	r.HandleFunc(path+"/{id}", a.deleteItem).Methods("DELETE")
	r.HandleFunc(path+"/{id}", a.appendDataToItem).Methods("PATCH")
	glog.V(2).Infof("set up route")
}
