package services

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/abbesm0hamed/portfolio/cmd/models"
)

type EmailService struct {
	smtpHost     string
	smtpPort     string
	smtpUser     string
	smtpPassword string
	fromEmail    string
	toEmail      string
}

func NewEmailService() *EmailService {
	return &EmailService{
		smtpHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		smtpPort:     getEnv("SMTP_PORT", "587"),
		smtpUser:     getEnv("SMTP_USER", "your-email@gmail.com"),
		smtpPassword: getEnv("SMTP_PASSWORD", "your-app-password"),
		fromEmail:    getEnv("FROM_EMAIL", "your-email@gmail.com"),
		toEmail:      getEnv("TO_EMAIL", "your-email@gmail.com"),
	}
}

func (e *EmailService) SendContactEmail(contact *models.ContactRequest) error {
	// Create the email message
	subject := fmt.Sprintf("Portfolio Contact: %s", contact.Subject)
	body := e.createEmailBody(contact)

	// Create the full message
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		e.fromEmail, e.toEmail, subject, body)

	// SMTP authentication
	auth := smtp.PlainAuth("", e.smtpUser, e.smtpPassword, e.smtpHost)

	// Send the email
	err := smtp.SendMail(
		e.smtpHost+":"+e.smtpPort,
		auth,
		e.fromEmail,
		[]string{e.toEmail},
		[]byte(message),
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (e *EmailService) createEmailBody(contact *models.ContactRequest) string {
	return fmt.Sprintf(`
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
			.container { max-width: 600px; margin: 0 auto; padding: 20px; }
			.header { background-color: #10b981; color: white; padding: 20px; text-align: center; }
			.content { padding: 20px; background-color: #f9f9f9; }
			.field { margin-bottom: 15px; }
			.label { font-weight: bold; color: #10b981; }
			.value { margin-top: 5px; }
			.footer { text-align: center; padding: 20px; color: #666; font-size: 12px; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>New Contact Form Submission</h1>
			</div>
			<div class="content">
				<div class="field">
					<div class="label">Name:</div>
					<div class="value">%s</div>
				</div>
				<div class="field">
					<div class="label">Email:</div>
					<div class="value">%s</div>
				</div>
				<div class="field">
					<div class="label">Subject:</div>
					<div class="value">%s</div>
				</div>
				<div class="field">
					<div class="label">Message:</div>
					<div class="value">%s</div>
				</div>
				<div class="field">
					<div class="label">Sent At:</div>
					<div class="value">%s</div>
				</div>
			</div>
			<div class="footer">
				<p>This email was sent from your portfolio contact form.</p>
			</div>
		</div>
	</body>
	</html>
	`, contact.Name, contact.Email, contact.Subject,
		strings.ReplaceAll(contact.Message, "\n", "<br>"),
		contact.SentAt.Format("January 2, 2006 at 3:04 PM"))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
