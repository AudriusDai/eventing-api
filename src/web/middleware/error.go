package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/audriusdai/eventing-api/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type WebErrorResponse struct {
	Description string `json:"description"`
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)

		if len(detectedErrors) == 0 {
			return
		}

		err := detectedErrors[0].Err

		switch err := err.(type) {
		case validator.ValidationErrors:
			c.JSON(http.StatusBadRequest, WebErrorResponse{Description: GetErrorText(err)})
			return
		}

		c.JSON(http.StatusInternalServerError, WebErrorResponse{Description: "Internal Server Error"})
	}
}

func GetErrorText(verrs validator.ValidationErrors) string {
	l := []string{}
	for _, e := range verrs {
		l = append(l, validationErrorToText(e))
	}

	if len(l) == 0 {
		return ""
	}

	l[0] = util.Capitalize(l[0])

	return strings.Join(l, ", ") + "."
}

var validationErrorToText = func(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("field `%s` is required", e.Field())
	case "max":
		return fmt.Sprintf("field `%s` max size %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("field `%s` min size %s", e.Field(), e.Param())
	case "gt":
		return fmt.Sprintf("field `%s` should be greater than %s", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("field `%s` should be less/equal than %s", e.Field(), e.Param())
	case "email":
		return fmt.Sprintf("value `%s` is an invalid email", e.Value())
	case "len":
		return fmt.Sprintf("field `%s` must be %s characters long", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf(
			"value `%s` is invalid for field `%s` (%s)",
			e.Value(),
			e.Field(),
			strings.Join(strings.Split(e.Param(), " "), ", "),
		)
	case "datetimeRFC3339":
		return fmt.Sprintf("value `%s` is invalid date (RFC3339)", e.Value())
	case "languageISO6391":
		return fmt.Sprintf("value `%s` is invalid language (ISO6391)", e.Value())
	}

	return fmt.Sprintf("field `%s` is invalid", e.Field())
}
