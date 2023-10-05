package validation

import (
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var BindValidators = func() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("datetimeRFC3339", validateDatetimeRFC3339); err != nil {
			log.Fatal("failed to register datetime validation")
		}
		if err := v.RegisterValidation("languageISO6391", validateLanguageISO6391); err != nil {
			log.Fatal("failed to register datetime validation")
		}

	}
}
