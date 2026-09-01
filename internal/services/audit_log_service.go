package services

import (
	"github.com/abdinep/Reminder-App/internal/models"
	"github.com/abdinep/Reminder-App/internal/repository"
)

type AuditLogService interface {
	ListAllLogs() ([]models.AuditLog, error)
	GetLogByID(id uint) (*models.AuditLog, error)
}

type auditLogService struct {
	repo repository.AuditLogRepository
}

func NewAuditLogService(repo repository.AuditLogRepository) AuditLogService {
	return &auditLogService{repo: repo}
}

func (s *auditLogService) ListAllLogs() ([]models.AuditLog, error) {
	return s.repo.GetAll()
}

func (s *auditLogService) GetLogByID(id uint) (*models.AuditLog, error) {
	return s.repo.GetByID(id)
}
