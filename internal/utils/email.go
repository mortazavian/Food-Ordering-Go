package utils

import "regexp"

func IsValidEmail(email string) bool {
	// Regular expression for validating email addresses
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	pattern := regexp.MustCompile(emailRegex)

	// Use the MatchString method to check if the email matches the pattern
	return pattern.MatchString(email)
}
