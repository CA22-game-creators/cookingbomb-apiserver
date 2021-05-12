package user

import (
	"fmt"
	"unicode/utf8"

	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

type Name string

const NameMinLen = 1
const NameMaxLen = 10

func NewName(v string) (Name, error) {
	if !utf8.ValidString(v) {
		return "", errors.InvalidArgument("user name string is invalid")
	}
	if utf8.RuneCountInString(v) < NameMinLen || utf8.RuneCountInString(v) > NameMaxLen {
		return "", errors.InvalidArgument(
			fmt.Sprintf("user name should be %d to %d characters", NameMinLen, NameMaxLen),
		)
	}

	return Name(v), nil
}
