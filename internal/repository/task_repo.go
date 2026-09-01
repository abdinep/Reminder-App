package repository

import (
	"github.com/abdinep/Reminder-App/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetPending() ([]models.Task, error)
	GetByID(id uint) (*models.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetPending() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("status = ?", "pending").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}
