package util

import (
	"strings"
)

var banned = map[string]bool{
	"\r":               true,
	"\n":               true,
	"\t":               true,
	"\v":               true,
	"\f":               true,
	string(byte(0x85)): true,
	string(byte(0xA0)): true,
}

func TrimSpaceCharacters(temp string) string {
	var buf strings.Builder

	for _, chr := range temp {
		if _, ok := banned[string(chr)]; !ok {
			buf.WriteRune(chr)
		}
	}

	return buf.String()
}
