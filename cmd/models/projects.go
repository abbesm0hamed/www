package models

import "time"

type Project struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Status       string     `json:"status"`
	Technologies []string   `json:"technologies"`
	GitHubURL    string     `json:"github_url,omitempty"`
	LiveURL      string     `json:"live_url,omitempty"`
	Category     string     `json:"category"`
	Featured     bool       `json:"featured"`
}

func (p *Project) GetDuration() string {
	if p.Status == "in_progress" {
		return "In Progress"
	}
	if p.EndDate != nil {
		return p.EndDate.Format("Jan 2006")
	}
	return ""
}

func (p *Project) GetStartDate() string {
	return p.StartDate.Format("Jan 2006")
}

func (p *Project) GetStatusDisplay() string {
	switch p.Status {
	case "completed":
		return "Completed"
	case "in_progress":
		return "In Progress"
	case "on_hold":
		return "On Hold"
	case "cancelled":
		return "Cancelled"
	default:
		return "Unknown"
	}
}

func (p *Project) IsActive() bool {
	return p.Status == "in_progress" || p.Status == "completed"
}
