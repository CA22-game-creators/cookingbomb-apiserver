package testdata

import (
	"strings"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
)

var Name = struct {
	Valid, TooShort, TooLong string
}{
	Valid:    "name",
	TooShort: strings.Repeat("a", user.NameMinLen-1),
	TooLong:  strings.Repeat("a", user.NameMaxLen+1),
}
