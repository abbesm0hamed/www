package repository

import (
	"fmt"
	"sort"

	"github.com/abbesm0hamed/www/cmd/data"
	"github.com/abbesm0hamed/www/cmd/models"
)

type SkillRepository interface {
	GetAll() ([]models.Skill, error)
	GetByID(id int) (*models.Skill, error)
	GetByCategory(category string) ([]models.Skill, error)
	GetByLevel(level string) ([]models.Skill, error)
}

type InMemorySkillRepository struct{}

func NewInMemorySkillRepository() *InMemorySkillRepository {
	return &InMemorySkillRepository{}
}

func (r *InMemorySkillRepository) GetAll() ([]models.Skill, error) {
	skills := data.GetSkills()
	// Sort by category and then by level (expert -> beginner)
	sort.Slice(skills, func(i, j int) bool {
		if skills[i].Category != skills[j].Category {
			return skills[i].Category < skills[j].Category
		}
		levelOrder := map[string]int{
			"expert":       0,
			"advanced":     1,
			"intermediate": 2,
			"beginner":     3,
		}
		return levelOrder[skills[i].Level] < levelOrder[skills[j].Level]
	})
	return skills, nil
}

func (r *InMemorySkillRepository) GetByID(id int) (*models.Skill, error) {
	skill := data.GetSkillByID(id)
	if skill == nil {
		return nil, fmt.Errorf("skill with ID %d not found", id)
	}
	return skill, nil
}

func (r *InMemorySkillRepository) GetByCategory(category string) ([]models.Skill, error) {
	skills := data.GetSkillsByCategory(category)
	// Sort by level (expert -> beginner)
	sort.Slice(skills, func(i, j int) bool {
		levelOrder := map[string]int{
			"expert":       0,
			"advanced":     1,
			"intermediate": 2,
			"beginner":     3,
		}
		return levelOrder[skills[i].Level] < levelOrder[skills[j].Level]
	})
	return skills, nil
}

func (r *InMemorySkillRepository) GetByLevel(level string) ([]models.Skill, error) {
	skills := data.GetSkillsByLevel(level)
	// Sort by category
	sort.Slice(skills, func(i, j int) bool {
		return skills[i].Category < skills[j].Category
	})
	return skills, nil
}
