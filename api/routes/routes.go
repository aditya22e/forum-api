package routes

import (
	"forum-api/api/handlers"
	"forum-api/api/middleware"
	"forum-api/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userService := services.NewUserService()
	topicService := services.NewTopicService()

	userHandler := handlers.NewUserHandler(userService)
	topicHandler := handlers.NewTopicHandler(topicService)
	commentHandler := handlers.NewCommentHandler(topicService)

	// Public routes
	r.POST("/api/users/register", userHandler.Register)
	r.POST("/api/users/login", userHandler.Login)

	// Protected routes
	protected := r.Group("/api", middleware.AuthMiddleware())
	{
		// Topic routes
		protected.POST("/topics", topicHandler.CreateTopic)
		protected.PUT("/topics/:id", topicHandler.UpdateTopic)
		protected.DELETE("/topics/:id", topicHandler.DeleteTopic)
		protected.POST("/topics/:id/subscribe", topicHandler.Subscribe)
		protected.DELETE("/topics/:id/unsubscribe", topicHandler.Unsubscribe)
		protected.POST("/topics/:id/comments", commentHandler.CreateComment)

		// Public topic routes
		r.GET("/topics/user/:user_id", topicHandler.GetTopicsByUser)
		r.GET("/topics/tags", topicHandler.GetTopicsByTags)
		r.GET("/topics/filter", topicHandler.FilterTopics)
	}
}
