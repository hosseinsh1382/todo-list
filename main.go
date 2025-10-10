package main

import (
	"ToDoList/handlers"
	"ToDoList/repositories"
	"ToDoList/services"
	"log"

	_ "ToDoList/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	host := gin.Default()
	host.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	host.GET("/", func(c *gin.Context) {
		log.Println("Hello World")
		c.IndentedJSON(200, gin.H{
			"message": "Connection Established",
		})
	})

	taskRepository, err := repositories.NewPostgresTaskRepository()
	if err != nil {
		log.Fatalf("failed to connect to repository: %v", err)
	}
	taskService := services.NewDefaultTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	taskHandler.RegisterRoutes(host)

	log.Println("Server started on port 8083")
	log.Fatal(host.Run(":8083"))
}
