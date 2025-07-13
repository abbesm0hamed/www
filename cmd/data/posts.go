package data

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/abbesm0hamed/portfolio/cmd/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

func GetPosts() ([]*models.Post, error) {
	postsDir := "content/posts"
	var posts []*models.Post

	err := filepath.WalkDir(postsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		post, err := parsePost(path)
		if err != nil {
			return fmt.Errorf("failed to parse post %s: %w", path, err)
		}

		posts = append(posts, post)
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort posts by date (newest first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ParsedDate.After(posts[j].ParsedDate)
	})

	return posts, nil
}

func parsePost(filename string) (*models.Post, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Split frontmatter and content
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid frontmatter format")
	}

	// Parse frontmatter
	var post models.Post
	if err := yaml.Unmarshal([]byte(parts[1]), &post); err != nil {
		return nil, err
	}

	// Parse date
	parsedDate, err := time.Parse("2006-01-02", post.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}
	post.ParsedDate = parsedDate

	// Convert markdown to HTML with enhanced configuration
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.TaskList,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithAttribute(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(), // Allow HTML in markdown (be careful with user input)
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(parts[2]), &buf); err != nil {
		return nil, err
	}

	// Post-process HTML to add image classes and handle relative paths
	processedContent := processImagePaths(buf.String())
	post.Content = processedContent

	return &post, nil
}

// processImagePaths handles image path processing and adds proper classes
func processImagePaths(content string) string {
	// Replace relative image paths with absolute paths
	// Images are stored in /cmd/web/assets/images/posts/
	content = strings.ReplaceAll(content, `src="./images/`, `src="/assets/images/posts/`)
	content = strings.ReplaceAll(content, `src="images/`, `src="/assets/images/posts/`)

	// Add responsive classes to images
	content = strings.ReplaceAll(content, `<img`, `<img class="blog-image"`)

	return content
}
