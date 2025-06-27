package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abbesm0hamed/portfolio/cmd/repository"
	"github.com/abbesm0hamed/portfolio/cmd/web"
	"github.com/abbesm0hamed/portfolio/cmd/web/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	experienceRepo := repository.NewInMemoryExperienceRepository()
	experienceHandler := handlers.NewExperienceHandler(experienceRepo)

	// Register routes
	mux.HandleFunc("/", s.portfolioHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)
	mux.HandleFunc("/experiences", experienceHandler.ExperiencesPage)

	// Wrap the mux with CORS middleware
	return s.CorsMiddleware(mux)
}

func (s *Server) portfolioHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Welcome to abbes portfolio"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
