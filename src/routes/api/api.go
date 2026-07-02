package api

import (
	"napoleon-email/src/app/infrastructure"
	v1 "napoleon-email/src/routes/api/v1"

	"github.com/gofiber/fiber/v2"
)

func RouterApi(app *fiber.App, c *infrastructure.Kernel) {
	v1.RouterApiV1(app, c)
}
