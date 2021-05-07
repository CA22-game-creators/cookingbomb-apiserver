package testdata

import (
	"strings"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
)

var Name = struct {
	Valid, Invalid, TooShort, TooLong string
}{
	Valid:    "name",
	Invalid:  string([]byte{0xff, 0xfe, 0xfd}),
	TooShort: strings.Repeat("a", user.NameMinLen-1),
	TooLong:  strings.Repeat("a", user.NameMaxLen+1),
}
