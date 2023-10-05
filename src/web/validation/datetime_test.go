package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateDatetimeRFC3339(t *testing.T) {
	t.Run("when date field is not given, return error", func(t *testing.T) {
		validate := validator.New()
		validate.RegisterValidation("datetimeRFC3339", validateDatetimeRFC3339)
		target := struct {
			Date string `validate:"datetimeRFC3339"`
		}{}

		err := validate.Struct(target)

		assert.NotNil(t, err)
	})

	t.Run("when date field is incorrect, return error", func(t *testing.T) {
		validate := validator.New()
		validate.RegisterValidation("datetimeRFC3339", validateDatetimeRFC3339)
		target := struct {
			Date string `validate:"datetimeRFC3339"`
		}{Date: "This is not a date in here"}

		err := validate.Struct(target)

		assert.NotNil(t, err)
	})

	t.Run("when date field is correct, return no error", func(t *testing.T) {
		validate := validator.New()
		validate.RegisterValidation("datetimeRFC3339", validateDatetimeRFC3339)
		target := struct {
			Date string `validate:"datetimeRFC3339"`
		}{Date: "2023-04-20T14:00:00Z"}

		err := validate.Struct(target)

		assert.Nil(t, err)
	})
}
