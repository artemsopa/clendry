package v1

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/service"
	"github.com/artomsopun/clendry/clendry-api/pkg/logger"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/labstack/echo/v4"
)

func (h *Handler) initStoragesRoutes(api *echo.Group) {
	profile := api.Group("/storage", h.checkAuth)
	{
		profile.PUT("/upload", h.uploadFile)
		profile.DELETE(":id", h.deleteFile)
	}
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

	tempFilename := fmt.Sprintf("%d-%s", userID, fileHeader.Filename)
	defer h.RemoveFile(tempFilename)

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

	f, err := os.OpenFile(tempFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, "failed to create temp file")
	}

	defer f.Close()

	if _, err := io.Copy(f, bytes.NewReader(buffer)); err != nil {
		return newResponse(c, http.StatusInternalServerError, "failed to write chunk to temp file")
	}

	err = h.services.Storages.UploadFile(c.Request().Context(), service.File{
		Title:       tempFilename,
		Size:        fileHeader.Size,
		ContentType: contentType,
		Type:        domain.Image,
		ForeignID:   userID,
	})
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "file uploaded")
}

func (h *Handler) deleteFile(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	fileIDStr := c.Param("id")

	fileID := types.ParseUUID(fileIDStr)
	err = h.services.Storages.DeleteFile(userID, fileID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "profile deleted")
}

func (s *Handler) RemoveFile(filename string) {
	if err := os.Remove(filename); err != nil {
		logger.Error("removeFile(): ", err)
	}
}
