package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GetAllAlbums reads the content/photos directory and returns a list of albums with their cover images
func GetAllAlbums() ([]Album, error) {
	var albums []Album
	c := cases.Title(language.English) // create title casing transformer

	entries, err := os.ReadDir("content/photos")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		albumName := entry.Name()
		albumPath := filepath.Join("content/photos", albumName)

		files, _ := os.ReadDir(albumPath)

		var cover string
		for _, f := range files {
			if !f.IsDir() {
				ext := strings.ToLower(filepath.Ext(f.Name()))
				if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
					cover = f.Name()
					break
				}
			}
		}

		albums = append(albums, Album{
			Name:  c.String(strings.ReplaceAll(albumName, "-", " ")),
			Slug:  albumName,
			Cover: fmt.Sprintf("/photos/%s/%s", albumName, cover),
		})
	}

	return albums, nil
}

// GetPhotosByAlbum returns all photos inside a given album folder
func GetPhotosByAlbum(albumSlug string) ([]Photo, error) {
	var photos []Photo
	albumPath := filepath.Join("content/photos", albumSlug)

	files, err := os.ReadDir(albumPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read album folder: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			continue
		}

		fullPath := filepath.Join(albumPath, file.Name())
		photos = append(photos, Photo{
			Filename: file.Name(),
			URL:      fmt.Sprintf("/photos/%s/%s", albumSlug, file.Name()),
			Size:     getFileSize(fullPath),
		})
	}

	return photos, nil
}

func CreateAlbum(album Album) error {
	albumDir := filepath.Join("content/photos", album.Slug)
	if _, err := os.Stat(albumDir); err == nil {
		return fmt.Errorf("album already exists")
	}

	if err := os.MkdirAll(albumDir, 0755); err != nil {
		return fmt.Errorf("failed to create album directory: %w", err)
	}

	// Write metadata file
	metaPath := filepath.Join(albumDir, "album.json")
	metaBytes, err := json.MarshalIndent(album, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}
	if err := os.WriteFile(metaPath, metaBytes, 0644); err != nil {
		return fmt.Errorf("failed to write metadata: %w", err)
	}

	return nil
}

func DeleteAlbum(slug string) error {
	albumPath := filepath.Join("content", "photos", slug)

	// Check if album exists
	if _, err := os.Stat(albumPath); os.IsNotExist(err) {
		return fmt.Errorf("album not found")
	}

	// Delete the folder and all its contents
	err := os.RemoveAll(albumPath)
	if err != nil {
		return fmt.Errorf("failed to delete album: %v", err)
	}

	return nil
}

func DeletePhotoFromAlbum(slug, filename string) error {
	photoPath := filepath.Join("content/photos", slug, filename)

	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		return fmt.Errorf("photo not found")
	}

	if err := os.Remove(photoPath); err != nil {
		return fmt.Errorf("failed to delete photo: %w", err)
	}

	return nil
}
