package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Is working!")
	})
	r.Run(":3000")
}
