package functions

import (
	"strconv"
	"unicode"
)

// Capitalize of the string.
func Capitalize(s string) string {
	mySlice := []rune(s)
	next := true
	for i := 0; i < len(mySlice); i++ {
		if unicode.IsLetter(mySlice[0]) && next {
			mySlice[i] = unicode.ToUpper(mySlice[i])
			next = false
		} else if next {
			mySlice[i] = unicode.ToLower(mySlice[i])

			next = true
		}
	}
	return string(mySlice)
}

// capitalize of the slice  of strings by the number.
func Change_To_Cap(res []string, intger int) []string {
	if intger < len(res) {
		for i := len(res) - 1; i >= 0; i-- {
			if intger > 0 {

				res[i] = Capitalize(res[i])
				intger--

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

// upper of the string.
func Upper(str string) string {
	res := ""
	for _, char := range str {
		res += string(unicode.ToUpper(char))
	}
	return res
}

// Upper of the slice of strings by the number.
func Change_To_Up(res []string, intger int) []string {
	if intger < len(res) {
		for i := len(res) - 1; i >= 0; i-- {
			if intger > 0 {

				res[i] = Upper(res[i])
				intger--

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

// Lower of the string.
func Lower(str string) string {
	res := ""
	for _, char := range str {
		res += string(unicode.ToLower(char))
	}
	return res
}

// Lower of the slice of strings by the number.
func Change_To_Low(res []string, intger int) []string {
	if intger < len(res) {
		for i := len(res) - 1; i >= 0; i-- {
			if intger > 0 {
				res[i] = Lower(res[i])
				intger--
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

// Binary converts the first binary string in the slice to a decimal number if "intger" is 1.
func Binary(res []string, intger int) []string {
	for i := len(res) - 1; i >= 0; i-- {
		if intger == 1 && isBinary(res[i]) {
			decimalValue, _ := strconv.ParseInt(res[i], 2, 64)
			res[i] = strconv.FormatInt(decimalValue, 10)
			intger--
		} else if intger == 0 {
			break
		}
	}

	return res
}

// Helper function to check if a string is a valid "binary" number.
func isBinary(str string) bool {
	for _, char := range str {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}

// Hexadecimal converts the first hexadecimal string in the slice to a decimal number if "intger" is 1.
func Hexadecimal(res []string, intger int) []string {
	for i := len(res) - 1; i >= 0; i-- {
		if intger == 1 && isHexadecimal(res[i]) {
			decimalValue, _ := strconv.ParseInt(res[i], 16, 64)
			res[i] = strconv.FormatInt(decimalValue, 10)
			intger--
		} else if intger == 0 {
			break
		}
	}

	return res
}

// Helper function to check if a string is a valid "hexadecimal" number.
func isHexadecimal(s string) bool {
	_, err := strconv.ParseInt(s, 16, 64)
	return err == nil
}
