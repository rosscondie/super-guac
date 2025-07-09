package data

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

type PostContent struct {
	Title string `json:"title"`
	HTML  string `json:"html"`
}

func GetPostBySlug(slug string) (*PostContent, error) {
	// Build the path e.g. "content/first-post.md"
	filename := filepath.Join("content", slug+".md")

	// Read the file
	mdBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read post: %w", err)
	}

	// Convert Markdown to HTML
	var buf bytes.Buffer
	err = goldmark.Convert(mdBytes, &buf)
	if err != nil {
		return nil, fmt.Errorf("markdown render failed %w", err)
	}

	return &PostContent{
		Title: slug,
		HTML:  buf.String(),
	}, nil
}
