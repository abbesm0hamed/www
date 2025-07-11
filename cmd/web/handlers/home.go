package handlers

import (
	"net/http"

	"github.com/abbesm0hamed/portfolio/cmd/web/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
