package models

import "time"

type Topic struct {
	ID          string   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID      string   `gorm:"type:uuid;not null"`
	Title       string   `gorm:"not null"`
	Content     string   `gorm:"not null"`
	Tags        []string `gorm:"type:text[]"`
	IsScheduled bool     `gorm:"default:false"`
	ScheduledAt *time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
