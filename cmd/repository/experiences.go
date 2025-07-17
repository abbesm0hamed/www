package repository

import (
	"fmt"
	"sort"

	"github.com/abbesm0hamed/www/cmd/data"
	"github.com/abbesm0hamed/www/cmd/models"
)

type ExperienceRepository interface {
	GetAll() ([]models.Experience, error)
	GetByID(id int) (*models.Experience, error)
}

type InMemoryExperienceRepository struct{}

func NewInMemoryExperienceRepository() *InMemoryExperienceRepository {
	return &InMemoryExperienceRepository{}
}

func (r *InMemoryExperienceRepository) GetAll() ([]models.Experience, error) {
	experiences := data.GetExperiences()

	// Sort by start date (most recent first)
	sort.Slice(experiences, func(i, j int) bool {
		return experiences[i].StartDate.After(experiences[j].StartDate)
	})

	return experiences, nil
}

func (r *InMemoryExperienceRepository) GetByID(id int) (*models.Experience, error) {
	experience := data.GetExperienceByID(id)
	if experience == nil {
		return nil, fmt.Errorf("experience with ID %d not found", id)
	}

	return experience, nil
}
