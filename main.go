package main

import (
	"fmt"
	"os"
	"taskmaster/cmd"
	"taskmaster/internal/config"
	"taskmaster/internal/database"
	"taskmaster/pkg/logger"
)

func main() {
	logger.Init()

	if err := config.Load(); err != nil {
		fmt.Printf("failed to load confituration: %v\n", err)
		os.Exit(1)
	}

	if err := database.Initialize(); err != nil {
		fmt.Printf("failed to initialize database: %v\n", err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
