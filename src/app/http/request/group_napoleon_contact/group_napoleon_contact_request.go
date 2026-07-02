package groupnapoleoncontact

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

type ContactGroupNapoleonRequest struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
	Nickname string `json:"nickname"`
}

func (r *ContactGroupNapoleonRequest) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	r.Name = strings.TrimSpace(r.Name)
	r.Phone = strings.TrimSpace(r.Phone)
	r.Message = strings.TrimSpace(r.Message)
	r.Nickname = strings.TrimSpace(r.Nickname)

	if r.Email == "" || r.Name == "" || r.Phone == "" || r.Message == "" || r.Nickname == "" {
		return fmt.Errorf("all fields are required")
	}
	if !emailRegex.MatchString(r.Email) {
		return fmt.Errorf("invalid email format")
	}
	if len(r.Name) > 100 {
		return fmt.Errorf("name must be 100 characters or less")
	}
	if len(r.Phone) > 30 {
		return fmt.Errorf("phone must be 30 characters or less")
	}
	if len(r.Message) > 5000 {
		return fmt.Errorf("message must be 5000 characters or less")
	}
	return nil
}
