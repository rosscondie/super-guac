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
	c := cases.Title(language.English)

	root := "content/photos"
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		slug := entry.Name()
		albumPath := filepath.Join(root, slug)
		metadataPath := filepath.Join(albumPath, "album.json")

		album := Album{
			Slug: slug,
		}

		// Try to read album.json
		if data, err := os.ReadFile(metadataPath); err == nil {
			var meta AlbumMetadata
			if err := json.Unmarshal(data, &meta); err == nil {
				// Use values from album.json
				if meta.Title != "" {
					album.Name = meta.Title
				}
				if meta.Cover != "" {
					album.Cover = meta.Cover
				}
			}
		}

		// Fallback if Name is still empty
		if album.Name == "" {
			album.Name = c.String(strings.ReplaceAll(slug, "-", " "))
		}

		// Fallback if Cover is still empty
		if album.Cover == "" {
			files, _ := os.ReadDir(albumPath)
			for _, f := range files {
				if !f.IsDir() {
					ext := strings.ToLower(filepath.Ext(f.Name()))
					if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
						album.Cover = fmt.Sprintf("/photos/%s/%s", slug, f.Name())
						break
					}
				}
			}
		}

		albums = append(albums, album)
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

func GetAlbumMetadata(slug string) (AlbumMetadata, error) {
	var meta AlbumMetadata
	albumPath := filepath.Join("content/photos", slug)
	metadataPath := filepath.Join(albumPath, "album.json")

	if _, err := os.Stat(albumPath); os.IsNotExist(err) {
		return meta, fmt.Errorf("album does not exist")
	}

	if bytes, err := os.ReadFile(metadataPath); err == nil {
		_ = json.Unmarshal(bytes, &meta)
	}

	// Fallbacks for missing fields
	if meta.Title == "" {
		meta.Title = strings.Title(strings.ReplaceAll(slug, "-", " "))
	}

	if meta.Slug == "" {
		meta.Slug = slug
	}

	if meta.Description == "" {
		meta.Description = ""
	}

	// If needed, default cover to empty string instead of crashing frontend
	if meta.Cover == "" {
		meta.Cover = ""
	}

	return meta, nil
}

// UpdateAlbumMetadata updates (or creates) the album.json file in the album folder
func UpdateAlbumMetadata(slug string, metadata AlbumMetadata) error {
	albumDir := filepath.Join("content", "photos", slug)

	// Ensure the album folder exists
	if _, err := os.Stat(albumDir); os.IsNotExist(err) {
		return fmt.Errorf("album folder %s does not exist", slug)
	}

	// Normalize and validate cover image
	if metadata.Cover != "" {
		filename := filepath.Base(metadata.Cover)

		files, err := os.ReadDir(albumDir)
		if err != nil {
			return fmt.Errorf("failed to read album folder: %v", err)
		}

		var found bool
		for _, file := range files {
			if !file.IsDir() && strings.EqualFold(file.Name(), filename) {
				// Use the actual casing from disk
				metadata.Cover = fmt.Sprintf("/photos/%s/%s", slug, file.Name())
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("cover image %s does not exist in album %s", filename, slug)
		}
	}

	// Marshal metadata into JSON
	jsonBytes, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal album metadata: %v", err)
	}

	// Write album.json to disk
	albumFilePath := filepath.Join(albumDir, "album.json")
	if err := os.WriteFile(albumFilePath, jsonBytes, 0644); err != nil {
		return fmt.Errorf("failed to write album.json: %v", err)
	}

	return nil
}
