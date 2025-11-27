package xvalidator

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/lib/constant"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

// Client holds the validator and translator instances used for struct validation.
type Client struct {
	validate    *validator.Validate
	translators map[string]ut.Translator
}

// NewClient creates a new Client instance with the provided options.
func NewClient(opts ...ValidatorOption) *Client {
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

	// Setup multiple locales
	enLocale := en.New()
	idLocale := id.New()
	uni := ut.New(enLocale, enLocale, idLocale)

	translators := make(map[string]ut.Translator)

	// English translator
	enTranslator, found := uni.GetTranslator("en")
	if !found {
		panic("english translator not found")
	}
	if err := enTranslations.RegisterDefaultTranslations(v, enTranslator); err != nil {
		panic(err)
	}
	translators["en-US"] = enTranslator

	// Indonesian translator
	idTranslator, found := uni.GetTranslator("id")
	if !found {
		panic("indonesian translator not found")
	}
	if err := idTranslations.RegisterDefaultTranslations(v, idTranslator); err != nil {
		panic(err)
	}
	translators["id"] = idTranslator

	for _, opt := range opts {
		if err := opt(v, translators); err != nil {
			panic(err)
		}
	}

	return &Client{
		validate:    v,
		translators: translators,
	}
}

// WithCustomValidator registers a custom validator and its translation function.
func WithCustomValidator(cv CustomValidator) ValidatorOption {
	return func(v *validator.Validate, translators map[string]ut.Translator) error {
		if err := v.RegisterValidation(cv.Tag(), cv.Func()); err != nil {
			return err
		}

		// Register translations for all supported languages
		translations := cv.Translations()

		for lang, translator := range translators {
			translationText, found := translations[lang]
			if !found {
				// Fallback to Indonesian if translation not found
				if fallback, exists := translations["id"]; exists {
					translationText = fallback
				} else {
					continue
				}
			}

			registerFn := func(ut ut.Translator) error {
				return ut.Add(cv.Tag(), translationText, true)
			}

			customTransFunc := func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T(cv.Tag(), fe.Field())
				return t
			}

			if err := v.RegisterTranslation(cv.Tag(), translator, registerFn, customTransFunc); err != nil {
				return err
			}
		}

		return nil
	}
}

// ValidateStruct validates the provided struct against the registered validation rules.
func (v *Client) ValidateStruct(s interface{}) []dto.ErrorValidationDto {
	return v.ValidateStructWithLang(s, constant.DefaultLanguage)
}

// ValidateStructWithLang validates the provided struct with a specific language.
func (v *Client) ValidateStructWithLang(s interface{}, lang string) []dto.ErrorValidationDto {
	var errValidations []dto.ErrorValidationDto

	translator, exists := v.translators[lang]
	if !exists {
		// Fallback to default language if language not supported
		translator = v.translators[constant.DefaultLanguage]
	}

	if err := v.validate.Struct(s); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				namespace := fe.Namespace()
				field := trimStructName(namespace)
				msg := fe.Translate(translator)

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
