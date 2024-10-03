package persianPhoneValidator

import (
	"regexp"
	"strings"
)

// Compile the regex once for efficiency
var nonDigitRegex = regexp.MustCompile("[^0-9]+")

func CleanPhoneNumber(phone string) string {
	// Remove all non-numeric characters
	cleanedPhone := nonDigitRegex.ReplaceAllString(phone, "")
	return cleanedPhone
}

func NormalizePhoneNumber(phone string) string {
	cleanedPhone := CleanPhoneNumber(phone)

	if len(cleanedPhone) == 0 {
		return ""
	}

	// Normalize phone number by handling different cases with switch
	switch {
	case strings.HasPrefix(cleanedPhone, "0098"):
		cleanedPhone = cleanedPhone[4:]
	case strings.HasPrefix(cleanedPhone, "98"):
		cleanedPhone = cleanedPhone[2:]
	case strings.HasPrefix(cleanedPhone, "+98"):
		cleanedPhone = cleanedPhone[3:]
	}

	// Ensure the phone number starts with '0' after normalization
	if !strings.HasPrefix(cleanedPhone, "0") {
		cleanedPhone = "0" + cleanedPhone
	}
	return cleanedPhone
}

func IsPhoneNumberValid(phone string) bool {
	// Normalize the phone number and check if it is valid
	normalPhone := NormalizePhoneNumber(phone)
	if len(normalPhone) == 11 && strings.HasPrefix(normalPhone, "09") {
		return true
	}
	return false
}

func ValidatePhoneNumber(phone string) (isvalid bool, normalizedPhoneNumber string) {
	// Normalize and validate the phone number
	normalizedPhoneNumber = NormalizePhoneNumber(phone)
	isvalid = IsPhoneNumberValid(normalizedPhoneNumber)
	return isvalid, normalizedPhoneNumber
}
