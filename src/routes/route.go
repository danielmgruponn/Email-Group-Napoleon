package routes

import (
	"napoleon-email/src/app/http/handler"
	"napoleon-email/src/app/infrastructure"
	"napoleon-email/src/routes/api"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	c := infrastructure.GetKernel()
	app.Get("/", handler.HealthCheck)
	api.RouterApi(app, c)
}
