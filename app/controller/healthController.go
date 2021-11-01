// ไฟล์ app/controller/healthController.go
package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /healthcheck [get]
func HealthCheckController(c *fiber.Ctx) error {
	msg := "OK"
	c.Status(http.StatusOK)
	return c.SendString(msg)
}
