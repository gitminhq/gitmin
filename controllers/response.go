package controllers

import "github.com/gofiber/fiber/v2"

// Response being sent out from the server
type Response struct {
	Status  int
	Message string
	Data    interface{}
}

// SendJSON is a wrapper over exisiting json response of fiber
func SendJSON(c *fiber.Ctx, statusCode int, message string, data interface{}) error {

	var resp interface{}
	switch val := data.(type) {

	// Some concrete type in case you want to test against
	// case something-special:

	default:
		resp = Response{
			Status:  statusCode,
			Message: message,
			Data:    val,
		}
	}

	if err := c.Status(statusCode).JSON(resp); err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	return nil
}
