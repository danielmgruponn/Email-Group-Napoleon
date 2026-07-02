package bankgoldhandler

import (
	"context"
	bankgold "napoleon-email/src/app/domain/application/bank_gold"
	bankgoldcontact "napoleon-email/src/app/http/request/back_gold_contact"
	"napoleon-email/src/config/app"
	"napoleon-email/src/pkg/logger"
	sendemail "napoleon-email/src/pkg/send_email"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BankGoldHandler struct {
	bankGoldApplication bankgold.ContactBankGoldEmailApplicationInterface
}

func NewBankGoldHandler(bankGoldApplication bankgold.ContactBankGoldEmailApplicationInterface) *BankGoldHandler {
	return &BankGoldHandler{
		bankGoldApplication: bankGoldApplication,
	}
}

func (h *BankGoldHandler) CreateBankGoldEmail(c *fiber.Ctx) error {
	var request bankgoldcontact.ContactBackGoldRequest
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

	newMessage := "Name: " + request.Name + " - " + "Email: " + request.Email + " - " + "Phone: " + request.Phone + " - " + "Nickname Napoleon: " + request.Nickname
	subject := "New Message of Contact Bank Gold"
	data := sendemail.CreateBodyEmail("Bank Gold", request.Name, request.Email, app.EmailBankGoldTo(), subject, newMessage, time.Now().Year())
	htmlContent, err := sendemail.GenerateContactEmailHTML(*data)
	if err != nil {
		logger.LogError("error generating email content", err, logger.LogStruct{Action: "generate_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating email content",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	contact, err := h.bankGoldApplication.CreateBankGoldEmail(ctx, request.Email, request.Name, app.EmailBankGoldTo(), subject, newMessage, htmlContent)
	if err != nil {
		logger.LogError("error creating Bank Gold email", err, logger.LogStruct{Action: "create_bank_gold_email_error"})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating Bank Gold email",
		})
	}

	err = sendemail.SendContactEmailHTML(htmlContent, app.EmailBankGoldTo(), subject)
	if err != nil {
		logger.LogError("error sending email", err, logger.LogStruct{Action: "send_email_error", Data: app.EmailBankGoldTo()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error sending email",
		})
	}

	sent := false
	contact.Sending = &sent
	_, err = h.bankGoldApplication.UpdateBankGoldEmail(ctx, contact)
	if err != nil {
		logger.LogError("error updating Bank Gold email sending status", err, logger.LogStruct{Action: "update_bank_gold_email_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Bank Gold email sent successfully",
	})
}
