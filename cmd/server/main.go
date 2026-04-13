package main

import (
	"go-swagger-example/internal/handlers"

	_ "go-swagger-example/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Swagger Example API
// @version         1.0
// @description
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", handlers.GetTodos)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
