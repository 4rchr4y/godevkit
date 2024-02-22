package basic

import "strings"

func EmptyStr(value string) bool {
	return strings.TrimSpace(value) != ""
}

func Exists(source map[any]any, value any) bool {
	_, exists := source[value]
	return exists
}
