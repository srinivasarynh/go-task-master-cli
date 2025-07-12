package cmd

import (
	"taskmaster/internal/database"
	"taskmaster/internal/repository"
	"taskmaster/internal/service"
	"taskmaster/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	taskService service.TaskService
	rootCmd     = &cobra.Command{
		Use:   "taskmaster",
		Short: "a powerful CLI task management tool",
		Long: `TaskMaster is a production-ready CLI application for managing tasks.
		It provides a clean interface for creating, updating, and tracking tasks with 
		priorities, due dates, and completion status`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initServices)

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
}

func initServices() {
	taskRepo := repository.NewTaskRepository(database.GetDB())
	taskService = service.NewTaskService(taskRepo, logger.GetLogger())
}
