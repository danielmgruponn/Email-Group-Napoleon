package napoleonminehandler

import (
	"context"
	"napoleon-email/src/app/domain/application/mine"
	napoleonminecontact "napoleon-email/src/app/http/request/napoleon_mine_contact"
	"napoleon-email/src/config/app"
	"napoleon-email/src/pkg/logger"
	sendemail "napoleon-email/src/pkg/send_email"
	"time"

	"github.com/gofiber/fiber/v2"
)

type NapoleonMineHandler struct {
	mineApplication mine.ContactMineEmailApplicationInterface
}

func NewNapoleonMineHandler(mineApplication mine.ContactMineEmailApplicationInterface) *NapoleonMineHandler {
	return &NapoleonMineHandler{
		mineApplication: mineApplication,
	}
}

func (h *NapoleonMineHandler) CreateNapoMineEmail(c *fiber.Ctx) error {
	var request napoleonminecontact.ContactNapoleonMineRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := request.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data := sendemail.CreateBodyEmail("Napoleon Gold Mine", request.Name, request.Email, app.EmailNapoleonMineTo(), request.Subject, request.Message, time.Now().Year())
	htmlContent, err := sendemail.GenerateContactEmailHTML(*data)
	if err != nil {
		logger.LogError("error generating email content", err, logger.LogStruct{Action: "generate_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating email content",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	contact, err := h.mineApplication.CreateNapoMineEmail(ctx, request.Email, request.Name, app.EmailNapoleonMineTo(), request.Subject, request.Message, htmlContent)
	if err != nil {
		logger.LogError("error creating Napoleon Mine email", err, logger.LogStruct{Action: "create_napoleon_mine_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating Napoleon Mine email",
		})
	}

	err = sendemail.SendContactEmailHTML(htmlContent, app.EmailNapoleonMineTo(), request.Subject)
	if err != nil {
		logger.LogError("error sending email", err, logger.LogStruct{Action: "send_email_error", Data: app.EmailNapoleonMineTo()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error sending email",
		})
	}

	sent := false
	contact.Sending = &sent
	_, err = h.mineApplication.UpdateNapoMineEmail(ctx, contact)
	if err != nil {
		logger.LogError("error updating Napoleon Mine email sending status", err, logger.LogStruct{Action: "update_napoleon_mine_email_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Napoleon Mine email sent successfully",
	})
}
