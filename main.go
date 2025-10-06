package main

import (
	"ToDoList/Repositories"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	repository, err := Repositories.NewPostgresTaskRepository()
	tasks, err := repository.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func main() {

	host := gin.Default()

	host.GET("/", func(c *gin.Context) {
		log.Println("Hello World")
		db, e := sql.Open("postgres", "postgres://postgres:postgres@localhost/todo?sslmode=disable")
		if e != nil {
			log.Printf("Error connecting to Postgres: %s", e)
			c.IndentedJSON(500, gin.H{"message": "Error connecting to Postgres"})
			return
		}
		e = db.Ping()
		if e != nil {
			log.Printf("Error connecting to Postgres: %s", e)
			c.IndentedJSON(500, gin.H{"message": "Error connecting to Postgres"})
			return
		}
		c.IndentedJSON(200, gin.H{
			"message": "Connection Established",
		})
	})

	host.GET("/tasks", GetAllTasks)

	log.Println("Server started on port 8083")
	log.Fatal(host.Run(":8083"))
}
