package apitest

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/audriusdai/eventing-api/core"
	"github.com/audriusdai/eventing-api/core/model"
	"github.com/audriusdai/eventing-api/web/middleware"
	"github.com/audriusdai/eventing-api/web/route"
	"github.com/google/uuid"
)

func TestPostEvent(t *testing.T) {
	_CreateEvent := core.CreateEvent
	defer func() { core.CreateEvent = _CreateEvent }()
	core.CreateEvent = func(event model.Event) (model.Event, error) {
		event.Id = uuid.New()
		return event, nil
	}
	validEventDto := route.EventDto{
		Name:         "The name",
		Date:         "2023-04-20T14:00:00Z",
		Languages:    []string{"English", "French"},
		VideoQuality: []string{"720p", "1080p"},
		AudioQuality: []string{"High", "Low"},
		Invitees:     []string{"anthony@email.com"},
	}

	t.Run("without a name, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Name = ""

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().Contains("`Name` is required")
	})

	t.Run("with incorrect date value, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Date = "this.is.not.a.date.in.here"

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`this.is.not.a.date.in.here` is invalid date (RFC3339)")
	})

	t.Run("without any language, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Languages = []string{}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`Languages` should be greater than 0")
	})

	t.Run("with unrecognized language, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Languages = []string{"not.a.language.1", "English", "not.a.language.2"}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`not.a.language.1` is invalid language (ISO6391)")
		res.JSON().Object().Value("error_message").String().
			Contains("`not.a.language.2` is invalid language (ISO6391)")
	})

	t.Run("without any video quality, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.VideoQuality = []string{}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`VideoQuality` should be greater than 0")
	})

	t.Run("with unrecognized video quality, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.VideoQuality = []string{"not.a.video.quality.1", "720p", "not.a.video.quality.2"}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`not.a.video.quality.1` is invalid for field `VideoQuality[0]`")
		res.JSON().Object().Value("error_message").String().
			Contains("`not.a.video.quality.2` is invalid for field `VideoQuality[2]`")
	})

	t.Run("without any audio quality, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.AudioQuality = []string{}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`AudioQuality` should be greater than 0")
	})

	t.Run("with unrecognized audio quality, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.AudioQuality = []string{"not.an.audio.quality", "Low"}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`not.an.audio.quality` is invalid for field `AudioQuality[0]`")
	})

	t.Run("without any invitee, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Invitees = []string{}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`Invitees` should be greater than 0")
	})

	t.Run("with an invalid invitee, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Invitees = []string{"this.is.not.an.email", "anthony@email.com"}

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`this.is.not.an.email` is an invalid email")
	})

	t.Run("with too many invitees, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Invitees = getInvitees(101)

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`Invitees` should be less/equal than 100")
	})

	t.Run("with too long description, return bad request", func(t *testing.T) {
		dto := validEventDto
		dto.Description = strings.Join(make([]string, 1026), "a")

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusBadRequest).JSON().Schema(middleware.WebResponseError{})
		res.JSON().Object().Value("error_message").String().
			Contains("`Description` max size 1024.")
	})

	t.Run("with valid details, return created", func(t *testing.T) {
		dto := validEventDto

		res := testClient(t).
			POST("/api/event").
			WithJSON(dto).
			Expect()

		res.Status(http.StatusCreated)
	})
}

func getInvitees(size int) []string {
	emails := make([]string, 0, size)
	for i := 0; i < size; i++ {
		emails = append(emails, fmt.Sprintf("user.%v@email.com", i))
	}
	return emails
}
