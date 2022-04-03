package files

import (
	"context"
	"fmt"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/service"
	"github.com/artomsopun/clendry/clendry-api/pkg/logger"
	"github.com/artomsopun/clendry/clendry-api/pkg/storage"
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
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

func NewFilesService(storage storage.Provider) *FilesManager {
	return &FilesManager{storage: storage}
}

func (s *FilesManager) GetObjectLink(fileName string) string {
	return s.storage.GetLink(fileName)
}

func (s *FilesManager) RemoveObject(ctx context.Context, object string) error {
	return s.storage.Delete(ctx, object)
}

func (s *FilesManager) UploadObject(ctx context.Context, folder string, file service.File) (string, error) {
	f, err := os.Open(file.Title)
	if err != nil {
		return "", err
	}

	info, _ := f.Stat()
	logger.Infof("file info: %+v", info)

	defer f.Close()

	fileName := s.generateFilename(folder, file)
	err = s.storage.Upload(ctx, storage.UploadInput{
		File:        f,
		Size:        file.Size,
		ContentType: file.ContentType,
		Name:        fileName,
	})
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (s *FilesManager) generateFilename(folder string, file service.File) string {
	fileName := fmt.Sprintf("%s.%s", uuid.New().String(), getFileExtension(file.Title))
	folderType := folders[file.Type]

	fileNameParts := strings.Split(file.Title, "-") // first part is userID

	return fmt.Sprintf("%s/%s/%s/%s", folder, fileNameParts[0], folderType, fileName)
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")

	return parts[len(parts)-1]
}

func (s *FilesManager) RemoveFile(filename string) {
	if err := os.Remove(filename); err != nil {
		logger.Error("removeFile(): ", err)
	}
}
