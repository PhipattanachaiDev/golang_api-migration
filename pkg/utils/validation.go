package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	re := regexp.MustCompile(`^[a-z0-9._%%+-]+@[a-z0-9.-]+\\.[a-z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidateRequired(field, name string) error {
	if field == "" {
		return errors.New(name + " is required")
	}
	return nil
}