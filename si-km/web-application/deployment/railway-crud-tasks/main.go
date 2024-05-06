package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     int    `form:"id" json:"id"`
	Title  string `form:"title" json:"title"`
	Status bool   `form:"status" json:"status"`
}

var (
	tasks  = []Task{}
	lastID = 0
	mu     sync.Mutex
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.POST("/tasks/delete/:id", deleteTask)

	router.Run(":8080")
}

func getTasks(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", tasks)
}

func createTask(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var task Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	lastID++
	task.ID = lastID
	tasks = append(tasks, task)

	c.Redirect(http.StatusMovedPermanently, "/tasks")
}

func updateTask(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var task Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			tasks[i].Title = task.Title
			tasks[i].Status = task.Status
			c.Redirect(http.StatusMovedPermanently, "/tasks")
			return
		}
	}

	c.JSON(404, gin.H{"error": "Task not found"})
}

func deleteTask(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.Redirect(http.StatusMovedPermanently, "/tasks")
			return
		}
	}

	c.JSON(404, gin.H{"error": "Task not found"})
}
