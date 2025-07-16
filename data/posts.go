package data

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// GetAllPosts reads the content/ directory and parses the frontmatter from Markdown files
func GetAllPosts() []Post {
	var posts []Post

	files, err := os.ReadDir("content/blog-posts")
	if err != nil {
		fmt.Println("Failed to read content directory:", err)
		return posts
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
			continue
		}

		slug := strings.TrimSuffix(file.Name(), ".md")
		fullPath := filepath.Join("content/blog-posts", file.Name())

		content, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Println("Failed to read file:", fullPath, err)
			continue
		}

		// Parse frontmatter
		text := string(content)
		if !strings.HasPrefix(text, "---") {
			continue // skip if no frontmatter
		}

		parts := strings.SplitN(text, "---", 3)
		if len(parts) < 3 {
			continue
		}

		fmText := parts[1]
		var fm frontmatter
		if err := yaml.Unmarshal([]byte(fmText), &fm); err != nil {
			fmt.Println("YAML error in", file.Name(), err)
			continue
		}

		post := Post{
			Title:   fm.Title,
			Slug:    slug,
			Summary: fm.Summary,
			Date:    fm.Date,
		}
		posts = append(posts, post)
	}

	fmt.Println("Loaded", len(posts), "posts from content/blog-posts")
	return posts
}
