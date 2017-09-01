package gocast

// FromUInts16ToBytes converts an uint16 slice to a byte slice
func FromUInts16ToBytes(from []uint16) []byte {
	result := make([]byte, len(from)*2)
	for i := 0; i < len(from); i++ {
		resultIndex := i * 2
		result[resultIndex] = byte(from[i])
		result[resultIndex+1] = byte(from[i] >> 8)
	}
	return result
}
