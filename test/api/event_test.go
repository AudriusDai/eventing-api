package apitest

import (
	"net/http"
	"testing"

	"github.com/audriusdai/eventing-api/core"
	"github.com/audriusdai/eventing-api/core/model"
	"github.com/audriusdai/eventing-api/util"
	"github.com/audriusdai/eventing-api/web/route"
)

func TestPostEvent(t *testing.T) {
	_CreateEvent := core.CreateEvent
	defer func() { core.CreateEvent = _CreateEvent }()
	validEventDto := route.EventDto{
		Name:         "The name",
		Date:         "2023-04-20T14:00:00Z",
		Languages:    []string{"English", "French"},
		VideoQuality: []string{"720p", "1080p"},
		AudioQuality: []string{"High", "Low"},
		Invitees:     []string{"anthony@email.com"},
		Description:  util.ToRef("This is an event."),
	}

	t.Run("with given event details, return created", func(t *testing.T) {
		core.CreateEvent = func(event model.Event) (model.Event, error) { return model.Event{}, nil }
		dto := validEventDto

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusCreated)
	})
}
