package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/abdinep/Reminder-App/internal/models"
	"github.com/abdinep/Reminder-App/internal/repository"
	"github.com/robfig/cron"
)

type Scheduler struct {
	ruleRepo  repository.ReminderRuleRepository
	taskRepo  repository.TaskRepository
	auditRepo repository.AuditLogRepository
}

func NewScheduler(ruleRepo repository.ReminderRuleRepository, taskRepo repository.TaskRepository, auditRepo repository.AuditLogRepository) *Scheduler {
	return &Scheduler{
		ruleRepo:  ruleRepo,
		taskRepo:  taskRepo,
		auditRepo: auditRepo,
	}
}

// Start launches the background ticker to run reminder evaluations every minute.
func (sch *Scheduler) Start() {
	c := cron.New()
	err := c.AddFunc("* * * * *", func() {
		sch.RunReminders()
	})
	if err != nil {
		log.Fatalf("[CRON ERROR] Failed to schedule reminder job: %v", err)
	}
	c.Start()
	log.Println("[SCHEDULER ENGINE] Reminder cron worker started (Interval: 1m).")
}

// RunReminders fetches active rules and pending tasks to evaluate reminder conditions.
func (sch *Scheduler) RunReminders() {
	rules, err := sch.ruleRepo.GetActive()
	if err != nil {
		log.Printf("[SCHEDULER ERROR] Failed fetching active rules: %v", err)
		return
	}

	tasks, err := sch.taskRepo.GetPending()
	if err != nil {
		log.Printf("[SCHEDULER ERROR] Failed fetching pending tasks: %v", err)
		return
	}

	for _, rule := range rules {
		for _, task := range tasks {
			if sch.evaluateRule(rule, task) {
				// Deduplication check: prevent triggering same rule for task within 1 hour
				alreadySent, err := sch.auditRepo.HasRecentTrigger(task.ID, rule.ID, "1 hour")
				if err != nil {
					log.Printf("[SCHEDULER WARNING] Deduplication check failed for task %d: %v", task.ID, err)
					continue
				}
				if alreadySent {
					continue
				}

				log.Printf("[REMINDER NOTIFICATION] Rule '%s' matched Task '%s'", rule.Name, task.Title)

				sch.auditRepo.Create(&models.AuditLog{
					EventType:   "reminder_triggered",
					EntityType:  "task",
					EntityID:    task.ID,
					RuleID:      rule.ID,
					Description: fmt.Sprintf("Triggered reminder rule '%s' for task: %s", rule.Name, task.Title),
				})
			}
		}
	}
}

// evaluateRule checks whether a reminder condition matches the given task's due date.
func (sch *Scheduler) evaluateRule(rule models.ReminderRule, task models.Task) bool {
	now := time.Now()
	timeRemaining := task.DueDate.Sub(now)

	switch rule.ConditionType {
	case "before_due":
		// Triggers if remaining time is positive and within the specified hour threshold
		thresholdHours := time.Duration(rule.ConditionValue) * time.Hour
		return timeRemaining > 0 && timeRemaining <= thresholdHours

	case "overdue":
		// Triggers if task deadline has passed and task is still pending
		return now.After(task.DueDate) && task.Status == "pending"

	case "on_due":
		// Triggers if current time is within a 30-minute window of the due date
		toleranceWindow := 30 * time.Minute
		diff := timeRemaining
		if diff < 0 {
			diff = -diff
		}
		return diff <= toleranceWindow

	default:
		return false
	}
}
