package service

import (
	"context"
	"crud-compartamos-api/db"
	"crud-compartamos-api/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IServiceClient interface {
	SaveClient(client types.Client) (bool, error)
	FinAllClients() ([]types.Client, error)
	UpdateClient(client types.Client) ([]types.Client, error)
	DeleteClient(client types.Client) ([]types.Client, error)
}

type ClientService struct {
	Clients IServiceClient
	mongo   db.IConnectMongo
}

func New(m db.IConnectMongo) IServiceClient {
	return &ClientService{
		mongo: m,
	}
}

func (s *ClientService) SaveClient(client types.Client) (bool, error) {
	collection, err := s.getCollection()
	if err != nil {
		return false, err
	}
	result, err := collection.InsertOne(context.TODO(), client)
	if err != nil {
		return false, err
	}
	// // ==== Create unit Index =======
	//indexName, err := collection.Indexes().CreateOne(
	//	context.Background(),
	//	mongo.IndexModel{
	//		Keys:    bson.D{{Key: "dni", Value: 1}},
	//		Options: options.Index().SetUnique(true),
	//	},
	//)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return true, nil
}

func (s *ClientService) FinAllClients() ([]types.Client, error) {
	collection, err := s.getCollection()
	if err != nil {
		return nil, err
	}

	filter := bson.D{}
	result, err := collection.Find(context.TODO(), filter)

	var clients []types.Client

	if err = result.All(context.TODO(), &clients); err != nil {
		return nil, err
	}
	return clients, nil
}

func (s *ClientService) UpdateClient(clientUpdate types.Client) ([]types.Client, error) {
	collection, err := s.getCollection()
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"dni", clientUpdate.DNI}}
	update := bson.D{{"$set", clientUpdate}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	return s.FinAllClients()
}

func (s *ClientService) DeleteClient(client types.Client) ([]types.Client, error) {
	collection, err := s.getCollection()
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"dni", client.DNI}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Document delete: %v\n", result)
	return s.FinAllClients()
}

func (s *ClientService) getCollection() (*mongo.Collection, error) {
	mongoClient, err := s.mongo.ConnectMongoDb()
	if err != nil {
		return nil, err
	}

	return mongoClient.Database("compartamos_clients").Collection("clients"), nil

}
