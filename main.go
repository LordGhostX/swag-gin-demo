package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "swag-gin-demo/docs"
)

// todo represents data about a task in the todo list
type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

// message represents request response with a message
type message struct {
	Message string `json:"message"`
}

// todo slice to seed todo list data
var todoList = []todo{
	{"1", "Learn Go"},
	{"2", "Build an API with Go"},
	{"3", "Document the API with swag"},
}

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	// configure the Gin server
	router := gin.Default()
	router.GET("/todo", getAllTodos)
	router.GET("/todo/:id", getTodoByID)
	router.POST("/todo", createTodo)
	router.DELETE("/todo/:id", deleteTodo)

	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run the Gin server
	router.Run()
}

// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} todo
// @Router /todo [get]
func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [get]
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
	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}

// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body todo true "todo data"
// @Success 200 {object} todo
// @Failure 400 {object} message
// @Router /todo [post]
func createTodo(c *gin.Context) {
	var newTodo todo

	// bind the received JSON data to newTodo
	if err := c.BindJSON(&newTodo); err != nil {
		r := message{"an error occurred while creating todo"}
		c.JSON(http.StatusBadRequest, r)
		return
	}

	// add the new todo item to todoList
	todoList = append(todoList, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [delete]
func deleteTodo(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and delete item with matching ID
	for index, todo := range todoList {
		if todo.ID == ID {
			todoList = append(todoList[:index], todoList[index+1:]...)
			r := message{"successfully deleted todo"}
			c.JSON(http.StatusOK, r)
			return
		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}
