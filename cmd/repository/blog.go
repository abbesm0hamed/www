package repository

import (
	"fmt"

	"github.com/abbesm0hamed/portfolio/cmd/data"
	"github.com/abbesm0hamed/portfolio/cmd/models"
)

type BlogRepository interface {
	GetAllPosts() ([]*models.Post, error)
	GetPostBySlug(slug string) (*models.Post, error)
}

type InMemoryBlogRepository struct {
	posts []*models.Post
}

func NewInMemoryBlogRepository() *InMemoryBlogRepository {
	posts, err := data.LoadPosts()
	if err != nil {
		// Handle error gracefully - return empty repository
		fmt.Printf("Error loading posts: %v\n", err)
		posts = []*models.Post{}
	}

	return &InMemoryBlogRepository{
		posts: posts,
	}
}

func (r *InMemoryBlogRepository) GetAllPosts() ([]*models.Post, error) {
	// Return only published posts
	var publishedPosts []*models.Post
	for _, post := range r.posts {
		if post.Published {
			publishedPosts = append(publishedPosts, post)
		}
	}
	return publishedPosts, nil
}

func (r *InMemoryBlogRepository) GetPostBySlug(slug string) (*models.Post, error) {
	for _, post := range r.posts {
		if post.Slug == slug && post.Published {
			return post, nil
		}
	}
	return nil, fmt.Errorf("post not found")
}
