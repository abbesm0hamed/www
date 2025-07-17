package handlers

import (
	"log"
	"net/http"

	"github.com/abbesm0hamed/www/cmd/web/views"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {
	component := views.AboutPage()
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering about page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
