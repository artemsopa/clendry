package v1

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/service"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/labstack/echo/v4"
)

func (h *Handler) initStoragesRoutes(api *echo.Group) {
	storage := api.Group("/storage", h.checkAuth)
	{
		storage.GET("", h.getMembership)
		storage.GET("/kb", h.getSumKB)
		files := storage.Group("/files")
		{
			files.GET("/", h.getAllFiles)
			files.GET("/:type", h.getFilesByType)
			files.GET("/folders/:id", h.getAllFolderFiles)
			files.POST("/upload", h.uploadFile)
			files.PUT("/folder", h.pushToFolder)
			files.PUT("/title", h.changeFileTitle)
			files.DELETE("", h.deleteFile)

			fav := files.Group("/fav")
			{
				fav.GET("/", h.getAllFavourite)
				fav.PUT("/", h.addToFavourite)
				fav.PUT("/remove", h.deleteFromFavourite)
			}

			trash := files.Group("/trash")
			{
				trash.GET("/", h.getAllTrash)
				trash.PUT("/", h.addToTrash)
				trash.PUT("/remove", h.deleteFromTrash)
			}
		}
		folders := storage.Group("/folders")
		{
			folders.GET("/", h.getAllFolders)
			folders.GET("/:id", h.getAllFilesByFolder)
			folders.POST("/", h.createFoler)
			folders.PUT("/", h.updateFolder)
			folders.DELETE("/:id", h.deleteFolder)
			folders.DELETE("/member/:id", h.deleteFromFolder)
		}
	}
}

func (h *Handler) getSumKB(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, h.services.Storages.GetFilesKBSum(userID))
}

func (h *Handler) uploadFile(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.Request().Body = http.MaxBytesReader(c.Response().Writer, c.Request().Body, maxUploadSize)

	file, fileHeader, err := c.Request().FormFile("file")
	if err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	tempFileName := fileHeader.Filename

	defer h.services.Storages.RemoveFile(tempFileName)

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	if _, err := file.Read(buffer); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	contentType := http.DetectContentType(buffer)
	// Validate File Type
	// if _, ex := imageTypes[contentType]; !ex {
	// 	return newResponse(c, http.StatusBadRequest, "file type is not supported")
	// }

	ctype := h.services.Storages.GetContentType(contentType)

	f, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, "failed to create temp file")
	}

	defer f.Close()

	if _, err := io.Copy(f, bytes.NewReader(buffer)); err != nil {
		return newResponse(c, http.StatusInternalServerError, "failed to write chunk to temp file")
	}

	id, err := h.services.Storages.UploadFile(c.Request().Context(), userID.String(), service.File{
		Title:       tempFileName,
		Size:        fileHeader.Size,
		ContentType: contentType,
		Type:        ctype,
		UserID:      userID,
	})
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, id)
}

type inputTitle struct {
	ID    types.BinaryUUID `json:"id"`
	Title string           `json:"title"`
}

func (h *Handler) changeFileTitle(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var file inputTitle
	if err := c.Bind(&file); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Storages.ChangeFileTitle(userID, file.ID, file.Title)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file title changed")
}

func (h *Handler) deleteFile(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	fileIDStr := c.QueryParam("id")

	fileID := types.ParseUUID(fileIDStr)
	err = h.services.Storages.DeleteFile(userID, fileID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file deleted")
}

type folderInput struct {
	Title string `json:"title"`
}

func (h *Handler) createFoler(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var folder folderInput
	if err := c.Bind(&folder); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Folders.CreateFolder(service.Folder{
		Title:  folder.Title,
		UserID: userID,
	})
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "folder created")
}

type folderUpdateInput struct {
	ID    types.BinaryUUID `json:"id"`
	Title string           `json:"title"`
}

func (h *Handler) updateFolder(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var folder folderUpdateInput
	if err := c.Bind(&folder); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Folders.ChangeFolderTitleUserID(userID, folder.ID, folder.Title)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "folder updated")
}

func (h *Handler) deleteFolder(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	folderIDStr := c.Param("id")

	folderID := types.ParseUUID(folderIDStr)
	err = h.services.Folders.DeleteFolderByID(userID, folderID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "folder deleted")
}

type folder struct {
	ID        types.BinaryUUID `json:"id"`
	Title     string           `json:"title"`
	CreatedAt time.Time        `json:"created_at"`
	UserID    types.BinaryUUID `json:"user_id"`
}

func (h *Handler) getAllFolders(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	folderServ, err := h.services.Folders.GetAllFoldersByUserID(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var folders []folder
	for _, value := range folderServ {
		folders = append(folders, folder{
			ID:        value.ID,
			Title:     value.Title,
			CreatedAt: value.CreatedAt,
			UserID:    value.UserID,
		})
	}
	return c.JSON(http.StatusOK, folders)
}

func (h *Handler) getAllFolderFiles(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	fileIDStr := c.Param("id")

	fileID := types.ParseUUID(fileIDStr)
	folderServ, err := h.services.Storages.GetAllFolderFilessByFileID(userID, fileID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var folders [][]folder
	var typedFolder []folder
	for _, folderServArr := range folderServ {
		for _, foldeServ := range folderServArr {
			typedFolder = append(typedFolder, folder{
				ID:        foldeServ.ID,
				Title:     foldeServ.Title,
				CreatedAt: foldeServ.CreatedAt,
				UserID:    foldeServ.UserID,
			})
		}
		folders = append(folders, typedFolder)
		typedFolder = []folder{}
	}
	return c.JSON(http.StatusOK, folders)
}

type input struct {
	FolderID types.BinaryUUID `json:"folder_id"`
	FileID   types.BinaryUUID `json:"file_id"`
}

type idInput struct {
	ID types.BinaryUUID `json:"id"`
}

type file struct {
	ID types.BinaryUUID `json:"id"`

	Title       string          `json:"title"`
	Url         string          `json:"url"`
	Size        int64           `json:"size"`
	ContentType string          `json:"c_type"`
	Type        domain.FileType `json:"type"`
	IsFavourite bool            `json:"is_fav"`
	IsTrash     bool            `json:"is_trash"`
	CreatedAt   time.Time       `json:"created_at"`
}

type fileFolder struct {
	ID types.BinaryUUID `json:"id"`

	Title       string          `json:"title"`
	Url         string          `json:"url"`
	Size        int64           `json:"size"`
	ContentType string          `json:"c_type"`
	Type        domain.FileType `json:"type"`
	IsFavourite bool            `json:"is_fav"`
	IsTrash     bool            `json:"is_trash"`
	CreatedAt   time.Time       `json:"created_at"`

	MemberID types.BinaryUUID `json:"member_id"`
}

func (h *Handler) getAllFilesByFolder(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	folderIDStr := c.Param("id")

	folderID := types.ParseUUID(folderIDStr)
	filesServ, err := h.services.Storages.GetAllFilesByFolderID(userID, folderID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var files []fileFolder
	for _, value := range filesServ {
		files = append(files, fileFolder{
			ID:          value.ID,
			Title:       value.Title,
			Url:         value.Url,
			Size:        value.Size,
			ContentType: value.ContentType,
			Type:        value.Type,
			IsFavourite: value.IsFavourite,
			IsTrash:     value.IsTrash,
			CreatedAt:   value.CreatedAt,
			MemberID:    value.MemberID,
		})
	}
	return c.JSON(http.StatusOK, files)
}

func (h *Handler) pushToFolder(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var folder input
	if err := c.Bind(&folder); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Storages.AddFileToFolder(userID, folder.FolderID, folder.FileID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file added to folder")
}

func (h *Handler) getFilesByType(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	typeStr := c.Param("type")
	filesServ, err := h.services.Storages.GetAllFilesByType(userID, domain.FileType(typeStr))
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var files []file
	for _, value := range filesServ {
		files = append(files, file{
			ID:          value.ID,
			Title:       value.Title,
			Url:         value.Url,
			Size:        value.Size,
			ContentType: value.ContentType,
			Type:        value.Type,
			IsFavourite: value.IsFavourite,
			IsTrash:     value.IsTrash,
			CreatedAt:   value.CreatedAt,
		})
	}
	return c.JSON(http.StatusOK, files)
}

func (h *Handler) getAllFiles(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	filesServ, err := h.services.Storages.GetAllFiles(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var files []file
	for _, value := range filesServ {
		files = append(files, file{
			ID:          value.ID,
			Title:       value.Title,
			Url:         value.Url,
			Size:        value.Size,
			ContentType: value.ContentType,
			Type:        value.Type,
			IsFavourite: value.IsFavourite,
			IsTrash:     value.IsTrash,
			CreatedAt:   value.CreatedAt,
		})
	}
	return c.JSON(http.StatusOK, files)
}

//Fav

func (h *Handler) getAllFavourite(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	filesServ, err := h.services.Storages.GetAllFavouriteByUserID(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var files []file
	for _, value := range filesServ {
		files = append(files, file{
			ID:          value.ID,
			Title:       value.Title,
			Url:         value.Url,
			Size:        value.Size,
			ContentType: value.ContentType,
			Type:        value.Type,
			IsFavourite: value.IsFavourite,
			IsTrash:     value.IsTrash,
			CreatedAt:   value.CreatedAt,
		})
	}
	return c.JSON(http.StatusOK, files)
}

func (h *Handler) addToFavourite(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var file idInput
	if err := c.Bind(&file); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Storages.AddToFavourite(userID, file.ID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file added to favourite")
}

func (h *Handler) deleteFromFavourite(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var file idInput
	if err := c.Bind(&file); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Storages.DeleteFromFavourite(userID, file.ID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file deleted from favourite")
}

//Trash

func (h *Handler) getAllTrash(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	filesServ, err := h.services.Storages.GetAllTrashByUserID(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var files []file
	for _, value := range filesServ {
		files = append(files, file{
			ID:          value.ID,
			Title:       value.Title,
			Url:         value.Url,
			Size:        value.Size,
			ContentType: value.ContentType,
			Type:        value.Type,
			IsFavourite: value.IsFavourite,
			IsTrash:     value.IsTrash,
			CreatedAt:   value.CreatedAt,
		})
	}
	return c.JSON(http.StatusOK, files)
}

func (h *Handler) addToTrash(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var file idInput
	if err := c.Bind(&file); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Storages.AddToTrash(userID, file.ID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file added to trash")
}

func (h *Handler) deleteFromTrash(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var file idInput
	if err := c.Bind(&file); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}
	err = h.services.Storages.DeleteFromTrash(userID, file.ID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file deleted")
}

func (h *Handler) deleteFromFolder(c echo.Context) error {
	_, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	memberIDStr := c.Param("id")

	memberID := types.ParseUUID(memberIDStr)
	err = h.services.Storages.DeleteFileFromFolder(memberID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "member deleted")
}

type member struct {
	ID types.BinaryUUID `json:"id"`

	FolderID types.BinaryUUID `json:"folder_id"`
	FileID   types.BinaryUUID `json:"file_id"`
}

func (h *Handler) getMembership(c echo.Context) error {
	folderID := c.QueryParam("folder")
	fileID := c.QueryParam("file")
	memberServ, err := h.services.Storages.GetByFolderFileID(types.ParseUUID(folderID), types.ParseUUID(fileID))
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, member{
		ID:       memberServ.ID,
		FolderID: memberServ.FolderID,
		FileID:   memberServ.FileID,
	})
}
