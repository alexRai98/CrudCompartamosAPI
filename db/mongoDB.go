package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type IConnectMongo interface {
	ConnectMongoDb() (*mongo.Client, error)
}

type ConnectionMongo struct{}

func NewConnection() IConnectMongo {
	return &ConnectionMongo{}
}

func (c *ConnectionMongo) ConnectMongoDb() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	credential := fmt.Sprintf("mongodb+srv://rai_98:%s@crudcomportamos.9rlbg78.mongodb.net/?retryWrites=true&w=majority&appName=crudComportamos", os.Getenv("PSS_DB"))
	opts := options.Client().ApplyURI(credential).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	return mongo.Connect(context.TODO(), opts)
}
