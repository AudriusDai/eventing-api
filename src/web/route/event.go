package route

import (
	"errors"
	"net/http"
	"time"

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
	VideoQuality []string  `json:"videoQuality" binding:"required,gt=0,dive,oneof=144p 240p 360p 480p 720p 1080p 1440p 2160p"`
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

	m, err := dtoToModelEvent(dto)
	if err != nil {
		ctx.Error(err)
		return
	}

	result, err := core.CreateEvent(m)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, modelToDtoEvent(result))
}

func dtoToModelEvent(d EventDto) (model.Event, error) {
	date, err := time.Parse(time.RFC3339, d.Date)

	if err != nil {
		return model.Event{}, errors.Join(err, errors.New("failed to parse date field"))
	}

	return model.Event{
		Name:         d.Name,
		Date:         date,
		Languages:    d.Languages,
		VideoQuality: d.VideoQuality,
		AudioQuality: d.AudioQuality,
		Invitees:     d.Invitees,
		Description:  d.Description,
	}, nil
}

func modelToDtoEvent(m model.Event) EventDto {
	return EventDto{
		Id:           m.Id,
		Name:         m.Name,
		Date:         m.Date.Format(time.RFC3339),
		Languages:    m.Languages,
		VideoQuality: m.VideoQuality,
		AudioQuality: m.AudioQuality,
		Invitees:     m.Invitees,
		Description:  m.Description,
	}
}
