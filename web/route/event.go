package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventDto struct {
}

func postEvent(app *gin.RouterGroup) {
	app.POST("/event", postEventFunc)
}

func postEventFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"great": "success"})
}
