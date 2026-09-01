package models

import (
	"time"

	"gorm.io/gorm"
)

type ReminderRule struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `json:"reminder_name"`
	Description    string         `json:"reminder_description"`
	ConditionType  string         `json:"condition_type"` // any of {before_due, overdue, on_due}
	ConditionValue int            `json:"condition_value"`
	IsActive       bool           `json:"is_active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
