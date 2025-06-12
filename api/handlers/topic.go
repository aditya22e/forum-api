package handlers

import (
	"forum-api/models"
	"forum-api/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TopicHandler struct {
	service   *services.TopicService
	validator *validator.Validate
}

func NewTopicHandler(service *services.TopicService) *TopicHandler {
	return &TopicHandler{service: service, validator: validator.New()}
}

func (h *TopicHandler) CreateTopic(c *gin.Context) {
	var input struct {
		Title       string     `json:"title" validate:"required,min=3,max=100"`
		Content     string     `json:"content" validate:"required"`
		Tags        []string   `json:"tags" validate:"dive,required"`
		IsScheduled bool       `json:"is_scheduled"`
		ScheduledAt *time.Time `json:"scheduled_at" validate:"required_if=IsScheduled true"`
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
	topic := models.Topic{
		UserID:      userID.(string),
		Title:       input.Title,
		Content:     input.Content,
		Tags:        input.Tags,
		IsScheduled: input.IsScheduled,
		ScheduledAt: input.ScheduledAt,
	}

	if err := h.service.CreateTopic(&topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, topic)
}

func (h *TopicHandler) UpdateTopic(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var input struct {
		Title   string   `json:"title" validate:"required,min=3,max=100"`
		Content string   `json:"content" validate:"required"`
		Tags    []string `json:"tags" validate:"dive,required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	topic, err := h.service.GetTopicByID(id)
	if err != nil || topic.UserID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized or topic not found"})
		return
	}

	topic.Title = input.Title
	topic.Content = input.Content
	topic.Tags = input.Tags

	if err := h.service.UpdateTopic(topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Notify subscribers
	go h.service.NotifySubscribers(id, topic.Title, topic.Content)
	c.JSON(http.StatusOK, topic)
}

func (h *TopicHandler) DeleteTopic(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	topic, err := h.service.GetTopicByID(id)
	if err != nil || topic.UserID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized or topic not found"})
		return
	}

	if err := h.service.DeleteTopic(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Topic deleted"})
}

func (h *TopicHandler) GetTopicsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	topics, err := h.service.GetTopicsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}

func (h *TopicHandler) GetTopicsByTags(c *gin.Context) {
	tags := strings.Split(c.Query("tags"), ",")
	topics, err := h.service.GetTopicsByTags(tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}

func (h *TopicHandler) FilterTopics(c *gin.Context) {
	var input struct {
		Tags      []string  `form:"tags" validate:"dive"`
		StartDate time.Time `form:"start_date"`
		EndDate   time.Time `form:"end_date"`
	}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	topics, err := h.service.FilterTopics(input.Tags, input.StartDate, input.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}

func (h *TopicHandler) Subscribe(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	if err := h.service.Subscribe(id, userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscribed to topic"})
}

func (h *TopicHandler) Unsubscribe(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	if err := h.service.Unsubscribe(id, userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unsubscribed from topic"})
}
