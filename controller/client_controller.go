package controller

import (
	"crud-compartamos-api/service"
	"crud-compartamos-api/types"
	"crud-compartamos-api/validator"
	"github.com/gin-gonic/gin"
)

type IClientController interface {
	SaveClient(ctx *gin.Context) (bool, error)
	FindAllClients() ([]types.Client, error)
	UpdateClient(ctx *gin.Context) ([]types.Client, error)
	DeleteClient(ctx *gin.Context) ([]types.Client, error)
}

type Controller struct {
	service   service.IServiceClient
	validator validator.IValidator
}

func New(service service.IServiceClient, v validator.IValidator) IClientController {
	return Controller{
		service:   service,
		validator: v,
	}
}

func (c Controller) SaveClient(ctx *gin.Context) (bool, error) {
	var client types.Client

	if err := ctx.BindJSON(&client); err != nil {
		return false, err
	}

	if err := c.validator.ValidateRequest(client); err != nil {
		return false, err
	}

	return c.service.SaveClient(client)

}

func (c Controller) FindAllClients() ([]types.Client, error) {
	return c.service.FinAllClients()
}

func (c Controller) UpdateClient(ctx *gin.Context) ([]types.Client, error) {
	var client types.Client
	if err := ctx.BindJSON(&client); err != nil {
		return nil, err
	}

	if err := c.validator.ValidateRequest(client); err != nil {
		return nil, err
	}
	return c.service.UpdateClient(client)
}

func (c Controller) DeleteClient(ctx *gin.Context) ([]types.Client, error) {
	var client types.Client
	if err := ctx.BindJSON(&client); err != nil {
		return nil, err
	}
	return c.service.DeleteClient(client)
}
