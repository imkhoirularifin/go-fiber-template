package xvalidator

import (
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
)

// ValidatorOption defines a functional option for configuring the validator instance.
type ValidatorOption func(*validator.Validate, map[string]ut.Translator) error

// CustomValidator defines the interface that custom validators must implement. It requires methods to return the validation tag, function, and translation details.
type CustomValidator interface {
	// Tag returns the tag identifier used in struct field validation tags (e.g., `validate:"tag"`).
	Tag() string
	// Func returns the validator.Func that performs the validation logic.
	Func() validator.Func
	// Translations returns a map of language codes to translation messages.
	// Supported languages: "en" (English), "id" (Indonesian)
	Translations() map[string]string
}
