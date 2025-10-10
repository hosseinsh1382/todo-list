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

func (h *TaskHandler) RegisterRoutes(router *gin.Engine) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("", h.Get)
		tasks.GET("/:id", h.GetById)
		tasks.POST("", h.Add)
		tasks.DELETE("/:id", h.Delete)
	}
}

// Get godoc
// @Summary		list tasks
// @Description	Get all tasks
// @Tags		tasks
// @Produce 	json
// @Success 	200 {array} Models.Task
// @Router		/tasks [get]
func (t *TaskHandler) Get(c *gin.Context) {
	tasks, err := t.TaskService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		t.ErrorLogger.Fatalf("Error in GetAll:\n %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
	return
}

// GetById 	godoc
// @Summary 	get by ID
// @Description Get a task by passing an ID
// @Tags		tasks
// @Produce		json
// @Success		200 {object} Models.Task
// @Failure		500 {object} map[string]string
// @Router		/tasks/{id} [get]
// @Param 		id	path	int	true	"id"
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

// Add godoc
// @Summary 	add new task
// @Description	add new task to database
// @Param		task body Models.Task true "new task" default
// @Tags		tasks
// @Router		/tasks [post]
func (h *TaskHandler) Add(c *gin.Context) {
	var task Models.Task
	if err := c.BindJSON(&task); err != nil {
		h.ErrorLogger.Printf("Error in Bind:\n %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
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

// Delete godoc
// @Summary		delete task
// @Description	delete a task by given ID
// @Param 		id path int true "id"
// @Tags 		tasks
// @Router		/tasks/{id} [delete]
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
