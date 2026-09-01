package models

import "time"

type AuditLog struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EventType   string    `json:"event_type"`  // rule_created, rule_updated, rule_deleted, rule_status_changed, reminder_triggered
	EntityType  string    `json:"entity_type"` // reminder_rule, task
	EntityID    uint      `json:"entity_id"`
	RuleID      uint      `json:"rule_id,omitempty"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
