package handlers

import (
	"go-swagger-example/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      ABABBABABA
// @Tags         todos
// @Produce      json
// @Success      200  {array}  models.Todo
// @Router       /todos [get]
func GetTodos(c *gin.Context) {
	// Giả sử lấy dữ liệu từ đâu đó
	todos := []models.Todo{{ID: "1", Task: "Học cách chia file", Completed: false}}
	c.JSON(http.StatusOK, todos)
}
