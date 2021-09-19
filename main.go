package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// todo represents data about a task in the todo list
type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

// todo slice to seed todo list data
var todoList = []todo{
	{"1", "Learn Go"},
	{"2", "Build an API with Go"},
	{"3", "Document the API with swag"},
}

func main() {
	router := gin.Default()
	router.GET("/todo", getAllTodos)
	router.GET("/todo/:id", getTodoByID)
	router.POST("/todo", createTodo)
	router.DELETE("/todo/:id", deleteTodo)

	err := router.Run()
	if err != nil {
		os.Exit(1)
	}
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

func getTodoByID(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and return item with matching ID
	for _, todo := range todoList {
		if todo.ID == ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	// return error message if todo is not found
	c.JSON(http.StatusNotFound, gin.H{
		"message": "todo not found",
	})
}

func createTodo(c *gin.Context) {
	var newTodo todo

	// bind the received JSON data to newTodo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "an error occurred while creating todo",
		})
		return
	}

	// add the new todo item to todoList
	todoList = append(todoList, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func deleteTodo(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and delete item with matching ID
	for index, todo := range todoList {
		if todo.ID == ID {
			todoList = append(todoList[:index], todoList[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "successfully deleted todo",
			})
			return
		}
	}

	// return error message if todo is not found
	c.JSON(http.StatusNotFound, gin.H{
		"message": "todo not found",
	})
}
