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

	taskRepository, _ := repositories.NewPostgresTaskRepository()
	taskService := services.NewDefaultTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	host.GET("/tasks", taskHandler.Get)
	host.GET("/tasks/:id", taskHandler.GetById)
	host.POST("/tasks", taskHandler.Add)
	host.DELETE("/tasks/:id", taskHandler.Delete)

	log.Println("Server started on port 8083")
	log.Fatal(host.Run(":8083"))
}
