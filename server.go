package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	server.Run(":8080")
}
