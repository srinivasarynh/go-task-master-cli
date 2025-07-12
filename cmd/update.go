package cmd

import (
	"fmt"
	"strconv"
	"taskmaster/internal/domain"
	"taskmaster/internal/ui"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [task_id]",
	Short: "update a task",
	Long:  `update a task's title, description, priority, or due date.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runUpdate,
}

func init() {
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "new task title")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "new task description")
	updateCmd.Flags().StringVarP(&priority, "priority", "p", "", "new task priority (low, medium, high)")
	updateCmd.Flags().StringVar(&dueDate, "due", "", "new due date (YYYY-MM-DD)")
}

func runUpdate(cmd *cobra.Command, args []string) error {
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	var taskPriority domain.TaskPriority
	if priority != "" {
		switch priority {
		case "low":
			taskPriority = domain.TaskPriorityLow
		case "medium":
			taskPriority = domain.TaskPriorityMedium
		case "high":
			taskPriority = domain.TaskPriorityHigh
		default:
			return fmt.Errorf("invalid priority: %s (must be low, medium or high)", priority)
		}
	}

	var parsedDueDate *time.Time
	if dueDate != "" {
		parsed, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return fmt.Errorf("invalid due date format: %s (use YYYY-MM-DD)", dueDate)
		}
		parsedDueDate = &parsed
	}

	task, err := taskService.UpdateTask(uint(id), title, description, taskPriority, parsedDueDate)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	fmt.Printf("✓ Task updated successfully!\n\n")
	ui.RenderTasksDetails(task)

	return nil
}
