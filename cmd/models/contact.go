package models

import (
	"time"
)

type ContactRequest struct {
	Name    string    `json:"name" validate:"required,min=2,max=100"`
	Email   string    `json:"email" validate:"required,email"`
	Subject string    `json:"subject" validate:"required,min=5,max=200"`
	Message string    `json:"message" validate:"required,min=10,max=1000"`
	SentAt  time.Time `json:"sent_at"`
}

func (c *ContactRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if len(c.Name) < 2 {
		errors["name"] = "Name must be at least 2 characters long"
	}
	if len(c.Name) > 100 {
		errors["name"] = "Name must be less than 100 characters"
	}

	if c.Email == "" {
		errors["email"] = "Email is required"
	}
	// Basic email validation
	if !isValidEmail(c.Email) {
		errors["email"] = "Please enter a valid email address"
	}

	if len(c.Subject) < 5 {
		errors["subject"] = "Subject must be at least 5 characters long"
	}
	if len(c.Subject) > 200 {
		errors["subject"] = "Subject must be less than 200 characters"
	}

	if len(c.Message) < 10 {
		errors["message"] = "Message must be at least 10 characters long"
	}
	if len(c.Message) > 1000 {
		errors["message"] = "Message must be less than 1000 characters"
	}

	return errors
}

func isValidEmail(email string) bool {
	// Basic email validation
	if len(email) < 5 || len(email) > 254 {
		return false
	}

	atCount := 0
	for _, char := range email {
		if char == '@' {
			atCount++
		}
	}

	return atCount == 1 && email[0] != '@' && email[len(email)-1] != '@'
}
