package tests

import (
	"testing"

	"github.com/google/uuid"

	"github.com/aditya22e/forum-api/config"
	"github.com/aditya22e/forum-api/services"

	"github.com/aditya22e/forum-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	// Set config.DB for service usage
	config.DB = db
	// Auto-migrate models
	err = db.AutoMigrate(&models.Topic{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}

func TestCreateTopic(t *testing.T) {
	db := setupTestDB(t)

	service := services.NewTopicService()

	topic := models.Topic{
		ID:      uuid.New().String(), // Manually set UUID for SQLite
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
	err = db.First(&savedTopic, "id = ?", topic.ID).Error
	if err != nil {
		t.Errorf("Failed to retrieve topic: %v", err)
	}
	if savedTopic.Title != topic.Title {
		t.Errorf("Expected title %v, got %v", topic.Title, savedTopic.Title)
	}
	if savedTopic.UserID != topic.UserID {
		t.Errorf("Expected user_id %v, got %v", topic.UserID, savedTopic.UserID)
	}
}
