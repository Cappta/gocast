package gocast

import (
	"unicode/utf16"
)

// FromUInts16ToString converts an UInt16 slice to string using UTF16 format
func FromUInts16ToString(from []uint16) string {
	return string(utf16.Decode(from))
}
