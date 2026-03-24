package handlers

import (
	"go-users-manager-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func (u *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := u.service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}
