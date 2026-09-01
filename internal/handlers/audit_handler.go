package handlers

import (
	"net/http"
	"strconv"

	 "github.com/abdinep/Reminder-App/internal/services"

	"github.com/gin-gonic/gin"
)

type AuditLogHandler struct {
	service services.AuditLogService
}

func NewAuditLogHandler(service services.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{service: service}
}

func (h *AuditLogHandler) List(c *gin.Context) {
	logs, err := h.service.ListAllLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (h *AuditLogHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log, err := h.service.GetLogByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Audit log not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}
