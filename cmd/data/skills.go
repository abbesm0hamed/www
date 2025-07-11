package data

import (
	"github.com/abbesm0hamed/portfolio/cmd/models"
)

var Skills = []models.Skill{
	{
		ID:           1,
		Name:         "Go",
		Category:     "Backend",
		Level:        "advanced",
		Description:  "Proficient in building scalable backend services, APIs, and microservices using Go. Experience with goroutines, channels, and Go's standard library.",
		YearsOfExp:   3,
		Technologies: []string{"Gin", "Fiber", "GORM", "Go Kit"},
	},
	{
		ID:           2,
		Name:         "React",
		Category:     "Frontend",
		Level:        "advanced",
		Description:  "Extensive experience in building modern web applications using React. Skilled in hooks, state management, and component architecture.",
		YearsOfExp:   4,
		Technologies: []string{"Redux", "Context API", "React Router", "Material-UI"},
	},
	{
		ID:           3,
		Name:         "PostgreSQL",
		Category:     "Database",
		Level:        "intermediate",
		Description:  "Experience in designing and optimizing database schemas, writing complex queries, and database administration.",
		YearsOfExp:   3,
		Technologies: []string{"SQL", "PL/pgSQL", "Database Design", "Performance Tuning"},
	},
	{
		ID:           4,
		Name:         "Docker",
		Category:     "DevOps",
		Level:        "intermediate",
		Description:  "Containerization of applications, multi-stage builds, and orchestration with Docker Compose.",
		YearsOfExp:   2,
		Technologies: []string{"Docker Compose", "Dockerfile", "Container Registry"},
	},
	{
		ID:           5,
		Name:         "JavaScript",
		Category:     "Frontend",
		Level:        "expert",
		Description:  "Deep understanding of JavaScript fundamentals, ES6+ features, and modern JavaScript development patterns.",
		YearsOfExp:   5,
		Technologies: []string{"ES6+", "TypeScript", "Node.js", "Webpack"},
	},
	{
		ID:           6,
		Name:         "AWS",
		Category:     "Cloud",
		Level:        "intermediate",
		Description:  "Experience with AWS services for building and deploying cloud-native applications.",
		YearsOfExp:   2,
		Technologies: []string{"EC2", "S3", "RDS", "Lambda", "CloudFormation"},
	},
}

func GetSkills() []models.Skill {
	return Skills
}

func GetSkillByID(id int) *models.Skill {
	for _, skill := range Skills {
		if skill.ID == id {
			return &skill
		}
	}
	return nil
}

func GetSkillsByCategory(category string) []models.Skill {
	var filteredSkills []models.Skill
	for _, skill := range Skills {
		if skill.Category == category {
			filteredSkills = append(filteredSkills, skill)
		}
	}
	return filteredSkills
}

func GetSkillsByLevel(level string) []models.Skill {
	var filteredSkills []models.Skill
	for _, skill := range Skills {
		if skill.Level == level {
			filteredSkills = append(filteredSkills, skill)
		}
	}
	return filteredSkills
}
