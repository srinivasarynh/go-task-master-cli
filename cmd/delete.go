package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task_id]",
	Short: "delete a task",
	Long:  "delete a task by providing its ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runDelete,
}

func runDelete(cmd *cobra.Command, args []string) error {
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	if err := taskService.DeleteTask(uint(id)); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	fmt.Printf("✓ Task %d deleted successfully!\n", id)
	return nil
}
