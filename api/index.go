package handler

import (
	"crud-compartamos-api/controller"
	"crud-compartamos-api/service"
	"crud-compartamos-api/validator"
	"github.com/gin-gonic/gin"
)

var (
	clientService    = service.New()
	validate         = validator.NewValidator()
	clientController = controller.New(clientService, validate)
)

func Handler() {
	server := gin.Default()
	server.GET("/clients", func(context *gin.Context) {
		context.JSON(200, clientController.FinAllClients())
	})
	server.POST("/client", func(context *gin.Context) {
		clients, err := clientController.SaveClient(context)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		context.JSON(201, clients)
	})
	server.Run(":8080")
}
