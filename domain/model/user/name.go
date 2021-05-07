package user

import (
	"fmt"
	"unicode/utf8"
)

type Name string

const NameMinLen = 1
const NameMaxLen = 10

func NewName(v string) (Name, error) {
	if utf8.RuneCountInString(v) < 1 || utf8.RuneCountInString(v) > 10 {
		return "", fmt.Errorf("user name should be %d to %d characters", NameMinLen, NameMaxLen)
	}

	return Name(v), nil
}
