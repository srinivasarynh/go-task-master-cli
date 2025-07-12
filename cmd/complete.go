package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task_id]",
	Short: "mark a task as complete",
	Long:  "mark a task as completed by providing its ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runComplete,
}

func runComplete(cmd *cobra.Command, args []string) error {
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}
	if err := taskService.CompleteTask(uint(id)); err != nil {
		return fmt.Errorf("failed to complete task: %w", err)
	}

	fmt.Printf("✓ Task %d marked as completed!\n", id)
	return nil
}
