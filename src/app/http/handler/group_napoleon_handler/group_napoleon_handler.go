package groupnapoleonhandler

import (
	"context"
	groupnapoleon "napoleon-email/src/app/domain/application/group_napoleon"
	groupnapoleoncontact "napoleon-email/src/app/http/request/group_napoleon_contact"
	"napoleon-email/src/config/app"
	"napoleon-email/src/pkg/logger"
	sendemail "napoleon-email/src/pkg/send_email"
	"time"

	"github.com/gofiber/fiber/v2"
)

type GroupNapoleonHandler struct {
	groupNapoleonApplication groupnapoleon.ContactGroupNapoleonEmailApplicationInterface
}

func NewGroupNapoleonHandler(groupNapoleonApplication groupnapoleon.ContactGroupNapoleonEmailApplicationInterface) *GroupNapoleonHandler {
	return &GroupNapoleonHandler{
		groupNapoleonApplication: groupNapoleonApplication,
	}
}

func (h *GroupNapoleonHandler) CreateGroupNapoEmail(c *fiber.Ctx) error {
	var request groupnapoleoncontact.ContactGroupNapoleonRequest
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

	newMessage := request.Message + " - My cell phone number (with WhatsApp or Telegram) is: " + request.Phone + " - My nickname in Napoleon is: " + request.Nickname
	subject := "New Message of Contact"
	data := sendemail.CreateBodyEmail("Grupo Napoleon", request.Name, request.Email, app.EmailGroupNapoleonTo(), subject, newMessage, time.Now().Year())
	htmlContent, err := sendemail.GenerateContactEmailHTML(*data)
	if err != nil {
		logger.LogError("error generating email content", err, logger.LogStruct{Action: "generate_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating email content",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	contact, err := h.groupNapoleonApplication.CreateNapoGroupEmail(ctx, request.Email, request.Name, app.EmailGroupNapoleonTo(), subject, newMessage, htmlContent)
	if err != nil {
		logger.LogError("error creating Group Napoleon email", err, logger.LogStruct{Action: "create_group_napoleon_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating Group Napoleon email",
		})
	}

	err = sendemail.SendContactEmailHTML(htmlContent, app.EmailGroupNapoleonTo(), subject)
	if err != nil {
		logger.LogError("error sending email", err, logger.LogStruct{Action: "send_email_error", Data: app.EmailGroupNapoleonTo()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error sending email",
		})
	}

	sent := false
	contact.Sending = &sent
	_, err = h.groupNapoleonApplication.UpdateNapoGroupEmail(ctx, contact)
	if err != nil {
		logger.LogError("error updating Group Napoleon email sending status", err, logger.LogStruct{Action: "update_group_napoleon_email_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Group Napoleon email sent successfully",
	})
}
