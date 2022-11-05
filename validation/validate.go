package validation

import (
	"errors"
	"regexp"
)

type Input struct {
	Value string
}

func AlphaNumExtra(input string) error {
	if !regexp.MustCompile(`^[\w -]`).MatchString(input) {
		return errors.New("forbidden characters")
	}

	return nil
}

func AlphaNum(input string) error {
	if !regexp.MustCompile(`^[\w]*$`).MatchString(input) {
		return errors.New("forbidden characters")
	}

	return nil
}

func Length(input string, min int, max int) error {
	l := len(input)
	if l < min {
		return errors.New("input too short")
	}
	if l > max {
		return errors.New("input too long")
	}
	return nil
}
