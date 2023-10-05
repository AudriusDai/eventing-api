package validation

import (
	iso6391 "github.com/emvi/iso-639-1"
	"github.com/go-playground/validator/v10"
)

var validateLanguageISO6391 validator.Func = func(fl validator.FieldLevel) bool {
	return iso6391.ValidName(fl.Field().String())
}
