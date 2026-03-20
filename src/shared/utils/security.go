package utils

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

type PasswordRequirements struct {
	MinLength      int  `json:"min_length"`
	RequireUpper   bool `json:"require_upper"`
	RequireLower   bool `json:"require_lower"`
	RequireNumber  bool `json:"require_number"`
	RequireSpecial bool `json:"require_special"`
}
type PasswordValidationResult struct {
	IsValid bool     `json:"is_valid"`
	Errors  []string `json:"errors"`
}

func getPasswordRequirements() PasswordRequirements {
	minLength, _ := strconv.Atoi(os.Getenv("PASSWORD_MIN_LENGTH"))
	if minLength == 0 {
		minLength = 8
	}
	return PasswordRequirements{
		MinLength:      minLength,
		RequireUpper:   os.Getenv("PASSWORD_REQUIRE_UPPER") != "false",
		RequireLower:   os.Getenv("PASSWORD_REQUIRE_LOWER") != "false",
		RequireNumber:  os.Getenv("PASSWORD_REQUIRE_NUMBER") != "false",
		RequireSpecial: os.Getenv("PASSWORD_REQUIRE_SPECIAL") != "false",
	}
}
func ValidatePassword(password string) PasswordValidationResult {
	reqs := getPasswordRequirements()
	var errors []string
	if len(password) < reqs.MinLength {
		errors = append(errors, "Password must be at least "+strconv.Itoa(reqs.MinLength)+" characters long")
	}
	if reqs.RequireUpper && !containsUppercase(password) {
		errors = append(errors, "Password must contain at least one uppercase letter")
	}
	if reqs.RequireLower && !containsLowercase(password) {
		errors = append(errors, "Password must contain at least one lowercase letter")
	}
	if reqs.RequireNumber && !containsNumber(password) {
		errors = append(errors, "Password must contain at least one number")
	}
	if reqs.RequireSpecial && !containsSpecialCharacter(password) {
		errors = append(errors, "Password must contain at least one special character")
	}

	weakPasswords := []string{"password", "123456", "qwerty", "admin", "letmein", "password123", "admin123", "12345678", "welcome"}
	if slices.Contains(weakPasswords, strings.ToLower(password)) {
		errors = append(errors, "Password is too common and easily guessable")
	}

	return PasswordValidationResult{
		IsValid: len(errors) == 0,
		Errors:  errors,
	}

}
func containsUppercase(s string) bool {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}
func containsLowercase(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}
func containsNumber(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {

			return true
		}
	}
	return false
}
func containsSpecialCharacter(s string) bool {
	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~"
	for _, char := range s {
		if containsRune(specialChars, char) {
			return true
		}
	}
	return false
}
func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
