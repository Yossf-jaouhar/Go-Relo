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

// treatment of the data line by line and apply (cap),(up),(low)
func treatment(cell string) string {
	mynewslice := strings.Fields(cell)
	var result []string

	for i := 0; i < len(mynewslice); i++ {
		if mynewslice[i] == "(cap)" && len(result) > 0 {
			result = functions.Change_To_Cap(result, 1)
		} else if mynewslice[i] == "(cap)" && len(result) == 0 {
		} else if len(mynewslice) >= 2 && mynewslice[i] == "(cap," && isvalidflag(mynewslice[i+1]) {
			boool, intger := isvalidnumber(mynewslice[i+1])
			if boool {
				result = functions.Change_To_Cap(result, intger)
				i += 1
			} else {
				result = append(result, mynewslice[i])
			}
		} else if mynewslice[i] == "(up)" && len(result) > 0 {
			result = functions.Change_To_Up(result, 1)
		} else if mynewslice[i] == "(up)" && len(result) == 0 {
		} else if len(mynewslice) >= 2 && mynewslice[i] == "(up," && isvalidflag(mynewslice[i+1]) {
			boool, intger := isvalidnumber(mynewslice[i+1])
			if boool {
				result = functions.Change_To_Up(result, intger)
				i += 1
			} else {
				result = append(result, mynewslice[i])
			}

		} else if mynewslice[i] == "(low)" && len(result) > 0 {
			result = functions.Change_To_Low(result, 1)
		} else if mynewslice[i] == "(low)" && len(result) == 0 {
		} else if len(mynewslice) >= 2 && mynewslice[i] == "(low," && isvalidflag(mynewslice[i+1]) {
			boool, intger := isvalidnumber(mynewslice[i+1])
			if boool {
				result = functions.Change_To_Low(result, intger)
				i += 1
			} else {
				result = append(result, mynewslice[i])
			}
		} else if mynewslice[i] == "(bin)" && len(result) > 0 {
			result[len(result)-1] = functions.Binary(result[len(result)-1])
		} else if mynewslice[i] == "(bin)" && len(result) == 0 {
		} else if mynewslice[i] == "(hex)" && len(result) > 0 {
			result[len(result)-1] = functions.Hexadecimal(result[len(result)-1])
		} else if mynewslice[i] == "(hex)" && len(result) == 0 {
		} else {
			result = append(result, mynewslice[i])
		}
	}
	Second_treatment := SecondTreatment(result)
	return strings.Join(Second_treatment, " ")
}

// check the flag
func isvalidflag(str string) bool {
	if str[len(str)-1] == ')' {
		return true
	} else {
		return false
	}
}

// check the number in the flag
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

// second treatment of the data line by line and apply "a"
func SecondTreatment(res []string) []string {
	var result []string
	for i := 0; i < len(res); i++ {
		if i < len(res)-1 && res[i] == "a" && isvalid_STRING(res[i+1]) {
			result = append(result, res[i]+"n")
		} else if i < len(res)-1 && res[i] == "A" && isvalid_STRING(res[i+1]) {
			result = append(result, res[i]+"N")
		} else {
			result = append(result, res[i])
		}
	}
	return Third_treatment(result)
}

// check the string to append "n" or "N"
func isvalid_STRING(str string) bool {
	g := "aAeEiIoOuUhH"
	if strings.Contains(g, string(str[0])) {
		return true
	} else {
		return false
	}
}

func Third_treatment(res []string) []string {
	return res
}
