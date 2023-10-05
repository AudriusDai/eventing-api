package route

import (
	"testing"
	"time"

	"github.com/audriusdai/eventing-api/core/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDtoToModelEvent(t *testing.T) {
	t.Run("with proper details, return mapped", func(t *testing.T) {
		input := EventDto{
			Name:         "The name",
			Date:         "2023-04-20T14:00:00Z",
			Languages:    []string{"English", "French"},
			VideoQuality: []string{"720p", "1080p"},
			AudioQuality: []string{"High", "Low"},
			Invitees:     []string{"anthony@email.com"},
		}

		actual, err := dtoToModelEvent(input)

		assert.Nil(t, err)
		assert.Equal(t, input.Name, actual.Name)
		assert.Equal(t, input.Date, actual.Date.Format(time.RFC3339))
		assert.Equal(t, input.Languages, actual.Languages)
		assert.Equal(t, input.VideoQuality, actual.VideoQuality)
		assert.Equal(t, input.AudioQuality, actual.AudioQuality)
		assert.Equal(t, input.Invitees, actual.Invitees)
	})

	t.Run("with invalid date, return error", func(t *testing.T) {
		input := EventDto{
			Date: "2023-04-20 14:00:00",
		}

		_, err := dtoToModelEvent(input)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "failed to parse date field")
	})
}

func TestModelToDtoEvent(t *testing.T) {
	t.Run("with proper details, return mapped", func(t *testing.T) {
		date, _ := time.Parse(time.RFC3339, "2023-04-20T14:00:00Z")
		input := model.Event{
			Id:           uuid.New(),
			Name:         "The name",
			Date:         date,
			Languages:    []string{"English", "French"},
			VideoQuality: []string{"720p", "1080p"},
			AudioQuality: []string{"High", "Low"},
			Invitees:     []string{"anthony@email.com"},
		}

		actual := modelToDtoEvent(input)

		assert.Equal(t, input.Id, actual.Id)
		assert.Equal(t, input.Name, actual.Name)
		assert.Equal(t, input.Date.Format(time.RFC3339), actual.Date)
		assert.Equal(t, input.Languages, actual.Languages)
		assert.Equal(t, input.VideoQuality, actual.VideoQuality)
		assert.Equal(t, input.AudioQuality, actual.AudioQuality)
		assert.Equal(t, input.Invitees, actual.Invitees)
	})
}
