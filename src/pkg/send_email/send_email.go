package sendemail

import (
	"bytes"
	"html/template"
	napoleontemplate "napoleon-email/src/app/resource/template/napoleon"
	"napoleon-email/src/config/app"
	"napoleon-email/src/pkg/logger"
	"napoleon-email/src/pkg/parse"

	"gopkg.in/gomail.v2"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("contacto").Parse(napoleontemplate.ContactoHTML))
}

type ContactEmailHTML struct {
	Company string
	Year    int
	Name    string
	Email   string
	To      string
	Subject string
	Message string
}

func CreateBodyEmail(company, name, email, to, subject, message string, year int) *ContactEmailHTML {
	return &ContactEmailHTML{
		Company: company,
		Name:    name,
		Email:   email,
		To:      to,
		Subject: subject,
		Message: message,
		Year:    year,
	}
}

func GenerateContactEmailHTML(data ContactEmailHTML) (string, error) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendContactEmailHTML(htmlContent, destiny, subject string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", app.EmailAddress())
	m.SetHeader("To", destiny)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlContent)

	port, err := parse.StringToInt(app.EmailPort())
	if err != nil {
		return err
	}

	d := gomail.NewDialer(app.EmailHost(), port, app.EmailAddress(), app.EmailPassword())
	if err := d.DialAndSend(m); err != nil {
		logger.LogError("Error sending email", err, logger.LogStruct{Action: "send_email_error", User: 0, Data: destiny})
		return err
	}

	return nil
}
