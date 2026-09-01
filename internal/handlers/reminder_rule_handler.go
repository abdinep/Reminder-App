package handlers

import (
	"net/http"
	"strconv"

	"github.com/abdinep/Reminder-App/internal/models"
	"github.com/abdinep/Reminder-App/internal/services"
	"github.com/gin-gonic/gin"
)

type ReminderRuleHandler struct {
	service services.ReminderRuleService
}

func NewReminderRuleHandler(service services.ReminderRuleService) *ReminderRuleHandler {
	return &ReminderRuleHandler{service: service}
}

func (h *ReminderRuleHandler) Create(c *gin.Context) {
	var rule models.ReminderRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateRule(&rule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rule)
}

func (h *ReminderRuleHandler) List(c *gin.Context) {
	rules, err := h.service.ListAllRules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rules)
}

func (h *ReminderRuleHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rule, err := h.service.GetRuleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rule not found"})
		return
	}

	c.JSON(http.StatusOK, rule)
}

func (h *ReminderRuleHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var rule models.ReminderRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rule.ID = uint(id)
	if err := h.service.UpdateRule(&rule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rule)
}

func (h *ReminderRuleHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteRule(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rule deleted"})
}

func (h *ReminderRuleHandler) Toggle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.ToggleRule(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rule status toggled"})
}
