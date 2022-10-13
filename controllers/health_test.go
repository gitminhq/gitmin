package controllers

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetHandlerHealthz(t *testing.T) {
	assert := assert.New(t)
	api := fiber.New()
	routers := New()

	api.Get("/healthz", routers.GetHandlerHealthz)

	req := httptest.NewRequest("GET", "/healthz", nil)
	resp, _ := api.Test(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail("Failed to read the response from server")
	}
	assert.Contains(strings.ToLower(string(body)), "healthy")
}
