package service

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/artomsopun/clendry/clendry-api/pkg/logger"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
)

type FilesService struct {
	repoFiles       repository.Files
	repoFolderFiles repository.FolderFiles
	files           files.Files
}

func NewFilesService(repoFiles repository.Files, repoFolderFiles repository.FolderFiles, files files.Files) *FilesService {
	return &FilesService{
		repoFiles:       repoFiles,
		repoFolderFiles: repoFolderFiles,
		files:           files,
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
			Url:         s.files.GetObjectLink(file.UrlTitle),
			Size:        file.Size,
			ContentType: file.ContentType,
			Type:        file.Type,
			IsFavourite: file.IsFavourite,
			IsTrash:     file.IsTrash,
			CreatedAt:   file.CreatedAt,
			UserID:      file.UserID,
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
			Url:         s.files.GetObjectLink(file.UrlTitle),
			Size:        file.Size,
			ContentType: file.ContentType,
			Type:        file.Type,
			IsFavourite: file.IsFavourite,
			IsTrash:     file.IsTrash,
			CreatedAt:   file.CreatedAt,
			UserID:      file.UserID,
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
		Url:         s.files.GetObjectLink(fileRepo.UrlTitle),
		Size:        fileRepo.Size,
		ContentType: fileRepo.ContentType,
		Type:        fileRepo.Type,
		CreatedAt:   fileRepo.CreatedAt,
		IsFavourite: fileRepo.IsFavourite,
		IsTrash:     fileRepo.IsTrash,
		UserID:      fileRepo.UserID,
	}
	return file, nil
}

func (s *FilesService) UploadFile(ctx context.Context, userID string, file File) (types.BinaryUUID, error) {
	title, url, err := s.files.UploadObject(ctx, userID, USERS, files.File{
		Title:       file.Title,
		Size:        file.Size,
		ContentType: file.ContentType,
		Type:        file.Type,
	})
	if err != nil {
		return types.BinaryUUID{}, err
	}
	id, err := s.repoFiles.Create(domain.File{
		Title:       title,
		UrlTitle:    url,
		Size:        file.Size,
		ContentType: file.ContentType,
		Type:        file.Type,
		IsFavourite: false,
		IsTrash:     false,
		CreatedAt:   time.Now(),
		UserID:      file.UserID,
	})
	if err != nil {
		return types.BinaryUUID{}, err
	}
	return id, nil
}

func (s *FilesService) ChangeFileTitle(userID, fileID types.BinaryUUID, title string) error {
	err := s.repoFiles.ChangeFileTitle(userID, fileID, title)
	return err
}

func (s *FilesService) DeleteFile(userID, fileID types.BinaryUUID) error {
	err := s.repoFiles.DeleteByID(userID, fileID)
	return err
}

func (s *FilesService) GetContentType(ctype string) domain.FileType {
	parts := strings.Split(ctype, "/")
	if parts[0] == "image" {
		return domain.Image
	}
	if parts[0] == "video" {
		return domain.Video
	}
	return domain.Other
}

func (s *FilesService) RemoveFile(filename string) {
	if err := os.Remove(filename); err != nil {
		logger.Error("removeFile(): ", err)
	}
}

func (s *FilesService) GetAllFilesByFolderID(userID, folderID types.BinaryUUID) ([]File, error) {
	members, err := s.repoFolderFiles.GetAllFilesByFolderID(folderID)
	if err != nil {
		return []File{}, err
	}
	var files []File
	for _, member := range members {
		fileRepo, err := s.repoFiles.GetFileByFolder(userID, member.FileID)
		if err != nil {
			continue
		}
		files = append(files, File{
			ID:          fileRepo.ID,
			Title:       fileRepo.Title,
			Url:         s.files.GetObjectLink(fileRepo.UrlTitle),
			Size:        fileRepo.Size,
			ContentType: fileRepo.ContentType,
			Type:        fileRepo.Type,
			IsFavourite: fileRepo.IsFavourite,
			IsTrash:     fileRepo.IsTrash,
			CreatedAt:   fileRepo.CreatedAt,
			MemberID:    member.ID,
		})
	}
	return files, nil
}

func (s *FilesService) AddFileToFolder(userID, folderID, fileID types.BinaryUUID) error {
	//check user ID
	err := s.repoFolderFiles.CreateFolderFile(domain.FolderFile{
		FolderID: folderID,
		FileID:   fileID,
	})
	return err
}

func (s *FilesService) DeleteFileFromFolder(memberID types.BinaryUUID) error {
	//check user ID
	err := s.repoFolderFiles.DeleteByFolderFileID(memberID)
	return err
}

func (s *FilesService) GetAllFavouriteByUserID(userID types.BinaryUUID) ([]File, error) {
	var files []File
	filesRepo, err := s.repoFiles.GetAllFavouriteByUserID(userID)
	if err != nil {
		return []File{}, err
	}
	for _, file := range filesRepo {
		files = append(files, File{
			ID:          file.ID,
			Title:       file.Title,
			Url:         s.files.GetObjectLink(file.UrlTitle),
			Size:        file.Size,
			ContentType: file.ContentType,
			Type:        file.Type,
			IsFavourite: file.IsFavourite,
			IsTrash:     file.IsTrash,
			CreatedAt:   file.CreatedAt,
			UserID:      file.UserID,
		})
	}
	return files, nil
}

func (s *FilesService) AddToFavourite(userID, fileID types.BinaryUUID) error {
	err := s.repoFiles.AddToFavourite(userID, fileID)
	return err
}
func (s *FilesService) DeleteFromFavourite(userID, fileID types.BinaryUUID) error {
	err := s.repoFiles.DeleteFromFavourite(userID, fileID)
	return err
}

func (s *FilesService) GetAllTrashByUserID(userID types.BinaryUUID) ([]File, error) {
	var files []File
	filesRepo, err := s.repoFiles.GetAllTrashByUserID(userID)
	if err != nil {
		return []File{}, err
	}
	for _, file := range filesRepo {
		files = append(files, File{
			ID:          file.ID,
			Title:       file.Title,
			Url:         s.files.GetObjectLink(file.UrlTitle),
			Size:        file.Size,
			ContentType: file.ContentType,
			Type:        file.Type,
			IsFavourite: file.IsFavourite,
			IsTrash:     file.IsTrash,
			CreatedAt:   file.CreatedAt,
			UserID:      file.UserID,
		})
	}
	return files, nil
}

func (s *FilesService) AddToTrash(userID, fileID types.BinaryUUID) error {
	err := s.repoFiles.AddToTrash(userID, fileID)
	return err
}

func (s *FilesService) DeleteFromTrash(userID, fileID types.BinaryUUID) error {
	err := s.repoFiles.DeleteFromTrash(userID, fileID)
	return err
}
