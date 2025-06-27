package handlers

import (
	"log"
	"net/http"

	"github.com/abbesm0hamed/portfolio/cmd/repository"
	"github.com/abbesm0hamed/portfolio/cmd/web/views"
)

type ExperienceHandler struct {
	repo repository.ExperienceRepository
}

func NewExperienceHandler(repo repository.ExperienceRepository) *ExperienceHandler {
	return &ExperienceHandler{
		repo: repo,
	}
}

func (h *ExperienceHandler) ExperiencesPage(w http.ResponseWriter, r *http.Request) {
	exps, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching experiences: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	component := views.ExperiencesPage(exps)
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering experiences page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
