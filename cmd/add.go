package cmd

import (
	"fmt"
	"taskmaster/internal/domain"
	"taskmaster/internal/ui"
	"time"

	"github.com/spf13/cobra"
)

var (
	title       string
	description string
	priority    string
	dueDate     string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Long:  `add a new task with optional description, priority, and due date.`,
	RunE:  runAdd,
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "task title (required)")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "task description")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "medium", "task priority (low, medium, high)")
	addCmd.Flags().StringVar(&dueDate, "due", "", "due date (YYY-MM-DD format)")

	addCmd.MarkFlagRequired("title")
}

func runAdd(cmd *cobra.Command, args []string) error {
	var taskPriority domain.TaskPriority
	switch priority {
	case "low":
		taskPriority = domain.TaskPriorityLow
	case "medium":
		taskPriority = domain.TaskPriorityMedium
	case "high":
		taskPriority = domain.TaskPriorityHigh
	default:
		return fmt.Errorf("invalid priority: %s (must be low, medium, or high)", priority)
	}

	var parsedDueDate *time.Time
	if dueDate != "" {
		parsed, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return fmt.Errorf("invalid due date format: %s (use YYYY-MM-DD)", dueDate)
		}
		parsedDueDate = &parsed
	}

	task, err := taskService.CreateTask(title, description, taskPriority, parsedDueDate)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	fmt.Printf("✓ Task created successfully!\n\n")
	ui.RenderTasksDetails(task)

	return nil
}
