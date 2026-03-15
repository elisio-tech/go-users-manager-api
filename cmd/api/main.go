package main

import (
	"go-users-manager-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
	r.Run(":3000")
}
