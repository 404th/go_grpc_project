package main

import (
	"github.com/404th/go_grpc_project/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Declaring GIN project
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	// ROUTES
	engine.GET("/ping", handlers.Ping)

	r := engine.Group("/book")
	{
		r.GET("", handlers.GetBookList)
		r.GET("/:id", handlers.GetBookByID)
		r.POST("", handlers.CreateBook)
		r.PUT("/:id", handlers.UpdateBook)
		r.DELETE("/:id", handlers.DeleteBook)
	}

	// listening port
	engine.Run(":4040")
}
