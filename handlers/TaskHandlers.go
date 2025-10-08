package handlers

import (
	"ToDoList/Models"
	"ToDoList/interfaces"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService interfaces.TaskService
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}

func NewTaskHandler(taskService interfaces.TaskService) *TaskHandler {
	errorLogger := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "INFO  ", log.Ldate|log.Ltime|log.Lshortfile)
	return &TaskHandler{
		TaskService: taskService,
		ErrorLogger: errorLogger,
		InfoLogger:  infoLogger,
	}
}

func (t *TaskHandler) Get(c *gin.Context) {
	tasks, err := t.TaskService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		t.ErrorLogger.Fatalf("Error in GetAll:\n %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
	return
}

func (h *TaskHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.TaskService.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
	return
}

func (h *TaskHandler) Add(c *gin.Context) {
	var task Models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	result, err := h.TaskService.Add(task)
	if err != nil {
		h.ErrorLogger.Printf("Error in Add:\n %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})

}
func (h *TaskHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.TaskService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		h.ErrorLogger.Printf("Error in Delete(%v:\n %v", id, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
	return
}
