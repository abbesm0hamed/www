package server

import (
	"net/http"

	"github.com/abbesm0hamed/www/cmd/repository"
	"github.com/abbesm0hamed/www/cmd/services"
	"github.com/abbesm0hamed/www/cmd/web"
	"github.com/abbesm0hamed/www/cmd/web/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Initialize repositories
	experienceRepo := repository.NewInMemoryExperienceRepository()
	projectRepo := repository.NewInMemoryProjectRepository()
	blogRepo := repository.NewInMemoryBlogRepository()

	// Initialize services
	emailService := services.NewEmailService()

	// Initialize handlers
	experienceHandler := handlers.NewExperienceHandler(experienceRepo)
	projectHandler := handlers.NewProjectHandler(projectRepo)
	contactHandler := handlers.NewContactHandler(emailService)
	blogHandler := handlers.NewBlogHandler(blogRepo)

	// Register routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			// Home page
			handlers.HomeHandler(w, r)
		} else {
			// Any other unmatched path → 404
			handlers.NotFoundHandler(w, r)
		}
	})

	mux.HandleFunc("/about", handlers.AboutPage)
	mux.HandleFunc("/contact", contactHandler.ContactPage)
	mux.HandleFunc("/contact/submit", contactHandler.HandleContactForm)
	mux.HandleFunc("/experiences", experienceHandler.ExperiencesPage)
	mux.HandleFunc("/projects", projectHandler.ProjectsPage)
	// Blog routes
	mux.HandleFunc("/blog", blogHandler.BlogListPage)
	mux.HandleFunc("/blog/", blogHandler.BlogPostPage)

	// Static files
	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)

	return s.CorsMiddleware(mux)
}
