package bankgoldcontact

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

type ContactBackGoldRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

func (r *ContactBackGoldRequest) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	r.Name = strings.TrimSpace(r.Name)
	r.Nickname = strings.TrimSpace(r.Nickname)
	r.Phone = strings.TrimSpace(r.Phone)

	if r.Email == "" || r.Name == "" || r.Phone == "" || r.Nickname == "" {
		return fmt.Errorf("all fields are required")
	}
	if !emailRegex.MatchString(r.Email) {
		return fmt.Errorf("invalid email format")
	}
	if len(r.Name) > 100 {
		return fmt.Errorf("name must be 100 characters or less")
	}
	if len(r.Nickname) > 50 {
		return fmt.Errorf("nickname must be 50 characters or less")
	}
	if len(r.Phone) > 30 {
		return fmt.Errorf("phone must be 30 characters or less")
	}
	return nil
}
