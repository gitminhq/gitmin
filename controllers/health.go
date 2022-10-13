package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// GetHandlerHealthz godoc
// @Summary Get the health status of application
// @Description Get the health status of application
// @Tags Health
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /healthz [get]
// GetHandlerHealthz returns the health
func (ctrl *Controller) GetHandlerHealthz(c *fiber.Ctx) error {
	empty := map[string]interface{}{}
	return SendJSON(c, fiber.StatusOK, "Healthy", empty)
}
