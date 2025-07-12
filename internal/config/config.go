package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

type LoggingConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

var AppConfig *Config

func Load() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(home, ".taskmaster")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.AddConfigPath("./configs")

	viper.SetDefault("database.path", filepath.Join(configDir, "tasks.db"))
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.file", filepath.Join(configDir, "taskmaster.log"))

	viper.SetEnvPrefix("TASKMASTER")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := createDefaultConfig(configDir); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	AppConfig = &Config{}
	return viper.Unmarshal(AppConfig)
}

func createDefaultConfig(configDir string) error {
	configPath := filepath.Join(configDir, "config.yaml")

	defaultConfig := `database:
	path: ~/.taskmaster/task.db 
	logging:
	level:info
	file: ~/.tadkmater/taskmaster.log
	`
	return os.WriteFile(configPath, []byte(defaultConfig), 0644)
}
