package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yaghoubi-mn/voter/internal/handlers"
)

func SetupRoutes(r *gin.Engine, handlers handlers.UserHandler) {

	// middlewares
	// r.Use(middleware.Auth())

	v1 := r.Group("/api/v1")

	v1.POST("/users/login", handlers.Login)
}
