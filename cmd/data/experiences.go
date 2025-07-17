package data

import (
	"time"

	"github.com/abbesm0hamed/www/cmd/models"
)

var Experiences = []models.Experience{
	{
		ID:          1,
		Title:       "Senior Software Engineer",
		Company:     "Tech Solutions Inc.",
		Location:    "San Francisco, CA",
		StartDate:   time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
		EndDate:     nil,
		Description: "Lead development of scalable web applications using Go and React. Mentor junior developers and collaborate with cross-functional teams to deliver high-quality software solutions.",
		Technologies: []string{"Go", "React", "PostgreSQL", "Docker", "Kubernetes"},
		Current:     true,
	},
	{
		ID:        2,
		Title:     "Full Stack Developer",
		Company:   "StartupXYZ",
		Location:  "Remote",
		StartDate: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   &[]time.Time{time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)}[0],
		Description: "Built and maintained multiple web applications from conception to deployment. Worked closely with designers and product managers to implement user-friendly features.",
		Technologies: []string{"JavaScript", "Node.js", "Vue.js", "MongoDB", "AWS"},
		Current:      false,
	},
	{
		ID:        3,
		Title:     "Junior Software Developer",
		Company:   "WebDev Agency",
		Location:  "New York, NY",
		StartDate: time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   &[]time.Time{time.Date(2021, 5, 31, 0, 0, 0, 0, time.UTC)}[0],
		Description: "Developed responsive websites and web applications for various clients. Gained experience in modern web technologies and agile development practices.",
		Technologies: []string{"HTML", "CSS", "JavaScript", "PHP", "MySQL"},
		Current:      false,
	},
}

func GetExperiences() []models.Experience {
	return Experiences
}

func GetExperienceByID(id int) *models.Experience {
	for _, exp := range Experiences {
		if exp.ID == id {
			return &exp
		}
	}
	return nil
}
