package project

import (
	"strings"
)

func CleanStr(data string) string {
	return strings.Join(strings.Fields(data), " ")
}