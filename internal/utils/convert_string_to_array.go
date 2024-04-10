package utils

import (
	"strings"
)

func ConvertEnvStringToArray(s string) []string {
	var arr []string
	arr = append(arr, strings.Split(s, ",")...)
	return arr
}
