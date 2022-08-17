package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

// ImageStore is for blob data storage.
type ImageStore interface {
	// Save puts a laptop image into the store
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error)
}

type DiskImageStore struct {
	mu       sync.Mutex
	imageDir string
	images   map[string]*ImageInfo
}

// ImageInfo contains information of the laptop image
type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
}

func NewDiskImageStore(imgDir string) *DiskImageStore {
	return &DiskImageStore{
		imageDir: imgDir,
		images:   make(map[string]*ImageInfo),
	}
}

func (s *DiskImageStore) Save(laptopId, imgType string, imgData bytes.Buffer) (string, error) {
	imgID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}
	imgPath := fmt.Sprintf("%s/%s%s", s.imageDir, imgID, imgType)
	file, err := os.Create(imgPath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}
	if _, err := imgData.WriteTo(file); err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.images[imgID.String()] = &ImageInfo{
		LaptopID: laptopId,
		Type:     imgType,
		Path:     imgPath,
	}
	return imgID.String(), nil
}
