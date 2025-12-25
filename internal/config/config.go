package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

const (
	// ConfigFileName is the name of the configuration file
	ConfigFileName = ".linear.toml"
)

// Config represents the CLI configuration
type Config struct {
	APIKey  string `toml:"api_key"`
	TeamID  string `toml:"team_id"`
	TeamKey string `toml:"team_key"`
}

// Manager handles configuration loading and saving
type Manager struct {
	config     *Config
	configPath string
}

// NewManager creates a new configuration manager
func NewManager() (*Manager, error) {
	// Look for config in current directory, then home directory
	paths := []string{
		filepath.Join(".", ConfigFileName),
	}

	home, err := os.UserHomeDir()
	if err == nil {
		paths = append(paths, filepath.Join(home, ConfigFileName))
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return &Manager{configPath: path}, nil
		}
	}

	// Default to current directory if no config found
	return &Manager{configPath: paths[0]}, nil
}

// Load loads the configuration from disk
func (m *Manager) Load() (*Config, error) {
	if m.config != nil {
		return m.config, nil
	}

	data, err := os.ReadFile(m.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			m.config = &Config{}
			return m.config, nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Also check environment variables
	if apiKey := os.Getenv("LINEAR_API_KEY"); apiKey != "" {
		cfg.APIKey = apiKey
	}

	m.config = &cfg
	return m.config, nil
}

// Save saves the configuration to disk
func (m *Manager) Save(cfg *Config) error {
	data, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(m.configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	m.config = cfg
	return nil
}

// Get returns a configuration value
func (m *Manager) Get(key string) (string, error) {
	cfg, err := m.Load()
	if err != nil {
		return "", err
	}

	switch key {
	case "api_key":
		return cfg.APIKey, nil
	case "team_id":
		return cfg.TeamID, nil
	case "team_key":
		return cfg.TeamKey, nil
	default:
		return "", fmt.Errorf("unknown config key: %s", key)
	}
}

// Set sets a configuration value
func (m *Manager) Set(key, value string) error {
	cfg, err := m.Load()
	if err != nil {
		return err
	}

	switch key {
	case "api_key":
		cfg.APIKey = value
	case "team_id":
		cfg.TeamID = value
	case "team_key":
		cfg.TeamKey = value
	default:
		return fmt.Errorf("unknown config key: %s", key)
	}

	return m.Save(cfg)
}

// Path returns the configuration file path
func (m *Manager) Path() string {
	return m.configPath
}

// IsConfigured returns whether the CLI is properly configured
func (m *Manager) IsConfigured() bool {
	cfg, err := m.Load()
	if err != nil {
		return false
	}
	return cfg.APIKey != "" && cfg.TeamKey != ""
}
