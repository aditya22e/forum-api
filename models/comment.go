package models

import "time"

type Comment struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TopicID   string    `gorm:"type:uuid;not null"`
	UserID    string    `gorm:"type:uuid;not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
