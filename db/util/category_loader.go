package util

import (
	"encoding/json"
	"fmt"
	"os"
)

// I18nEntry represents a single localized name for a category
type I18nEntry struct {
	Lang string `json:"lang"`
	Name string `json:"name"`
}

// Category represents a single activity category with its aliases and translations
type Category struct {
	Name    string      `json:"name"`
	Aliases []string    `json:"aliases"`
	I18n    []I18nEntry `json:"i18n"`
}

// CategoryConfig represents the entire category configuration structure
type CategoryConfig struct {
	Categories []Category `json:"categories"`
}

// LoadCategoryConfig reads and parses the category configuration JSON file
func LoadCategoryConfig(filePath string) (*CategoryConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read category config file: %w", err)
	}

	var config CategoryConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse category config JSON: %w", err)
	}

	return &config, nil
}
