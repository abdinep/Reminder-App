package repository

import (
	"github.com/abdinep/Reminder-App/internal/models"
	"gorm.io/gorm"
)

type ReminderRuleRepository interface {
	Create(rule *models.ReminderRule) error
	GetAll() ([]models.ReminderRule, error)
	GetByID(id uint) (*models.ReminderRule, error)
	Update(rule *models.ReminderRule) error
	Delete(id uint) error
	GetActive() ([]models.ReminderRule, error)
}

type reminderRuleRepository struct {
	db *gorm.DB
}

func NewReminderRuleRepository(db *gorm.DB) ReminderRuleRepository {
	return &reminderRuleRepository{db: db}
}

func (r *reminderRuleRepository) Create(rule *models.ReminderRule) error {
	return r.db.Create(rule).Error
}

func (r *reminderRuleRepository) GetAll() ([]models.ReminderRule, error) {
	var rules []models.ReminderRule
	err := r.db.Find(&rules).Error
	return rules, err
}

func (r *reminderRuleRepository) GetByID(id uint) (*models.ReminderRule, error) {
	var rule models.ReminderRule
	err := r.db.First(&rule, id).Error
	return &rule, err
}

func (r *reminderRuleRepository) Update(rule *models.ReminderRule) error {
	return r.db.Save(rule).Error
}

func (r *reminderRuleRepository) Delete(id uint) error {
	return r.db.Delete(&models.ReminderRule{}, id).Error
}

func (r *reminderRuleRepository) GetActive() ([]models.ReminderRule, error) {
	var rules []models.ReminderRule
	err := r.db.Where("is_active = ?", true).Find(&rules).Error
	return rules, err
}
