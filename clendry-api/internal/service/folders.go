package service

import (
	"time"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
)

type FoldersService struct {
	repoFolders repository.Folders
}

func NewFoldersService(repoFolders repository.Folders) *FoldersService {
	return &FoldersService{
		repoFolders: repoFolders,
	}
}

func (s *FoldersService) GetAllFoldersByUserID(userID types.BinaryUUID) ([]Folder, error) {
	var folders []Folder
	foldersRepo, err := s.repoFolders.GetAllFoldersByUserID(userID)
	if err != nil {
		return []Folder{}, err
	}
	for _, folder := range foldersRepo {
		folders = append(folders, Folder{
			ID:        folder.ID,
			Title:     folder.Title,
			CreatedAt: folder.CreatedAt,
			UserID:    folder.UserID,
		})
	}
	return folders, nil
}

func (s *FoldersService) GetFolderByUserID(userID, folderID types.BinaryUUID) (Folder, error) {
	var folder Folder
	folderRepo, err := s.repoFolders.GetFolderByUserID(userID, folderID)
	if err != nil {
		return Folder{}, err
	}
	folder = Folder{
		ID:        folderRepo.ID,
		Title:     folderRepo.Title,
		CreatedAt: folderRepo.CreatedAt,
		UserID:    folderRepo.UserID,
	}
	return folder, nil
}

func (s *FoldersService) ChangeFolderTitleUserID(userID, folderID types.BinaryUUID, title string) error {
	err := s.repoFolders.ChangeFolderTitleUserID(userID, folderID, title)
	return err
}

func (s *FoldersService) CreateFolder(folder Folder) error {
	err := s.repoFolders.Create(domain.Folder{
		Title:     folder.Title,
		CreatedAt: time.Now(),
		UserID:    folder.UserID,
	})
	return err
}

func (s *FoldersService) DeleteFolderByID(userID, folderID types.BinaryUUID) error {
	err := s.repoFolders.DeleteByID(userID, folderID)
	return err
}
