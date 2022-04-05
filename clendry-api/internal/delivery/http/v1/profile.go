package v1

import (
	"bytes"
	"fmt"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/internal/service"
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func (h *Handler) initProfilesRoutes(api *echo.Group) {
	profile := api.Group("/profile", h.checkAuth)
	{
		profile.GET("", h.getProfile)
		profile.PUT("/password", h.changePassword)
		profile.PUT("/avatar", h.changeAvatar)
		profile.DELETE("", h.deleteProfile)
	}
}

type userInfo struct {
	ID     types.BinaryUUID `json:"id"`
	Nick   string           `json:"nick"`
	Email  string           `json:"email"`
	Avatar string           `json:"avatar"`
}

func (h *Handler) getProfile(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	user, err := h.services.Profiles.GetProfile(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userInfo{
		ID:     user.ID,
		Nick:   user.Nick,
		Email:  user.Email,
		Avatar: user.Avatar,
	})
}

type changePassReq struct {
	OldPassword string `json:"old_password"`
	Password    string `json:"password" binding:"required,min=8,max=64"`
	Confirm     string `json:"confirm" binding:"required,min=8,max=64"`
}

func (h *Handler) changePassword(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	var pass changePassReq
	if err := c.Bind(&pass); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.Profiles.ChangePassword(service.PasswordConfirm{
		UserID:      userID,
		OldPassword: pass.OldPassword,
		Passwords: service.Passwords{
			Password: pass.Password,
			Confirm:  pass.Confirm,
		},
	}); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}
	return newResponse(c, http.StatusOK, "password changed")
}

const maxUploadSize = 5 << 20

var imageTypes = map[string]interface{}{
	"image/jpeg": nil,
	"image/jpg":  nil,
	"image/png":  nil,
}

func (h *Handler) changeAvatar(c echo.Context) error {
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
	defer h.filesManager.RemoveFile(tempFilename)

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	if _, err := file.Read(buffer); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	contentType := http.DetectContentType(buffer)
	// Validate File Type
	if _, ex := imageTypes[contentType]; !ex {
		return newResponse(c, http.StatusBadRequest, "file type is not supported")
	}

	f, err := os.OpenFile(tempFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, "failed to create temp file")
	}

	defer f.Close()

	if _, err := io.Copy(f, bytes.NewReader(buffer)); err != nil {
		return newResponse(c, http.StatusInternalServerError, "failed to write chunk to temp file")
	}

	err = h.services.Profiles.UploadAvatar(c.Request().Context(), service.File{
		Title:       tempFilename,
		Size:        fileHeader.Size,
		ContentType: contentType,
		Type:        domain.Image,
		ForeignID:   userID,
	})
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "avatar changed")
}

func (h *Handler) deleteProfile(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	err = h.services.Profiles.DeleteProfile(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "profile deleted")
}
