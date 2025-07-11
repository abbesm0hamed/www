package handlers

import (
	"net/http"
	"github.com/abbesm0hamed/portfolio/cmd/web/views"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	views.NotFound().Render(r.Context(), w)
}
