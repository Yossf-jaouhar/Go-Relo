package treatment

import (
	"strings"

	"goreloaded/functions"
)

// Split data line by line and return 2D array
func Div(data string) [][]string {
	var res string
	var newdata [][]string
	var myslice []string

	for _, char := range data {
		if char == '\n' {
			if len(res) > 0 {
				myslice = append(myslice, treatment(res))
				res = ""
			} else if len(res) == 0 {
				myslice = append(myslice, res)
			}
			if len(myslice) > 0 {

				newdata = append(newdata, myslice)
				myslice = nil
			}
		} else {
			res += string(char)
		}
	}

	if res != "" {
		myslice = append(myslice, treatment(res))
	}
	if len(myslice) > 0 {
		newdata = append(newdata, myslice)
	}

	return newdata
}

// treatment of the data line by line
func treatment(cell string) string {
	mynewslice := strings.Fields(cell)
	var result []string

	for i := 0; i < len(mynewslice); i++ {
		if mynewslice[i] == "(cap)" && len(result) > 0 {
			result[len(result)-1] = functions.Capitalize(result[len(result)-1])
		} else if mynewslice[i] == "(cap)" && len(result) == 0 {
		} else if mynewslice[i] == "(up)" && len(result) > 0 {
			result[len(result)-1] = functions.Upper(result[len(result)-1])
		} else if mynewslice[i] == "(up)" && len(result) == 0 {
		} else if len(mynewslice) >= 2 && mynewslice[i] == "(up," && isvalidflag(mynewslice[i+1]) {
			boool, intger := isvalidnumber(mynewslice[i+1])
			if boool {
				result = ff(result, intger)
				i += 2
			} else {
				result = append(result, mynewslice[i])
			}

		} else if mynewslice[i] == "(low)" && len(result) > 0 {
			result[len(result)-1] = functions.Lower(result[len(result)-1])
		} else if mynewslice[i] == "(low)" && len(result) == 0 {
		} else if mynewslice[i] == "(bin)" && len(result) > 0 {
			result[len(result)-1] = functions.Binary(result[len(result)-1])
		} else if mynewslice[i] == "(bin)" && len(result) == 0 {
		} else if mynewslice[i] == "(hex)" && len(result) > 0 {
			result[len(result)-1] = functions.Hexadecimal(result[len(result)-1])
		} else {
			result = append(result, mynewslice[i])
		}
	}

	return strings.Join(result, " ")
}

// // treatflag if "(flag,int)"
// func treatflag(part1, part2 string) string {
// 	if part1 == "(up," {
// 	} else if part1 == "(low," {
// 		flag := part1 + part2
// 		return flag
// 	} else if part1 == "(cap," {
// 		flag := part1 + part2
// 		return flagreturn false, 0
// 	} else {
// 		return part1
// 	}
// }

func isvalidflag(str string) bool {
	if str[len(str)-1] == ')' {
		return true
	} else {
		return false
	}
}

func isvalidnumber(str string) (bool, int) {
	res := str[:len(str)-1]
	ress := ""
	for i := 0; i < len(res); i++ {
		if res[i] >= '0' && res[i] <= '9' {
			ress += string(res[i])
		} else {
			return false, 0
		}
	}
	return true, functions.Atoi(ress)
}

func ff(res []string, intger int) []string {
	if intger <= len(res) {
		for j := len(res) - intger; j < len(res); j++ {
			res[j] = functions.Upper(res[j])
		}
	} else if intger != len(res) {
		for j := 0; j < len(res); j++ {
			res[j] = functions.Upper(res[j])
		}
	}
	return res
}
