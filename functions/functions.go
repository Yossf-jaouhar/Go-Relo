package functions

import (
	"strconv"
	"unicode"
)

// Capitalize the string
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

// capitalize the sllice  of strings by the number
func Change_To_Cap(res []string, intger int) []string {
	if intger < len(res) {
		for i := len(res) - 1; i >= 0; i-- {
			if intger > 0 {
				if isposible(res[i]) {
					res[i] = Capitalize(res[i])
					intger--
				}
			} else if intger == 0 {
				break
			}
		}
	} else {
		for j := 0; j < len(res); j++ {
			res[j] = Capitalize(res[j])
		}
	}
	return res
}

// upper the string
func Upper(str string) string {
	res := ""
	for _, char := range str {
		res += string(unicode.ToUpper(char))
	}
	return res
}

// Upper the sllice  of strings by the number
func Change_To_Up(res []string, intger int) []string {
	if intger < len(res) {
		for i := len(res) - 1; i >= 0; i-- {
			if intger > 0 {
				if isposible(res[i]) {
					res[i] = Upper(res[i])
					intger--
				}
			} else if intger == 0 {
				break
			}
		}
	} else {
		for j := 0; j < len(res); j++ {
			res[j] = Upper(res[j])
		}
	}
	return res
}

// Lower the string
func Lower(str string) string {
	res := ""
	for _, char := range str {
		res += string(unicode.ToLower(char))
	}
	return res
}

// Lower the slice  of strings by the number
func Change_To_Low(res []string, intger int) []string {
	if intger < len(res) {
		for i := len(res) - 1; i >= 0; i-- {
			if intger > 0 {
				if isposible(res[i]) {
					res[i] = Lower(res[i])
					intger--
				}
			} else if intger == 0 {
				break
			}
		}
	} else {
		for j := 0; j < len(res); j++ {
			res[j] = Lower(res[j])
		}
	}
	return res
}

// Check if the string is possible to change
func isposible(str string) bool {
	for _, char := range str {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

// Change the number to binary
func Binary(binaryStr string) string {
	decimalValue, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return binaryStr
	}
	return strconv.FormatInt(decimalValue, 10)
}

// Change the number to Hexadecimal
func Hexadecimal(hexadecimalStr string) string {
	decimalValue, err := strconv.ParseInt(hexadecimalStr, 16, 64)
	if err != nil {
		return hexadecimalStr
	}
	return strconv.FormatInt(decimalValue, 10)
}

// convert string to integer
func Atoi(s string) int {
	negative := false
	res := 0
	start := 0
	new := []rune(s)
	if len(s) == 0 {
		return 1
	}
	if new[0] == '-' {
		negative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	for i := start; i < len(new); i++ {
		if new[i] < '0' || new[i] > '9' {
			return 0
		}
		res = res*10 + int(new[i]-'0')
	}
	if negative {
		return -res
	}
	return res
}
