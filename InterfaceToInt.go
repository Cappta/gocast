package gocast

import (
	"errors"
	"strconv"
)

// InterfaceToInt converts an interface to integer
func InterfaceToInt(from interface{}) (int, error) {
	switch value := from.(type) {
	case float64:
		return int(value), nil
	case float32:
		return int(value), nil
	case int64:
		return int(value), nil
	case int32:
		return int(value), nil
	case int16:
		return int(value), nil
	case int8:
		return int(value), nil
	case uint64:
		return int(value), nil
	case uint32:
		return int(value), nil
	case uint16:
		return int(value), nil
	case uint8:
		return int(value), nil
	case int:
		return int(value), nil
	case uint:
		return int(value), nil
	case string:
		floatValue, err := strconv.ParseFloat(value, 64)
		return int(floatValue), err
	default:
		return 0, errors.New("Unknown type")
	}
}
