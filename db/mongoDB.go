package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	opts := options.Client().ApplyURI("mongodb+srv://rai_98:mongo123456@crudcomportamos.9rlbg78.mongodb.net/?retryWrites=true&w=majority&appName=crudComportamos").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	return mongo.Connect(context.TODO(), opts)
}
