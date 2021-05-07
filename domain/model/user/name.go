package user

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Name string

const NameMinLen = 1
const NameMaxLen = 10

func NewName(v string) (Name, error) {
	if !utf8.ValidString(v) {
		return "", errors.New("user name string is invalid")
	}
	if utf8.RuneCountInString(v) < NameMinLen || utf8.RuneCountInString(v) > NameMaxLen {
		return "", fmt.Errorf("user name should be %d to %d characters", NameMinLen, NameMaxLen)
	}

	return Name(v), nil
}
