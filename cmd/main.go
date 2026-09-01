package main

import (
	"log"
	"os"

	"github.com/abdinep/Reminder-App/internal/database"
	"github.com/abdinep/Reminder-App/internal/handlers"
	"github.com/abdinep/Reminder-App/internal/repository"
	"github.com/abdinep/Reminder-App/internal/routes"
	"github.com/abdinep/Reminder-App/internal/scheduler"
	"github.com/abdinep/Reminder-App/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("[CONFIG] No .env file loaded; defaulting to system environment variables")
	}

	// 1. Establish database connection & seed initial data
	database.InitDB()
	database.Seed()

	// 2. Instantiate repository layer
	ruleRepo := repository.NewReminderRuleRepository(database.DB)
	auditRepo := repository.NewAuditLogRepository(database.DB)
	taskRepo := repository.NewTaskRepository(database.DB)

	// 3. Instantiate business service layer
	ruleService := services.NewReminderRuleService(ruleRepo, auditRepo)
	auditService := services.NewAuditLogService(auditRepo)

	// 4. Instantiate API handler layer
	ruleHandler := handlers.NewReminderRuleHandler(ruleService)
	auditHandler := handlers.NewAuditLogHandler(auditService)

	// 5. Launch background scheduler for reminder processing
	reminderScheduler := scheduler.NewScheduler(ruleRepo, taskRepo, auditRepo)
	reminderScheduler.Start()

	// 6. Configure Gin HTTP router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Register API endpoints
	routes.SetupRoutes(router, ruleHandler, auditHandler)

	// 7. Start HTTP server listener
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	log.Printf("==================================================")
	log.Printf("Server starting on port :%s", serverPort)
	log.Printf("==================================================")

	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf("[FATAL] Server launch failed: %v", err)
	}
}
