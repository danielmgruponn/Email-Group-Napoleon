package napoleonhandler

import (
	"context"
	"napoleon-email/src/app/domain/application/napoleon"
	napoleoncontact "napoleon-email/src/app/http/request/napoleon_contact"
	"napoleon-email/src/config/app"
	"napoleon-email/src/pkg/logger"
	sendemail "napoleon-email/src/pkg/send_email"
	"time"

	"github.com/gofiber/fiber/v2"
)

type NapoleonHandler struct {
	napoleonApplication napoleon.ContactEmailApplicationInterface
}

func NewNapoleonHandler(napoleonApplication napoleon.ContactEmailApplicationInterface) *NapoleonHandler {
	return &NapoleonHandler{
		napoleonApplication: napoleonApplication,
	}
}

func (h *NapoleonHandler) CreateNapoEmail(c *fiber.Ctx) error {
	var request napoleoncontact.ContactNapoleonRequest
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

	data := sendemail.CreateBodyEmail("Napoleon", request.Name, request.Email, app.EmailNapoleonTo(), request.Subject, request.Message, time.Now().Year())
	htmlContent, err := sendemail.GenerateContactEmailHTML(*data)
	if err != nil {
		logger.LogError("error generating email content", err, logger.LogStruct{Action: "generate_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating email content",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	contact, err := h.napoleonApplication.CreateNapoEmail(ctx, request.Email, request.Name, app.EmailNapoleonTo(), request.Subject, request.Message, htmlContent)
	if err != nil {
		logger.LogError("error creating Napoleon email", err, logger.LogStruct{Action: "create_napoleon_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating Napoleon email",
		})
	}

	err = sendemail.SendContactEmailHTML(htmlContent, app.EmailNapoleonTo(), request.Subject)
	if err != nil {
		logger.LogError("error sending email", err, logger.LogStruct{Action: "send_email_error", Data: app.EmailNapoleonTo()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error sending email",
		})
	}

	sent := false
	contact.Sending = &sent
	_, err = h.napoleonApplication.UpdateNapoEmail(ctx, contact)
	if err != nil {
		logger.LogError("error updating Napoleon email sending status", err, logger.LogStruct{Action: "update_napoleon_email_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Napoleon email created successfully",
	})
}
