package repository

import (
	"github.com/abdinep/Reminder-App/internal/models"
	"gorm.io/gorm"
)

type AuditLogRepository interface {
	Create(log *models.AuditLog) error
	GetAll() ([]models.AuditLog, error)
	GetByID(id uint) (*models.AuditLog, error)
	HasRecentTrigger(taskID, ruleID uint, duration string) (bool, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) Create(log *models.AuditLog) error {
	return r.db.Create(log).Error
}

func (r *auditLogRepository) GetAll() ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Order("created_at desc").Find(&logs).Error
	return logs, err
}

func (r *auditLogRepository) GetByID(id uint) (*models.AuditLog, error) {
	var log models.AuditLog
	err := r.db.First(&log, id).Error
	return &log, err
}

func (r *auditLogRepository) HasRecentTrigger(taskID, ruleID uint, duration string) (bool, error) {
	var count int64
	err := r.db.Model(&models.AuditLog{}).
		Where("event_type = ? AND entity_id = ? AND rule_id = ? AND created_at > NOW() - ?::interval",
			"reminder_triggered", taskID, ruleID, duration).
		Count(&count).Error
	return count > 0, err
}
