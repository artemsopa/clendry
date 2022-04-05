package v1

import (
	"github.com/artomsopun/clendry/clendry-api/pkg/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *echo.Group) {
	users := api.Group("/users", h.checkAuth)
	{
		users.GET("", h.getAllUsers)
		users.GET("/:id", h.getDefProfile)

		friends := users.Group("/friends")
		{
			friends.GET("", h.getAllFriends)

			reqs := friends.Group("/reqs")
			{
				reqs.GET("/in", h.getAllIncomingRequests)
				reqs.GET("/out", h.getAllSentRequests)
				reqs.POST("", h.sendFriendRequest)
				reqs.PUT("", h.confirmFriendRequest)
				reqs.DELETE("", h.deleteFriendRequest)
			}

			blocks := friends.Group("/blocks")
			{
				blocks.GET("", h.getAllBlockedUsers)
				blocks.POST("", h.addToBlock)
				blocks.DELETE("", h.deleteBlock)
			}
		}
	}
}

type userRelates struct {
	ID       types.BinaryUUID `json:"id"`
	Nick     string           `json:"nick"`
	Email    string           `json:"email"`
	Avatar   string           `json:"avatar"`
	BlockTo  bool             `json:"block_to"`
	BlockBy  bool             `json:"block_by"`
	Sent     bool             `json:"sent"`
	Incoming bool             `json:"incoming"`
	Friend   bool             `json:"friend"`
}

func (h *Handler) getAllUsers(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	usersServ, err := h.services.Users.GetAllUsers(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var users []userRelates
	for _, value := range usersServ {
		users = append(users, userRelates{
			ID:       value.ID,
			Nick:     value.Nick,
			Email:    value.Email,
			Avatar:   value.Avatar,
			BlockTo:  value.BlockTo,
			BlockBy:  value.BlockBy,
			Sent:     value.Sent,
			Incoming: value.Incoming,
			Friend:   value.Friend,
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) getDefProfile(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	defIDStr := c.Param("id")

	defID := types.ParseUUID(defIDStr)
	userServ, err := h.services.Users.GetUserByID(userID, defID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	user := userRelates{
		ID:       userServ.ID,
		Nick:     userServ.Nick,
		Email:    userServ.Email,
		Avatar:   userServ.Avatar,
		BlockTo:  userServ.BlockTo,
		BlockBy:  userServ.BlockBy,
		Sent:     userServ.Sent,
		Incoming: userServ.Incoming,
		Friend:   userServ.Friend,
	}
	return c.JSON(http.StatusOK, user)
}

type userInfos struct {
	ID     types.BinaryUUID `json:"id"`
	Nick   string           `json:"nick"`
	Email  string           `json:"email"`
	Avatar string           `json:"avatar"`
}

func (h *Handler) getAllFriends(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	usersServ, err := h.services.Users.GetAllFriends(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var users []userInfos
	for _, value := range usersServ {
		users = append(users, userInfos{
			ID:     value.ID,
			Nick:   value.Nick,
			Email:  value.Email,
			Avatar: value.Avatar,
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) getAllIncomingRequests(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	usersServ, err := h.services.Users.GetAllIncomingRequests(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var users []userInfos
	for _, value := range usersServ {
		users = append(users, userInfos{
			ID:     value.ID,
			Nick:   value.Nick,
			Email:  value.Email,
			Avatar: value.Avatar,
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) getAllSentRequests(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	usersServ, err := h.services.Users.GetAllFriends(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var users []userInfos
	for _, value := range usersServ {
		users = append(users, userInfos{
			ID:     value.ID,
			Nick:   value.Nick,
			Email:  value.Email,
			Avatar: value.Avatar,
		})
	}
	return c.JSON(http.StatusOK, users)
}

type defBody struct {
	DefID types.BinaryUUID `json:"def_id"`
}

func (h *Handler) sendFriendRequest(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var def defBody
	if err := c.Bind(&def); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}

	err = h.services.Users.SendFriendRequest(userID, def.DefID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusOK, "friend request created")
}

func (h *Handler) confirmFriendRequest(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var def defBody
	if err := c.Bind(&def); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}

	err = h.services.Users.ConfirmFriendRequest(userID, def.DefID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusOK, "friend request updated")
}

func (h *Handler) deleteFriendRequest(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	defIDStr := c.Param("id")
	defID := types.ParseUUID(defIDStr)

	err = h.services.Users.DeleteRequest(userID, defID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusOK, "friend request deleted")
}

func (h *Handler) getAllBlockedUsers(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	usersServ, err := h.services.Users.GetAllBlockedUsers(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var users []userInfos
	for _, value := range usersServ {
		users = append(users, userInfos{
			ID:     value.ID,
			Nick:   value.Nick,
			Email:  value.Email,
			Avatar: value.Avatar,
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) addToBlock(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var def defBody
	if err := c.Bind(&def); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}

	err = h.services.Users.CreateBlock(userID, def.DefID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusOK, "friend request updated")
}

func (h *Handler) deleteBlock(c echo.Context) error {
	userID, err := getUserId(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	defIDStr := c.Param("id")
	defID := types.ParseUUID(defIDStr)

	err = h.services.Users.DeleteRequest(userID, defID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusOK, "friend request deleted")
}
