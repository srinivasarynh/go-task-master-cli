package ui

import (
	"fmt"
	"strconv"
	"strings"
	"taskmaster/internal/domain"
	"time"
)

func RenderTasksTable(tasks []*domain.Task) {
	if len(tasks) == 0 {
		fmt.Println("no tasks found")
		return
	}
	fmt.Printf("%-5s %-40s %-15s %-10s %-12s %-10s\n",
		"ID", "Title", "Status", "Priority", "Due Date", "Created")
	fmt.Println(strings.Repeat("-", 95))

	for _, task := range tasks {
		fmt.Printf("%-5s %-40s %-15s %-10s %-12s %-10s\n",
			strconv.FormatUint(uint64(task.ID), 10),
			truncateString(task.Title, 40),
			colorizeStatus(string(task.Status)),
			colorizePriority(string(task.Priority)),
			formatDueDate(task.DueDate),
			task.CreatedAt.Format("2006-01-02"),
		)
	}
}

func RenderTasksDetails(task *domain.Task) {
	fmt.Printf("task details: \n")
	fmt.Printf("  ID: %d\n", task.ID)
	fmt.Printf("  Title: %s\n", task.Title)

	if task.Description != "" {
		fmt.Printf("  Description: %s\n", task.Description)
	}
	fmt.Printf("  Status: %s\n", colorizeStatus(string(task.Status)))
	fmt.Printf("  Priority: %s\n", colorizePriority(string(task.Priority)))

	if task.DueDate != nil {
		fmt.Printf("  Due Date: %s\n", formatDueDate(task.DueDate))
	}

	fmt.Printf("  Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Updated: %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
}

func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length-3] + "..."
}

func colorizeStatus(status string) string {
	switch strings.ToLower(status) {
	case "completed":
		return fmt.Sprintf("✓ %s", status)
	case "pending":
		return fmt.Sprintf("⏳ %s", status)
	case "canceled":
		return fmt.Sprintf("✗ %s", status)
	default:
		return status
	}
}

func colorizePriority(priority string) string {
	switch strings.ToLower(priority) {
	case "high":
		return fmt.Sprintf("🔴 %s", priority)
	case "medium":
		return fmt.Sprintf("🟡 %s", priority)
	case "low":
		return fmt.Sprintf("🟢 %s", priority)
	default:
		return priority
	}
}

func formatDueDate(dueDate *time.Time) string {
	if dueDate == nil {
		return "N/A"
	}

	now := time.Now()
	if dueDate.Before(now) {
		return fmt.Sprintf("⚠️ %s", dueDate.Format("2006-01-02"))
	}

	return dueDate.Format("2006-01-02")
}
