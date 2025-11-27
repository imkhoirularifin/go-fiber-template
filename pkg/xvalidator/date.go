package xvalidator

import (
	"regexp"

	val "github.com/go-playground/validator/v10"
)

var regexDate = `^\d{4}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$`

type DateValidator struct{}

func (v *DateValidator) Tag() string {
	return "x_date"
}

func (v *DateValidator) Func() val.Func {
	return func(fl val.FieldLevel) bool {
		if fl.Field().IsZero() {
			return true
		}

		regex := regexp.MustCompile(regexDate)
		return regex.MatchString(fl.Field().String())
	}
}

func (v *DateValidator) Translations() map[string]string {
	return map[string]string{
		"en-US": "Invalid date format, expected format: YYYY-MM-DD",
		"id":    "Format tanggal tidak valid, format yang diharapkan: YYYY-MM-DD",
	}
}
