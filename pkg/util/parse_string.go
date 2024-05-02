package util

import (
	"strconv"
)

func ParseStringToType(str string) interface{} {
	intType, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return intType
	}

	floatType, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return floatType
	}

	boolType, err := strconv.ParseBool(str)
	if err == nil {
		return boolType
	}

	return str
}
