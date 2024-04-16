package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/service/auth"
	"github.com/onainadapdap1/golang_kantin/internal/service/user"
)

type UserHandler interface {
	Login(c *gin.Context)
}

type userHandler struct {
	userService user.UserService
	authService auth.AuthService
}

func NewUserHandler(userService user.UserService, authService auth.AuthService) UserHandler {
	return &userHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *userHandler) Login(c *gin.Context) {
	var input api.LoginInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unproccessable entity bro"})
		return
	}

	loggedUser, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "login failed!"})
		return
	}

	token, err := h.authService.GenerateToken(int(loggedUser.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token!"})
		return
	}

	c.SetCookie("access_token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token": token,
		"message": "successfully logged in",
		"user": loggedUser,
	})
}
