package service

import "crud-compartamos-api/types"

type IServiceClient interface {
	SaveClient(client types.Client) []types.Client
	FinAllClients() []types.Client
	UpdateClient(client types.Client) []types.Client
	DeleteClient(client types.Client) bool
}

type ClientService struct {
	Clients IServiceClient
}

var clientsFake []types.Client

func New() IServiceClient {
	return &ClientService{}
}

func (s *ClientService) SaveClient(client types.Client) []types.Client {

	clientsFake = append(clientsFake, client)
	return clientsFake
}

func (s *ClientService) FinAllClients() []types.Client {
	return clientsFake
}

func (s *ClientService) UpdateClient(clientUpdate types.Client) []types.Client {
	for index, client := range clientsFake {
		if client.DNI == clientUpdate.DNI {
			clientsFake[index] = clientUpdate
		}
	}
	return clientsFake
}

func (s *ClientService) DeleteClient(client types.Client) bool {
	return true
}
