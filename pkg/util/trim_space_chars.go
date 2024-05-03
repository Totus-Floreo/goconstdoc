package util

import (
	"strings"
)

type buffer interface {
	String() string
	Reset()
	WriteString(string) (int, error)
}

func TrimSpaceCharacters(buf buffer) string {
	defer buf.Reset()

	temp := strings.ReplaceAll(buf.String(), "\r", "")
	buf.Reset()
	buf.WriteString(strings.ReplaceAll(temp, "\n", ""))

	temp = strings.ReplaceAll(buf.String(), "\t", "")
	buf.Reset()
	buf.WriteString(strings.ReplaceAll(temp, "\v", ""))

	temp = strings.ReplaceAll(buf.String(), "\f", "")
	buf.Reset()
	buf.WriteString(strings.ReplaceAll(temp, string(byte(0x85)), ""))

	return strings.ReplaceAll(buf.String(), string(byte(0xA0)), "")
}
