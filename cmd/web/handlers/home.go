package handlers

import (
	"net/http"

	"github.com/abbesm0hamed/portfolio/cmd/data"
	"github.com/abbesm0hamed/portfolio/cmd/models"
	"github.com/abbesm0hamed/portfolio/cmd/web/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	experiences := data.GetExperiences()
	skills := data.GetSkills()
	projects := data.GetProjects()
	posts, err := data.GetPosts()
	if err != nil {
		// Handle error - for simplicity, we'll log it and continue with an empty slice
		posts = []*models.Post{}
	}

	views.Home(experiences, skills, projects, posts).Render(r.Context(), w)
}
