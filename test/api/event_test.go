package apitest

import (
	"net/http"
	"testing"

	"github.com/audriusdai/eventing-api/core"
	"github.com/audriusdai/eventing-api/core/model"
	"github.com/audriusdai/eventing-api/web/route"
)

func TestPostEvent(t *testing.T) {
	_CreateEvent := core.CreateEvent
	defer func() { core.CreateEvent = _CreateEvent }()

	t.Run("with given event details, return created", func(t *testing.T) {
		core.CreateEvent = func(event model.Event) (model.Event, error) { return model.Event{}, nil }

		res := testClient(t).POST("/api/event").Expect()

		res.Status(http.StatusCreated)
		res.JSON().Equal(route.EventDto{})
	})
}
