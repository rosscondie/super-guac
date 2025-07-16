package data

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetAllPhotos() ([]Photo, error) {
	var photos []Photo

	err := filepath.WalkDir("content/photos", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil // skip directories
		}

		// Filter for image extensions (case-insensitive)
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return nil // skip non-image files
		}

		photos = append(photos, Photo{
			Filename: d.Name(),
			URL:      "/photos/" + d.Name(),
			Size:     getFileSize(path),
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to read photo folder: %w", err)
	}

	return photos, nil
}

func getFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}
