package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/abbesm0hamed/www/cmd/models"
	"github.com/abbesm0hamed/www/cmd/services"
	"github.com/abbesm0hamed/www/cmd/web/views"
)

type ContactHandler struct {
	emailService *services.EmailService
}

func NewContactHandler(emailService *services.EmailService) *ContactHandler {
	return &ContactHandler{
		emailService: emailService,
	}
}

func (h *ContactHandler) ContactPage(w http.ResponseWriter, r *http.Request) {
	component := views.ContactPage()
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering contact page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *ContactHandler) HandleContactForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %v", err)
		h.renderErrorResponse(w, r, "Invalid form data")
		return
	}

	// Create contact request
	contactRequest := &models.ContactRequest{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
		SentAt:  time.Now(),
	}

	// Validate the form
	if errors := contactRequest.Validate(); len(errors) > 0 {
		h.renderValidationErrors(w, r, errors)
		return
	}

	// Send email
	if err := h.emailService.SendContactEmail(contactRequest); err != nil {
		log.Printf("Error sending email: %v", err)
		h.renderErrorResponse(w, r, "Failed to send message. Please try again later.")
		return
	}

	// Return success response
	h.renderSuccessResponse(w, r)
}

func (h *ContactHandler) renderValidationErrors(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	w.Header().Set("Content-Type", "text/html")
	component := views.ContactFormErrors(errors)
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering validation errors: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *ContactHandler) renderErrorResponse(w http.ResponseWriter, r *http.Request, message string) {
	w.Header().Set("Content-Type", "text/html")
	component := views.ContactFormError(message)
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering error response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *ContactHandler) renderSuccessResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	component := views.ContactFormSuccess()
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering success response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
