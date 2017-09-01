package gocast

import (
	"unicode/utf16"
)

// FromStringToUInts16 converts a string to an UInt16 slice using UTF16 format
func FromStringToUInts16(from string) []uint16 {
	return utf16.Encode([]rune(from))
}
