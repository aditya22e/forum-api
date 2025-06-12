package tests

import (
	"forum-api/models"
	"forum-api/services"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	db.AutoMigrate(&models.Topic{})
	return db
}

func TestCreateTopic(t *testing.T) {
	db := setupTestDB(t)
	service := &services.TopicService{db: db}

	topic := models.Topic{
		UserID:  "user1",
		Title:   "Test Topic",
		Content: "Test Content",
		Tags:    []string{"test", "example"},
	}

	err := service.CreateTopic(&topic)
	if err != nil {
		t.Errorf("Failed to create topic: %v", err)
	}

	var savedTopic models.Topic
	db.First(&savedTopic, "id = ?", topic.ID)
	if savedTopic.Title != topic.Title {
		t.Errorf("Expected title %v, got %v", topic.Title, savedTopic.Title)
	}
}
