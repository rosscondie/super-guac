package data

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

func GetPostBySlug(slug string) (*PostContent, error) {
	filename := filepath.Join("content", slug+".md")

	mdBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read post: %w", err)
	}

	content := string(mdBytes)

	if !strings.HasPrefix(content, "---") {
		return nil, fmt.Errorf("missing frontmatter")
	}

	// Split frontmatter and markdown body
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid frontmatter format")
	}

	fmText := parts[1]
	mdBody := parts[2]

	// Parse YAML frontmatter
	var fm frontmatter
	if err := yaml.Unmarshal([]byte(fmText), &fm); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Convert markdown body to HTML
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(mdBody), &buf); err != nil {
		return nil, fmt.Errorf("markdown render failed: %w", err)
	}

	return &PostContent{
		Title: fm.Title,
		HTML:  buf.String(),
		Date:  fm.Date,
	}, nil
}
