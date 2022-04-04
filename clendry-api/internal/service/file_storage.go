package service

import (
	"context"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"time"
)

type FilesService struct {
	repoUser  repository.Users
	repoFiles repository.Files
	files     files.Files
}

func NewFilesService(repoUser repository.Users, repoFiles repository.Files, files files.Files) *FilesService {
	return &FilesService{
		repoUser:  repoUser,
		repoFiles: repoFiles,
		files:     files,
	}
}

func (s *FilesService) GetAllFiles(userID types.BinaryUUID) ([]File, error) {
	var files []File
	filesRepo, err := s.repoFiles.GetAllFilesByUserID(userID)
	if err != nil {
		return []File{}, err
	}
	for _, file := range filesRepo {
		files = append(files, File{
			ID:          file.ID,
			Title:       file.Title,
			URL:         s.files.GetObjectLink(file.Title),
			Size:        file.Size,
			ContentType: file.ContentType,
			Type:        file.Type,
			CreatedAt:   file.CreatedAt,
			ForeignID:   file.UserID,
		})
	}
	return files, nil
}

func (s *FilesService) GetAllFilesByType(userID types.BinaryUUID, fileType domain.FileType) ([]File, error) {
	var files []File
	filesRepo, err := s.repoFiles.GetAllTypeFilesByUserID(userID, fileType)
	if err != nil {
		return []File{}, err
	}
	for _, file := range filesRepo {
		files = append(files, File{
			ID:          file.ID,
			Title:       file.Title,
			URL:         s.files.GetObjectLink(file.Title),
			Size:        file.Size,
			ContentType: file.ContentType,
			Type:        file.Type,
			CreatedAt:   file.CreatedAt,
			ForeignID:   file.UserID,
		})
	}
	return files, nil
}

func (s *FilesService) GetFile(userID, fileID types.BinaryUUID) (File, error) {
	var file File
	fileRepo, err := s.repoFiles.GetFileByUserID(userID, fileID)
	if err != nil {
		return File{}, err
	}
	file = File{
		ID:          fileRepo.ID,
		Title:       fileRepo.Title,
		URL:         s.files.GetObjectLink(fileRepo.Title),
		Size:        fileRepo.Size,
		ContentType: fileRepo.ContentType,
		Type:        fileRepo.Type,
		CreatedAt:   fileRepo.CreatedAt,
		ForeignID:   fileRepo.UserID,
	}
	return file, nil
}

func (s *FilesService) UploadFile(ctx context.Context, file File) error {
	title, err := s.files.UploadObject(ctx, users, file)
	err = s.repoFiles.Create(domain.File{
		Title:       title,
		Size:        file.Size,
		Current:     false,
		ContentType: file.ContentType,
		Type:        file.Type,
		CreatedAt:   time.Now(),
		UserID:      file.ForeignID,
	})
	return err
}

func (s *FilesService) DeleteFile(userID, fileID types.BinaryUUID) error {
	err := s.repoFiles.DeleteByID(userID, fileID)
	return err
}

//TODO check fileType
