package data

import (
	"time"

	"github.com/abbesm0hamed/portfolio/cmd/models"
)

var Projects = []models.Project{
	{
		ID:           1,
		Title:        "E-commerce Platform",
		Description:  "A full-stack e-commerce platform with user authentication, payment processing, and inventory management. Built with microservices architecture.",
		StartDate:    time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      &[]time.Time{time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)}[0],
		Status:       "completed",
		Technologies: []string{"Go", "React", "PostgreSQL", "Docker", "AWS", "Stripe"},
		GitHubURL:    "https://github.com/abbesm0hamed/ecommerce-platform",
		LiveURL:      "https://ecommerce-demo.example.com",
		Category:     "Web Application",
		Featured:     true,
	},
	{
		ID:           2,
		Title:        "Task Management API",
		Description:  "RESTful API for task management with team collaboration features. Includes real-time notifications and file attachments.",
		StartDate:    time.Date(2023, 3, 10, 0, 0, 0, 0, time.UTC),
		EndDate:      &[]time.Time{time.Date(2023, 7, 20, 0, 0, 0, 0, time.UTC)}[0],
		Status:       "completed",
		Technologies: []string{"Go", "PostgreSQL", "Redis", "WebSocket", "JWT"},
		GitHubURL:    "https://github.com/abbesm0hamed/task-management-api",
		LiveURL:      "",
		Category:     "API",
		Featured:     false,
	},
	{
		ID:           3,
		Title:        "Real-time Chat Application",
		Description:  "A real-time messaging application with group chats, file sharing, and emoji reactions. Built with WebSocket technology.",
		StartDate:    time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
		EndDate:      nil,
		Status:       "in_progress",
		Technologies: []string{"Go", "React", "WebSocket", "MongoDB", "Docker"},
		GitHubURL:    "https://github.com/abbesm0hamed/chat-app",
		LiveURL:      "",
		Category:     "Web Application",
		Featured:     true,
	},
	{
		ID:           4,
		Title:        "Portfolio Website",
		Description:  "Personal portfolio website showcasing projects, skills, and experience. Built with modern web technologies and responsive design.",
		StartDate:    time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
		EndDate:      &[]time.Time{time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)}[0],
		Status:       "completed",
		Technologies: []string{"Go", "HTML", "CSS", "JavaScript", "Docker"},
		GitHubURL:    "https://github.com/abbesm0hamed/portfolio",
		LiveURL:      "https://abbesm0hamed.dev",
		Category:     "Web Application",
		Featured:     true,
	},
	{
		ID:           5,
		Title:        "Weather Dashboard",
		Description:  "A weather dashboard that displays current weather conditions and forecasts for multiple cities. Includes weather alerts and historical data.",
		StartDate:    time.Date(2022, 8, 15, 0, 0, 0, 0, time.UTC),
		EndDate:      &[]time.Time{time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)}[0],
		Status:       "completed",
		Technologies: []string{"JavaScript", "Vue.js", "Weather API", "Chart.js"},
		GitHubURL:    "https://github.com/abbesm0hamed/weather-dashboard",
		LiveURL:      "https://weather-dashboard-demo.example.com",
		Category:     "Web Application",
		Featured:     false,
	},
}

func GetProjects() []models.Project {
	return Projects
}

func GetProjectByID(id int) *models.Project {
	for _, project := range Projects {
		if project.ID == id {
			return &project
		}
	}
	return nil
}

func GetProjectsByStatus(status string) []models.Project {
	var filteredProjects []models.Project
	for _, project := range Projects {
		if project.Status == status {
			filteredProjects = append(filteredProjects, project)
		}
	}
	return filteredProjects
}

func GetProjectsByCategory(category string) []models.Project {
	var filteredProjects []models.Project
	for _, project := range Projects {
		if project.Category == category {
			filteredProjects = append(filteredProjects, project)
		}
	}
	return filteredProjects
}

func GetFeaturedProjects() []models.Project {
	var featuredProjects []models.Project
	for _, project := range Projects {
		if project.Featured {
			featuredProjects = append(featuredProjects, project)
		}
	}
	return featuredProjects
}
