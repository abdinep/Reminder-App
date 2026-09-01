package routes

import (
	"github.com/abdinep/Reminder-App/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ruleHandler *handlers.ReminderRuleHandler, auditHandler *handlers.AuditLogHandler) {
	api := r.Group("/api")
	{
		// Reminder Rules
		api.POST("/reminder-rules", ruleHandler.Create)
		api.GET("/reminder-rules", ruleHandler.List)
		api.GET("/reminder-rules/:id", ruleHandler.Get)
		api.PUT("/reminder-rules/:id", ruleHandler.Update)
		api.DELETE("/reminder-rules/:id", ruleHandler.Delete)
		api.PATCH("/reminder-rules/:id/toggle", ruleHandler.Toggle)

		// Activity Logs
		api.GET("/activity-logs", auditHandler.List)
		api.GET("/activity-logs/:id", auditHandler.Get)
	}
}
