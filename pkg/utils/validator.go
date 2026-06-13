package utils

import "github.com/go-playground/validator/v10"

func ValidateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	for _, c := range username {
		if !((c >= 'A' && c <= 'z') || c == '_' || (c >= '0' && c <= '9')) {
			return false
		}
	}

	return true
}
