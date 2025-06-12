package models

import "time"

type Subscription struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TopicID   string    `gorm:"type:uuid;not null"`
	UserID    string    `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
