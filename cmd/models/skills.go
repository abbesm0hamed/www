package models

import "fmt"

type Skill struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Category     string   `json:"category"`
	Level        string   `json:"level"`
	Description  string   `json:"description"`
	YearsOfExp   int      `json:"years_of_experience"`
	Technologies []string `json:"technologies,omitempty"`
}

func (s *Skill) GetProficiencyLevel() string {
	switch s.Level {
	case "beginner":
		return "Beginner"
	case "intermediate":
		return "Intermediate"
	case "advanced":
		return "Advanced"
	case "expert":
		return "Expert"
	default:
		return "Unknown"
	}
}

func (s *Skill) GetExperienceRange() string {
	switch {
	case s.YearsOfExp < 1:
		return "Less than 1 year"
	case s.YearsOfExp == 1:
		return "1 year"
	case s.YearsOfExp < 5:
		return fmt.Sprintf("%d years", s.YearsOfExp)
	default:
		return "5+ years"
	}
}
