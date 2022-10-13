package controllers

import "github.com/gofiber/fiber/v2"

// GetIndexHTML returns the index.html file from pkger
func (ctrl *Controller) GetIndexHTML(c *fiber.Ctx) error {
	return c.Render("index", map[string]interface{}{
		"name": "Hello, World!",
	})
}
