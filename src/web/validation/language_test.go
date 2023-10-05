package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateLanguageISO6391(t *testing.T) {
	t.Run("when language field is not given, return error", func(t *testing.T) {
		validate := validator.New()
		validate.RegisterValidation("languageISO6391", validateLanguageISO6391)
		target := struct {
			Language string `validate:"languageISO6391"`
		}{}

		err := validate.Struct(target)

		assert.NotNil(t, err)
	})

	t.Run("when language field is incorrect, return error", func(t *testing.T) {
		validate := validator.New()
		validate.RegisterValidation("languageISO6391", validateLanguageISO6391)
		target := struct {
			Language string `validate:"languageISO6391"`
		}{Language: "not a language"}

		err := validate.Struct(target)

		assert.NotNil(t, err)
	})

	t.Run("when language field is correct, return no error", func(t *testing.T) {
		validate := validator.New()
		validate.RegisterValidation("languageISO6391", validateLanguageISO6391)
		target := struct {
			Language string `validate:"languageISO6391"`
		}{Language: "Lithuanian"}

		err := validate.Struct(target)

		assert.Nil(t, err)
	})
}
