package xvalidator

import (
	"unicode"

	val "github.com/go-playground/validator/v10"
)

type PasswordValidator struct{}

func (v *PasswordValidator) Tag() string {
	return "x_strong_password"
}

func (v *PasswordValidator) Func() val.Func {
	return func(fl val.FieldLevel) bool {
		if fl.Field().IsZero() {
			return true
		}

		password := fl.Field().String()

		// Check minimum length (8 characters)
		if len(password) < 8 {
			return false
		}

		// Check maximum length (128 characters)
		if len(password) > 128 {
			return false
		}

		var (
			hasUpper   = false
			hasLower   = false
			hasNumber  = false
			hasSpecial = false
		)

		// Check for character requirements
		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsNumber(char):
				hasNumber = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true
			}
		}

		// All requirements must be met
		return hasUpper && hasLower && hasNumber && hasSpecial
	}
}

func (v *PasswordValidator) Translations() map[string]string {
	return map[string]string{
		"en-US": "password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character",
		"id":    "kata sandi harus minimal 8 karakter dan mengandung setidaknya satu huruf besar, satu huruf kecil, satu angka, dan satu karakter khusus",
	}
}
