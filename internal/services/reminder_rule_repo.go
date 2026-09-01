package services

import (
	"fmt"

	"github.com/abdinep/Reminder-App/internal/models"
	"github.com/abdinep/Reminder-App/internal/repository"
)

type ReminderRuleService interface {
	CreateRule(rule *models.ReminderRule) error
	ListAllRules() ([]models.ReminderRule, error)
	GetRuleByID(id uint) (*models.ReminderRule, error)
	UpdateRule(rule *models.ReminderRule) error
	DeleteRule(id uint) error
	ToggleRule(id uint) error
}

type reminderRuleService struct {
	ruleRepo  repository.ReminderRuleRepository
	auditRepo repository.AuditLogRepository
}

func NewReminderRuleService(ruleRepo repository.ReminderRuleRepository, auditRepo repository.AuditLogRepository) ReminderRuleService {
	return &reminderRuleService{ruleRepo: ruleRepo, auditRepo: auditRepo}
}

func (svc *reminderRuleService) CreateRule(rule *models.ReminderRule) error {
	if rule.Name == "" {
		return models.ErrInvalidInput
	}
	if rule.ConditionType != "before_due" && rule.ConditionType != "overdue" && rule.ConditionType != "on_due" {
		return models.ErrInvalidInput
	}
	if rule.ConditionValue < 0 {
		return models.ErrInvalidInput
	}

	err := svc.ruleRepo.Create(rule)
	if err == nil {
		svc.auditRepo.Create(&models.AuditLog{
			EventType:   "rule_created",
			EntityType:  "reminder_rule",
			EntityID:    rule.ID,
			Description: fmt.Sprintf("Created reminder rule '%s' (Condition: %s)", rule.Name, rule.ConditionType),
		})
	}
	return err
}

func (svc *reminderRuleService) ListAllRules() ([]models.ReminderRule, error) {
	return svc.ruleRepo.GetAll()
}

func (svc *reminderRuleService) GetRuleByID(id uint) (*models.ReminderRule, error) {
	return svc.ruleRepo.GetByID(id)
}

func (svc *reminderRuleService) UpdateRule(rule *models.ReminderRule) error {
	if rule.ConditionType != "" && rule.ConditionType != "before_due" && rule.ConditionType != "overdue" && rule.ConditionType != "on_due" {
		return models.ErrInvalidInput
	}
	if rule.ConditionValue < 0 {
		return models.ErrInvalidInput
	}

	err := svc.ruleRepo.Update(rule)
	if err == nil {
		svc.auditRepo.Create(&models.AuditLog{
			EventType:   "rule_updated",
			EntityType:  "reminder_rule",
			EntityID:    rule.ID,
			Description: fmt.Sprintf("Updated properties for reminder rule ID #%d (%s)", rule.ID, rule.Name),
		})
	}
	return err
}

func (svc *reminderRuleService) DeleteRule(id uint) error {
	err := svc.ruleRepo.Delete(id)
	if err == nil {
		svc.auditRepo.Create(&models.AuditLog{
			EventType:   "rule_deleted",
			EntityType:  "reminder_rule",
			EntityID:    id,
			Description: fmt.Sprintf("Removed reminder rule ID #%d", id),
		})
	}
	return err
}

func (svc *reminderRuleService) ToggleRule(id uint) error {
	rule, err := svc.ruleRepo.GetByID(id)
	if err != nil {
		return err
	}
	rule.IsActive = !rule.IsActive
	err = svc.ruleRepo.Update(rule)
	if err == nil {
		statusStr := "deactivated"
		if rule.IsActive {
			statusStr = "activated"
		}
		svc.auditRepo.Create(&models.AuditLog{
			EventType:   "rule_status_changed",
			EntityType:  "reminder_rule",
			EntityID:    id,
			Description: fmt.Sprintf("Rule ID #%d status changed to %s", id, statusStr),
		})
	}
	return err
}
