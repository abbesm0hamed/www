package repository

import (
	"fmt"
	"sort"

	"github.com/abbesm0hamed/www/cmd/data"
	"github.com/abbesm0hamed/www/cmd/models"
)

type ProjectRepository interface {
	GetAll() ([]models.Project, error)
	GetByID(id int) (*models.Project, error)
	GetByStatus(status string) ([]models.Project, error)
	GetByCategory(category string) ([]models.Project, error)
	GetFeatured() ([]models.Project, error)
}

type InMemoryProjectRepository struct{}

func NewInMemoryProjectRepository() *InMemoryProjectRepository {
	return &InMemoryProjectRepository{}
}

func (r *InMemoryProjectRepository) GetAll() ([]models.Project, error) {
	projects := data.GetProjects()
	// Sort by start date (most recent first)
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].StartDate.After(projects[j].StartDate)
	})
	return projects, nil
}

func (r *InMemoryProjectRepository) GetByID(id int) (*models.Project, error) {
	project := data.GetProjectByID(id)
	if project == nil {
		return nil, fmt.Errorf("project with ID %d not found", id)
	}
	return project, nil
}

func (r *InMemoryProjectRepository) GetByStatus(status string) ([]models.Project, error) {
	projects := data.GetProjectsByStatus(status)
	// Sort by start date (most recent first)
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].StartDate.After(projects[j].StartDate)
	})
	return projects, nil
}

func (r *InMemoryProjectRepository) GetByCategory(category string) ([]models.Project, error) {
	projects := data.GetProjectsByCategory(category)
	// Sort by start date (most recent first)
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].StartDate.After(projects[j].StartDate)
	})
	return projects, nil
}

func (r *InMemoryProjectRepository) GetFeatured() ([]models.Project, error) {
	projects := data.GetFeaturedProjects()
	// Sort by start date (most recent first)
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].StartDate.After(projects[j].StartDate)
	})
	return projects, nil
}
