package handlers

import (
	"log"
	"net/http"

	"github.com/abbesm0hamed/www/cmd/repository"
	"github.com/abbesm0hamed/www/cmd/web/views"
)

type ProjectHandler struct {
	repo repository.ProjectRepository
}

func NewProjectHandler(repo repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{
		repo: repo,
	}
}

func (h *ProjectHandler) ProjectsPage(w http.ResponseWriter, r *http.Request) {
	exps, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching projects: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	component := views.ProjectsPage(exps)
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering projects page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
