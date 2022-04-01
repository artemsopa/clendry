package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) checkAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := h.getAccessCookie(c)
		if err != nil {
			log.Println(err.Error())
			return newResponse(c, http.StatusUnauthorized, err.Error())
		}

		c.Set(userCtx, id)
		return next(c)
	}
}

func (h *Handler) getAccessCookie(c echo.Context) (string, error) {
	accessCookie, err := c.Cookie(AccessToken)
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return "", errors.New("you don't have any cookie")
		}
		return "", err
	}
	return h.tokenManager.Parse(accessCookie.Value)
}

func (h *Handler) parseAuthHeader(c echo.Context) (string, error) {
	header := c.Request().Header.Get(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}

func getUserId(c echo.Context) (uint, error) {
	return getIdByContext(c, userCtx)
}

func getIdByContext(c echo.Context, context string) (uint, error) {
	idFromCtx := c.Get(context)

	idStr, ok := idFromCtx.(string)
	if !ok {
		return 0, errors.New("userCtx is of invalid type")
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("userCtx is of invalid type")
	}

	id := uint(idInt)

	return id, nil
}
