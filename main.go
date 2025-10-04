package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	host := gin.Default()
	host.GET("/", func(c *gin.Context) {
		log.Println("Hello World")
		c.IndentedJSON(200, gin.H{
			"message": "Hello World",
		})
	})
	host.Run(":8083")
	log.Println("Server started on port 8083")
}
