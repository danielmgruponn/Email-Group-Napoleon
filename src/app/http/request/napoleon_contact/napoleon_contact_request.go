package napoleoncontact

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

type ContactNapoleonRequest struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (r *ContactNapoleonRequest) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	r.Name = strings.TrimSpace(r.Name)
	r.Subject = strings.TrimSpace(r.Subject)
	r.Message = strings.TrimSpace(r.Message)

	if r.Email == "" || r.Name == "" || r.Subject == "" || r.Message == "" {
		return fmt.Errorf("all fields are required")
	}
	if !emailRegex.MatchString(r.Email) {
		return fmt.Errorf("invalid email format")
	}
	if len(r.Name) > 100 {
		return fmt.Errorf("name must be 100 characters or less")
	}
	if len(r.Subject) > 200 {
		return fmt.Errorf("subject must be 200 characters or less")
	}
	if len(r.Message) > 5000 {
		return fmt.Errorf("message must be 5000 characters or less")
	}
	return nil
}
