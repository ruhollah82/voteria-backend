package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yaghoubi-mn/voter/internal/handlers"
	"github.com/yaghoubi-mn/voter/internal/middleware"
)

func SetupRoutes(r *gin.Engine, authMiddleware middleware.AuthMiddleware, userHandlers handlers.UserHandler, postHandler handlers.PostHandler, commentHandler handlers.CommentHandler) {

	// middlewares
	// r.Use(middleware.Auth())

	v1 := r.Group("/api/v1")

	// users
	v1.POST("/users/login", userHandlers.Login)

	authV1 := v1.Group("/")
	authV1.Use(authMiddleware.Auth())

	// posts
	authV1.POST("/posts", postHandler.Create)
	authV1.PUT("/posts/:postId", postHandler.Update)
	authV1.DELETE("/posts/:postId", postHandler.Delete)
	v1.GET("/posts", postHandler.GetAll)
	v1.GET("/posts/:postId", postHandler.GetByID)
	authV1.POST("/posts/:postId/upvote", postHandler.UpVote)
	authV1.POST("/posts/:postId/downvote", postHandler.DownVote)
	authV1.DELETE("/posts/:postId/votes", postHandler.DeleteVote)

	// comments
	authV1.POST("posts/:postId/comments", commentHandler.Create)
	authV1.PUT("/comments/:commentId", commentHandler.Update)
	authV1.DELETE("/comments/:commentId", commentHandler.Delete)
	v1.GET("/posts/:postId/comments", commentHandler.GetAll)
	authV1.POST("/comments/:commentId/upvote", commentHandler.UpVote)
	authV1.POST("/comments/:commentId/downvote", commentHandler.DownVote)
	authV1.DELETE("/comments/:commentId/votes", commentHandler.DeleteVote)

}
