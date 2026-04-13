package handlers

import (
	"go-swagger-example/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: "1", Task: "Task 1", Completed: false},
	{ID: "2", Task: "Task 2", Completed: false},
	{ID: "3", Task: "Task 3", Completed: false},
	{ID: "4", Task: "Task 4", Completed: false},
}

// @Summary      ABABBABABA
// @Tags         todos
// @Produce      json
// @Success      200  {array}  models.Todo
// @Router       /todos [get]
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// @Summary  Tạo
// @Tags         todos
// @Accept       json
// Product        json
// @Param        input  body      models.CreateTodoInput  true  "Nội dung công việc"
// @Success      201    {object}  models.Todo
// @Router       /todos [post]
func CreateTodo(c *gin.Context) {
	var input models.CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo := models.Todo{
		ID:        strconv.Itoa(len(todos) + 1),
		Task:      input.Task,
		Completed: false,
	}

	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}
