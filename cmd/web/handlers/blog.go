package handlers

import (
	"net/http"
	"strings"

	"github.com/abbesm0hamed/portfolio/cmd/repository"
	"github.com/abbesm0hamed/portfolio/cmd/web/views"
)

type BlogHandler struct {
	repo repository.BlogRepository
}

func NewBlogHandler(repo repository.BlogRepository) *BlogHandler {
	return &BlogHandler{
		repo: repo,
	}
}

func (h *BlogHandler) BlogListPage(w http.ResponseWriter, r *http.Request) {
	posts, err := h.repo.GetAllPosts()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	views.BlogList(posts).Render(r.Context(), w)
}

func (h *BlogHandler) BlogPostPage(w http.ResponseWriter, r *http.Request) {
	// Extract slug from URL path
	path := strings.TrimPrefix(r.URL.Path, "/blog/")
	if path == "" {
		http.Redirect(w, r, "/blog", http.StatusSeeOther)
		return
	}

	post, err := h.repo.GetPostBySlug(path)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	views.BlogPost(post).Render(r.Context(), w)
}
