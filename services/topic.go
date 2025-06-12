package services

import (
	"time"

	"github.com/aditya22e/forum-api/config"
	"github.com/aditya22e/forum-api/models"

	"gorm.io/gorm"
)

type TopicService struct {
	db           *gorm.DB
	emailService *EmailService
}

func NewTopicService() *TopicService {
	return &TopicService{db: config.DB, emailService: NewEmailService()}
}

func (s *TopicService) CreateTopic(topic *models.Topic) error {
	return s.db.Create(topic).Error
}

func (s *TopicService) UpdateTopic(topic *models.Topic) error {
	return s.db.Save(topic).Error
}

func (s *TopicService) DeleteTopic(id string) error {
	return s.db.Delete(&models.Topic{}, "id = ?", id).Error
}

func (s *TopicService) GetTopicByID(id string) (*models.Topic, error) {
	var topic models.Topic
	err := s.db.Where("id = ?", id).First(&topic).Error
	return &topic, err
}

func (s *TopicService) GetTopicsByUser(userID string) ([]models.Topic, error) {
	var topics []models.Topic
	err := s.db.Where("user_id = ?", userID).Find(&topics).Error
	return topics, err
}

func (s *TopicService) GetTopicsByTags(tags []string) ([]models.Topic, error) {
	var topics []models.Topic
	err := s.db.Where("tags @> ?", tags).Find(&topics).Error
	return topics, err
}

func (s *TopicService) FilterTopics(tags []string, startDate, endDate time.Time) ([]models.Topic, error) {
	var topics []models.Topic
	query := s.db.Model(&models.Topic{})
	if len(tags) > 0 {
		query = query.Where("tags @> ?", tags)
	}
	if !startDate.IsZero() {
		query = query.Where("created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("created_at <= ?", endDate)
	}
	err := query.Find(&topics).Error
	return topics, err
}

func (s *TopicService) CreateComment(comment *models.Comment) error {
	return s.db.Create(comment).Error
}

func (s *TopicService) Subscribe(topicID, userID string) error {
	subscription := models.Subscription{TopicID: topicID, UserID: userID}
	return s.db.Create(&subscription).Error
}

func (s *TopicService) Unsubscribe(topicID, userID string) error {
	return s.db.Delete(&models.Subscription{}, "topic_id = ? AND user_id = ?", topicID, userID).Error
}

func (s *TopicService) NotifySubscribers(topicID, title, content string) {
	var subscriptions []models.Subscription
	s.db.Where("topic_id = ?", topicID).Find(&subscriptions)

	var userIDs []string
	for _, sub := range subscriptions {
		userIDs = append(userIDs, sub.UserID)
	}

	var users []models.User
	s.db.Where("id IN ?", userIDs).Find(&users)

	for _, user := range users {
		s.emailService.SendEmail(user.Email, "Topic Updated: "+title, "The topic has been updated:\n\n"+content)
	}
}

func (s *TopicService) GetScheduledTopics(now time.Time) ([]models.Topic, error) {
	var topics []models.Topic
	err := s.db.Where("is_scheduled = ? AND scheduled_at <= ?", true, now).Find(&topics).Error
	return topics, err
}
