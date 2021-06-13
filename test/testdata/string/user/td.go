package testdata

import (
	"strings"
)

var Name = struct {
	Valid, Invalid, InvalidChar, TooShort, TooLong string
}{
	Valid:       "name",
	Invalid:     `\m\t@#$%`,
	InvalidChar: string([]byte{0xff, 0xfe, 0xfd}),
	TooShort:    "",
	TooLong:     strings.Repeat("a", 100),
}
