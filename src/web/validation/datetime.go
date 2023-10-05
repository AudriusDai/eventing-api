package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validateDatetimeRFC3339 validator.Func = func(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
