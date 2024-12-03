package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"

	"github.com/surajNirala/job_portal_api/internal/models"
)

func GenerateFromPassword(charCount int) string {
	const digit = "0123456789abcdef"
	var password strings.Builder
	password.Grow(charCount)
	for i := 0; i < charCount; i++ {
		password.WriteByte(digit[rand.Intn(len(digit))])
	}
	return password.String()
}

func ValidatePasswordStrength(password string) (bool, []string) {
	validation := models.PasswordValidation{
		MinLength:  8,
		HasUpper:   true,
		HasLower:   true,
		HasNumber:  true,
		HasSpecial: true,
	}

	var validationErros []string
	// Check for number length
	if len(password) < validation.MinLength {
		validationErros = append(validationErros, fmt.Sprintf("Password must be atleast %d characers long.", validation.MinLength))
	}
	// Check for uppercase letter
	if validation.HasUpper && !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		validationErros = append(validationErros, "Password must contain atleast one uppercase letter.")
	}
	// Check for lowercase letter
	if validation.HasLower && !strings.ContainsAny(password, "abcdefghijklmzopqrstuvwxyz") {
		validationErros = append(validationErros, "Password must contain atleast one lowercase letter.")
	}
	// Check for number
	if validation.HasNumber && !strings.ContainsAny(password, "0123456789") {
		validationErros = append(validationErros, "Password must contain atleast one number.")
	}
	// Check for special characters
	if validation.HasSpecial {
		hasspecial := false
		for _, char := range password {
			if unicode.IsPunct(char) || unicode.IsSymbol(char) {
				hasspecial = true
				break
			}
		}
		if !hasspecial {
			validationErros = append(validationErros, "Password must contain atleast one special character.")
		}
	}
	return len(validationErros) == 0, validationErros
}
