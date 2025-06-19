package xvalidator

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
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

func (v *DateValidator) Translation() (string, val.TranslationFunc) {
	msg := "Invalid Date Format, Standard Format: YYYY-MM-DD"
	return msg, func(ut ut.Translator, fe val.FieldError) string {
		return msg
	}
}
