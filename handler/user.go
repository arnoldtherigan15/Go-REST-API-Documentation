package handler

import (
	"net/http"

	"github.com/arnoldtherigan15/Go-REST-API-Documentation/user"
	"github.com/gin-gonic/gin"
)

type handler struct {
	userService user.Service
}

func NewHandler(userService user.Service) *handler {
	return &handler{userService}
}

func (h *handler) Register(c *gin.Context) {
	var userInput *user.RegisterInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "failed to register",
		})
		return
	}
	createdUser, err := h.userService.Create(*userInput)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": "failed to register",
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
