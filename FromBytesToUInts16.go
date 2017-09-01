package gocast

import "errors"

// FromBytesToUInts16 converts data to uint16 format
func FromBytesToUInts16(data []byte) (result []uint16, err error) {
	if len(data)%2 != 0 {
		return nil, errors.New("Invalid data")
	}
	length := len(data) / 2
	result = make([]uint16, length)
	for i := 0; i < length; i++ {
		dataIndex := i * 2
		result[i] = uint16(data[dataIndex]) + uint16(data[dataIndex+1])<<8
	}
	return
}
