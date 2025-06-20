package xvalidator

import (
	"fmt"
	"go-fiber-template/internal/domain/dto"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// Validator holds the validator and translator instances used for struct validation.
type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

// ValidatorOption defines a functional option for configuring the validator instance.
type ValidatorOption func(*validator.Validate, ut.Translator) error

// NewValidator creates a new Validator instance with the provided options.
func NewValidator(opts ...ValidatorOption) (*Validator, error) {
	v := validator.New(validator.WithRequiredStructEnabled())

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "" {
			name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		}
		if name == "-" {
			return ""
		}
		return name
	})

	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)

	translator, found := uni.GetTranslator("en")
	if !found {
		return nil, fmt.Errorf("translator not found for locale 'en'")
	}

	if err := enTranslations.RegisterDefaultTranslations(v, translator); err != nil {
		return nil, err
	}

	for _, opt := range opts {
		if err := opt(v, translator); err != nil {
			return nil, err
		}
	}

	return &Validator{
		validate:   v,
		translator: translator,
	}, nil
}

// WithCustomValidator registers a custom validator and its translation function.
func WithCustomValidator(cv CustomValidator) ValidatorOption {
	return func(v *validator.Validate, translator ut.Translator) error {
		if err := v.RegisterValidation(cv.Tag(), cv.Func()); err != nil {
			return err
		}

		translationText, customTransFunc := cv.Translation()

		if translationText == "" || customTransFunc == nil {
			return nil
		}

		registerFn := func(ut ut.Translator) error {
			return ut.Add(cv.Tag(), translationText, true)
		}

		return v.RegisterTranslation(cv.Tag(), translator, registerFn, customTransFunc)
	}
}

// ValidateStruct validates the provided struct against the registered validation rules.
func (v *Validator) ValidateStruct(s interface{}) []dto.ErrorValidationDto {
	var errValidations []dto.ErrorValidationDto

	if err := v.validate.Struct(s); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				namespace := fe.Namespace()
				field := trimStructName(namespace)
				msg := fe.Translate(v.translator)

				errValidations = append(errValidations, dto.ErrorValidationDto{
					Field:   field,
					Message: msg,
				})
			}

			return errValidations
		}

		return []dto.ErrorValidationDto{
			{
				Field:   "unknown",
				Message: err.Error(),
			},
		}
	}

	return nil
}

func trimStructName(field string) string {
	// Remove the struct name prefix if it exists
	if idx := strings.Index(field, "."); idx != -1 {
		return field[idx+1:]
	}
	return field
}
