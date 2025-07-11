package models

import "time"

type Post struct {
	Title       string    `yaml:"title"`
	Date        string    `yaml:"date"`
	Slug        string    `yaml:"slug"`
	Description string    `yaml:"description"`
	Tags        []string  `yaml:"tags"`
	Published   bool      `yaml:"published"`
	Content     string    // HTML content after markdown processing
	ParsedDate  time.Time // Parsed date for sorting
}
