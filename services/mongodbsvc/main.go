package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/hobbyfarm/gargantua/v3/pkg/microservices"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	serviceConfig  *microservices.ServiceConfig
	URI            string
	dbName         string
	collectionName string
	servicePort    int

	username string
	password string
)

func init() { //todo move flags to config
	serviceConfig = microservices.BuildServiceConfig()
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

func main() { //todo: add new microservice setup structure
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
