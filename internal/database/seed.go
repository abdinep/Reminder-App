package database

import (
	"log"
	"time"

	"github.com/abdinep/Reminder-App/internal/models"
)

// Seed populates the database with initial sample tasks if the table is empty.
func Seed() {
	var count int64
	DB.Model(&models.Task{}).Count(&count)
	if count > 0 {
		return
	}

	sampleTasks := []models.Task{
		{
			Title:       "Review Q3 Project Milestones",
			Description: "Prepare draft overview of Q3 team progress and deliverables",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "pending",
		},
		{
			Title:       "System Security & Dependency Audit",
			Description: "Perform security scan and check on third-party libraries",
			DueDate:     time.Now().Add(72 * time.Hour),
			Status:      "pending",
		},
		{
			Title:       "Submit Expense Reimbursements",
			Description: "Submit monthly receipts for office software licenses",
			DueDate:     time.Now().Add(-24 * time.Hour),
			Status:      "pending",
		},
		{
			Title:       "Client Sync & Demo Presentation",
			Description: "Prepare demo slides and environment for client presentation",
			DueDate:     time.Now().Add(1 * time.Hour),
			Status:      "pending",
		},
		{
			Title:       "Infrastructure Capacity Planning",
			Description: "Evaluate server resource usage and scaling for next quarter",
			DueDate:     time.Now().Add(7 * 24 * time.Hour),
			Status:      "pending",
		},
	}

	for _, task := range sampleTasks {
		if err := DB.Create(&task).Error; err != nil {
			log.Printf("[SEEDER ERROR] Failed to insert task '%s': %v", task.Title, err)
		}
	}

	log.Println("[SEEDER] Database seeded with 5 initial sample tasks.")
}
