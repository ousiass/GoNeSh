// Package config handles configuration file loading and management for GoNeSh.
package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds all GoNeSh configuration
type Config struct {
	// General settings
	Theme    string `mapstructure:"theme"`
	Language string `mapstructure:"language"`

	// AI settings
	AI AIConfig `mapstructure:"ai"`

	// SSH settings
	Connections []ConnectionConfig `mapstructure:"connections"`

	// Transfers settings
	Transfers []TransferConfig `mapstructure:"transfers"`

	// Git settings
	Git GitConfig `mapstructure:"git"`
}

// AIConfig holds AI-related configuration
type AIConfig struct {
	DefaultProvider string `mapstructure:"default_provider"`
	DefaultModel    string `mapstructure:"default_model"`
}

// ConnectionConfig holds SSH connection configuration
type ConnectionConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	User string `mapstructure:"user"`
	Key  string `mapstructure:"key"`
	Env  string `mapstructure:"env"`
}

// TransferConfig holds Quick Transfer configuration
type TransferConfig struct {
	Name       string `mapstructure:"name"`
	Connection string `mapstructure:"connection"`
	RemotePath string `mapstructure:"remote_path"`
	LocalPath  string `mapstructure:"local_path"`
}

// GitConfig holds Git Auto Commit configuration
type GitConfig struct {
	AutoCommit AutoCommitConfig `mapstructure:"auto_commit"`
}

// AutoCommitConfig holds auto commit settings
type AutoCommitConfig struct {
	Enabled      bool              `mapstructure:"enabled"`
	Model        string            `mapstructure:"model"`
	Language     string            `mapstructure:"language"`
	Emoji        bool              `mapstructure:"emoji"`
	Candidates   int               `mapstructure:"candidates"`
	MaxDiffLines int               `mapstructure:"max_diff_lines"`
	EmojiMap     map[string]string `mapstructure:"emoji_map"`
}

// configDir returns the GoNeSh configuration directory
func configDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gonesh"), nil
}

// Load reads configuration from files
func Load() (*Config, error) {
	dir, err := configDir()
	if err != nil {
		return nil, err
	}

	// 設定ディレクトリが存在しない場合は作成
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	// Viperの設定
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	// デフォルト値を設定
	setDefaults()

	// 設定ファイルを読み込む
	if err := viper.ReadInConfig(); err != nil {
		// 設定ファイルが存在しない場合はデフォルトを作成
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return createDefaultConfig(dir)
		}
		// その他のエラーの場合もデフォルト設定を返す
		return createDefaultConfig(dir)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// setDefaults sets default configuration values
func setDefaults() {
	viper.SetDefault("theme", "tokyo-night")
	viper.SetDefault("language", "ja")
	viper.SetDefault("ai.default_provider", "gemini")
	viper.SetDefault("ai.default_model", "gemini-1.5-flash")
	viper.SetDefault("git.auto_commit.enabled", true)
	viper.SetDefault("git.auto_commit.model", "gemini-1.5-flash")
	viper.SetDefault("git.auto_commit.language", "ja")
	viper.SetDefault("git.auto_commit.emoji", true)
	viper.SetDefault("git.auto_commit.candidates", 5)
	viper.SetDefault("git.auto_commit.max_diff_lines", 500)
}

// createDefaultConfig creates default configuration files
func createDefaultConfig(dir string) (*Config, error) {
	cfg := &Config{
		Theme:    "tokyo-night",
		Language: "ja",
		AI: AIConfig{
			DefaultProvider: "gemini",
			DefaultModel:    "gemini-1.5-flash",
		},
		Git: GitConfig{
			AutoCommit: AutoCommitConfig{
				Enabled:      true,
				Model:        "gemini-1.5-flash",
				Language:     "ja",
				Emoji:        true,
				Candidates:   5,
				MaxDiffLines: 500,
				EmojiMap: map[string]string{
					"feat":     "✨",
					"fix":      "🐛",
					"update":   "📝",
					"refactor": "♻️",
					"style":    "💄",
					"test":     "✅",
					"docs":     "📚",
					"chore":    "📦",
					"perf":     "⚡",
					"security": "🔒",
				},
			},
		},
	}

	// config.yaml を作成
	configPath := filepath.Join(dir, "config.yaml")

	// ファイルが存在しない場合のみ作成
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		viper.SetConfigFile(configPath)
		setDefaults()
		_ = viper.SafeWriteConfig() // エラーは無視（権限の問題など）
	}

	return cfg, nil
}

// GetConfigDir returns the configuration directory path
func GetConfigDir() (string, error) {
	return configDir()
}
