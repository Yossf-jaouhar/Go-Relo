package functions

import (
	"fmt"
	"strconv"
	"unicode"
)

func Capitalize(s string) string {
	res := ""
	for i, char := range s {
		if i == 0 {
			res += string(unicode.ToUpper(char))
		} else {
			res += string(unicode.ToLower(char))
		}
	}
	return res
}

func Upper(str string) string {
	res := ""
	for _, char := range str {
		res += string(unicode.ToUpper(char))
	}
	return res
}

func Lower(str string) string {
	res := ""
	for _, char := range str {
		res += string(unicode.ToLower(char))
	}
	return res
}

func Binary(binaryStr string, line int) string {
	decimalValue, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("you have a non-binary number in this line", ">>>", line, "<<<", "I will write this number as it is.")
		return binaryStr
	}
	return strconv.FormatInt(decimalValue, 10)
}
