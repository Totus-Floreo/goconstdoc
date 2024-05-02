/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package util

import "strings"

func RemoveCommentSymbols(input string) string {
	input = strings.ReplaceAll(input, "//", "")
	input = strings.ReplaceAll(input, "/*", "")
	input = strings.ReplaceAll(input, "*/", "")
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.TrimSpace(input)
	return input
}
