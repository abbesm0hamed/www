package models

import "time"

type Experience struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Description string    `json:"description"`
	Technologies []string  `json:"technologies"`
	Current     bool      `json:"current"`
}

func (e *Experience) GetDuration() string {
	if e.Current {
		return "Present"
	}
	if e.EndDate != nil {
		return e.EndDate.Format("Jan 2006")
	}
	return ""
}

func (e *Experience) GetStartDate() string {
	return e.StartDate.Format("Jan 2006")
}
