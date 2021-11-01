package routes

import (
	"learnFiberSwagger/app/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {
	app.Get("/healthcheck", controller.HealthCheckController)
}
