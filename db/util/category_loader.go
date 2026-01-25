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

// GetCategoryByName finds a category by its primary name
func (cc *CategoryConfig) GetCategoryByName(name string) *Category {
	for i := range cc.Categories {
		if cc.Categories[i].Name == name {
			return &cc.Categories[i]
		}
	}
	return nil
}

// GetCategoryByAlias finds a category by one of its aliases (case-insensitive)
func (cc *CategoryConfig) GetCategoryByAlias(alias string) *Category {
	for i := range cc.Categories {
		for _, a := range cc.Categories[i].Aliases {
			if a == alias {
				return &cc.Categories[i]
			}
		}
	}
	return nil
}

// GetLocalizedName returns the localized name for a category in the specified language
func (c *Category) GetLocalizedName(lang string) string {
	for _, entry := range c.I18n {
		if entry.Lang == lang {
			return entry.Name
		}
	}
	// Return the primary name if language not found
	return c.Name
}

// GetAllLanguages returns all available language codes from the config
func (cc *CategoryConfig) GetAllLanguages() map[string]bool {
	languages := make(map[string]bool)
	for _, category := range cc.Categories {
		for _, i18n := range category.I18n {
			languages[i18n.Lang] = true
		}
	}
	return languages
}

// PrintCategoryInfo prints human-readable information about all categories
func (cc *CategoryConfig) PrintCategoryInfo() {
	fmt.Printf("Total categories: %d\n\n", len(cc.Categories))
	for _, cat := range cc.Categories {
		fmt.Printf("Category: %s\n", cat.Name)
		fmt.Printf("  Aliases: %v\n", cat.Aliases)
		fmt.Printf("  Available languages: %d\n", len(cat.I18n))
		for _, i18n := range cat.I18n {
			fmt.Printf("    - %s: %s\n", i18n.Lang, i18n.Name)
		}
		fmt.Println()
	}
}
