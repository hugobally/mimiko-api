package validation

import (
	"errors"
)

type Input struct {
	Value string
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
