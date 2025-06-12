package services

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type SchedulerService struct {
	topicService *TopicService
}

func NewSchedulerService(topicService *TopicService) *SchedulerService {
	return &SchedulerService{topicService: topicService}
}

func (s *SchedulerService) StartScheduler() {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		topics, err := s.topicService.GetScheduledTopics(time.Now())
		if err != nil {
			logrus.Error("Failed to fetch scheduled topics: ", err)
			return
		}
		for _, topic := range topics {
			topic.IsScheduled = false
			if err := s.topicService.UpdateTopic(&topic); err != nil {
				logrus.Error("Failed to publish topic: ", err)
			}
		}
	})
	c.Start()
}
