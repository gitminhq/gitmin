package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django"

	"gitminhq/gitmin/config"
)

// App struct consists of our application structure
type App struct {
	Server   *fiber.App
	Config   *config.Config
}

type appConfig struct {
	// Put all your html/jinja templates here, defaults to /templates in pkger
	htmlTemplateDir string

	// Put all your assets like css, js etc.
	staticAssetDir string

	// configuration
	config *config.Config

	// DB instance
}

// ConfigOption will be a set of configuration options for APP
type ConfigOption struct {
	setup func(ro *appConfig)
}

// WithHTMLTemplateDir will provide the base dir to load templates from pkger
func WithHTMLTemplateDir(htmlTemplateDir string) ConfigOption {
	return ConfigOption{func(ac *appConfig) {
		ac.htmlTemplateDir = htmlTemplateDir
	}}
}

// WithStaticAssetDir will provide the base dir to load static assets from pkger
func WithStaticAssetDir(staticAssetDir string) ConfigOption {
	return ConfigOption{func(ac *appConfig) {
		ac.staticAssetDir = staticAssetDir
	}}
}

// WithConfiguration will register a configuration of type *config.Config
func WithConfiguration(config *config.Config) ConfigOption {
	return ConfigOption{func(ac *appConfig) {
		ac.config = config
	}}
}

// New creates a new instance of App
func New(configOptions ...ConfigOption) *App {

	ac := &appConfig{}
	for _, option := range configOptions {
		option.setup(ac)
	}

	var engine *django.Engine
	// Register the templates directory which is packaged using pkger
	if ac.htmlTemplateDir != "" {
		engine = django.New(ac.htmlTemplateDir, ".html")

	}
	engine.Reload(true)

	// Instantiate a fiber application
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Register the static assets like css, js etc
	if ac.staticAssetDir != "" {
		app.Static("", ac.staticAssetDir)
	}

	// Use default recoverer middleware/recover
	app.Use(recover.New())

	return &App{
		Server:   app,
		Config:   ac.config,
	}
}
