package files

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/logger"
	"github.com/artomsopun/clendry/clendry-api/pkg/storage"
	"github.com/google/uuid"
)

const (
	_workersCount   = 2
	_workerInterval = time.Second * 10
)

var folders = map[domain.FileType]string{
	domain.Image: "images",
	domain.Video: "videos",
	domain.Other: "other",
}

type FilesManager struct {
	storage storage.Provider
}

func NewFilesManager(storage storage.Provider) *FilesManager {
	return &FilesManager{storage: storage}
}

func (s *FilesManager) GetObjectLink(fileName string) string {
	return s.storage.GetLink(fileName)
}

func (s *FilesManager) RemoveObject(ctx context.Context, object string) error {
	return s.storage.Delete(ctx, object)
}

func (s *FilesManager) UploadObject(ctx context.Context, userID string, folder string, file File) (string, string, error) {
	f, err := os.Open(file.Title)
	if err != nil {
		return "", "", err
	}

	info, _ := f.Stat()
	logger.Infof("file info: %+v", info)

	defer f.Close()

	fileTitle, fileURL := s.generateFilename(userID, folder, file)
	err = s.storage.Upload(ctx, storage.UploadInput{
		File:        f,
		Size:        file.Size,
		ContentType: file.ContentType,
		Name:        fileURL,
	})
	if err != nil {
		return "", "", err
	}
	return fileTitle, fileURL, nil
}

func (s *FilesManager) generateFilename(userID string, folder string, file File) (string, string) {
	fileName := fmt.Sprintf("%s-%s.%s", getFileTitle(file.Title), uuid.New().String(), getFileExtension(file.Title))
	folderType := folders[file.Type]

	return fileName, fmt.Sprintf("%s/%s/%s/%s", folder, userID, folderType, fileName)
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")

	return parts[len(parts)-1]
}

func getFileTitle(filename string) string {
	parts := strings.Split(filename, ".")
	title := ""
	for i := 0; i < len(parts)-1; i++ {
		title += parts[i]
	}
	return title
}

func (s *FilesManager) RemoveFile(filename string) {
	if err := os.Remove(filename); err != nil {
		logger.Error("removeFile(): ", err)
	}
}
