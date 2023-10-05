package route

import (
	"net/http"

	"github.com/audriusdai/eventing-api/core"
	"github.com/audriusdai/eventing-api/core/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventDto struct {
	Id           uuid.UUID `json:"id,omitempty"`
	Name         string    `json:"name" binding:"required,max=256"`
	Date         string    `json:"date" binding:"required,datetimeRFC3339"`
	Languages    []string  `json:"languages" binding:"required,gt=0,dive,languageISO6391"`
	VideoQuality []string  `json:"videoQuality" binding:"required,gt=0,dive,oneof=720p 1080p"`
	AudioQuality []string  `json:"audioQuality" binding:"required,gt=0,dive,oneof=Low Medium High"`
	Invitees     []string  `json:"invitees" binding:"required,gt=0,lte=100,dive,email"`
	Description  string    `json:"description,omitempty" binding:"max=1024"`
}

func postEvent(app *gin.RouterGroup) {
	app.POST("/event", postEventFunc)
}

func postEventFunc(ctx *gin.Context) {
	dto := EventDto{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.Error(err)
		return
	}

	result, err := core.CreateEvent(dtoToModelEvent(dto))

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, modelToDtoEvent(result))
}

func dtoToModelEvent(d EventDto) model.Event {
	// todo: map properly
	return model.Event{
		Name: d.Name,
	}
}

func modelToDtoEvent(m model.Event) EventDto {
	// todo: map properly
	return EventDto{
		Id:   m.Id,
		Name: m.Name,
	}
}
