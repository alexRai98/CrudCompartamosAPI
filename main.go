package main

import (
	"crud-compartamos-api/controller"
	"crud-compartamos-api/db"
	"crud-compartamos-api/service"
	"crud-compartamos-api/validator"
	"github.com/gin-gonic/gin"
)

var (
	mongoDB          = db.NewConnection()
	clientService    = service.New(mongoDB)
	validate         = validator.NewValidator()
	clientController = controller.New(clientService, validate)
)

func main() {
	server := gin.Default()
	server.GET("/", func(context *gin.Context) {
		context.JSON(200, "Hello ..!!")
	})
	server.GET("/clients", func(context *gin.Context) {
		clients, err := clientController.FindAllClients()
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		context.JSON(200, clients)
	})
	server.POST("/client", func(context *gin.Context) {
		response, err := clientController.SaveClient(context)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		context.JSON(201, gin.H{
			"ok": response,
		})
	})
	server.PUT("/client", func(context *gin.Context) {
		clients, err := clientController.UpdateClient(context)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		context.JSON(200, clients)
	})
	server.DELETE("/client", func(context *gin.Context) {
		clients, err := clientController.DeleteClient(context)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		context.JSON(200, clients)
	})
	server.Run()
}
