package handlers

import (
	"forum-api/models"
	"forum-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentHandler struct {
	service   *services.TopicService
	validator *validator.Validate
}

func NewCommentHandler(service *services.TopicService) *CommentHandler {
	return &CommentHandler{service: service, validator: validator.New()}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var input struct {
		Content string `json:"content" validate:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	comment := models.Comment{
		TopicID: c.Param("id"),
		UserID:  userID.(string),
		Content: input.Content,
	}

	if err := h.service.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
