package cmd

import (
	"fmt"
	"taskmaster/internal/domain"
	"taskmaster/internal/ui"

	"github.com/spf13/cobra"
)

var (
	status   string
	overdue  bool
	detailed bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  "list all tasks or filter by status",
	RunE:  runList,
}

func init() {
	listCmd.Flags().StringVarP(&status, "status", "s", "", "filter by status (pending, completed, canceled)")
	listCmd.Flags().BoolVar(&overdue, "overdue", false, "show only overdue tasks")
	listCmd.Flags().BoolVar(&detailed, "detailed", false, "show detailed task information")
}

func runList(cmd *cobra.Command, args []string) error {
	var tasks []*domain.Task
	var err error

	if overdue {
		tasks, err = taskService.GetOverdueTasks()
		if err != nil {
			return fmt.Errorf("failed to get overdue tasks: %w", err)
		}
		fmt.Println("overdue tasks: ")
	} else if status != "" {
		var taskStatus domain.TaskStatus
		switch status {
		case "pending":
			taskStatus = domain.TaskStatusPending

		case "completed":
			taskStatus = domain.TaskStatusCompleted

		case "canceled":
			taskStatus = domain.TaskStatusCanceled

		default:
			return fmt.Errorf("invalid status: %s (must be pending, complete, canceled)", status)
		}

		tasks, err = taskService.GetTaskByStatus(taskStatus)
		if err != nil {
			return fmt.Errorf("failed to get tasks by status: %w", err)
		}
		fmt.Printf("tasks with status '%s':\n", status)
	} else {
		tasks, err = taskService.GetAllTasks()
		if err != nil {
			return fmt.Errorf("failed to get task: %w", err)
		}
		fmt.Println("all tasks: ")
	}

	if detailed {
		for i, task := range tasks {
			if i > 0 {
				fmt.Println()
			}
			ui.RenderTasksDetails(task)
		}
	} else {
		ui.RenderTasksTable(tasks)
	}
	return nil
}
