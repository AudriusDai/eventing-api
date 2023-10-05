package route

import (
	"net/http"

	"github.com/audriusdai/eventing-api/core"
	"github.com/audriusdai/eventing-api/core/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventDto struct {
	Id           uuid.UUID `json:"id" binding:"uuid"`
	Name         string    `json:"name"`
	Date         string    `json:"date"`      // should be utc timestampz
	Languages    []string  `json:"languages"` // check if there is ISO validator for languages
	VideoQuality []string  `json:"videoQuality"`
	AudioQuality []string  `json:"audioQuality"`
	Invitees     []string  `json:"invitees"`    // should have email validation
	Description  string    `json:"description"` // length optional
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
