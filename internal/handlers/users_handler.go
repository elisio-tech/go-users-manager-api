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

func (u *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	users, err := u.service.GetUserByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"massage": err})
		return
	}
	ctx.JSON(200, users)
}

func (u *UserHandler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}

	err := u.service.DeleteUser(id)
	if err != nil {

	}
	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
