package config

import "html/template"

// Appconfig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
