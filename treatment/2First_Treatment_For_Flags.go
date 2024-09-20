package treatment

import (
	"strconv"
	"strings"

	"goreloaded/functions"
)

// First Treatment of the data line by line and apply (cap),(up),(low)
func Treatment(cell string) string {
	mynewslice := strings.Fields(cell)
	var result []string

	for i := 0; i < len(mynewslice); i++ {
		if len(result) > 0 && isValidflag(mynewslice[i]) {
			switch mynewslice[i] {
			case "(up)":
				result = functions.Change_To_Up(result, 1)
			case "(cap)":
				result = functions.Change_To_Cap(result, 1)
			case "(low)":
				result = functions.Change_To_Low(result, 1)
			case "(bin)":
				result = functions.Binary(result, 1)
			case "(hex)":
				result = functions.Hexadecimal(result, 1)
			}
		} else if len(result) == 0 && isValidflag(mynewslice[i]) {
			continue
		} else if i < len(mynewslice)-1 && Flag_part1(mynewslice[i]) && Flag_part2(mynewslice[i+1]) {
			boool, intger := isvalidnumber(mynewslice[i+1])
			if boool {
				switch mynewslice[i] {
				case "(cap,":
					result = functions.Change_To_Cap(result, intger)
				case "(up,":
					result = functions.Change_To_Up(result, intger)
				case "(low,":
					result = functions.Change_To_Low(result, intger)
				}
				i += 1

			} else {
				result = append(result, mynewslice[i])
			}
		} else {
			result = append(result, mynewslice[i])
		}
	}

	result = SecondTreatement(result)
	result = Third_Treatment(result)
	result = Last_Treatment(result)
	StrRes := strings.Join(result, " ")
	return StrRes
}

// isValidFalgs just for (up) (cap) (low) (bin) (hex).
func isValidflag(flag string) bool {
	switch flag {
	case "(cap)", "(up)", "(low)", "(bin)", "(hex)":
		return true
	}
	return false
}

// check the part1 of the flag, isValid: (up, (cap, (low, 
func Flag_part1(flag string) bool {
	switch flag {
	case "(up,", "(low,", "(cap,":
		return true
	}
	return false
}

// check the part2 of the flag, isValid: 2) , 3) ....
func Flag_part2(str string) bool {
	if str[len(str)-1] == ')' {
		return true
	} else {
		return false
	}
}

// check the number in the flags and convert it to a intger.
func isvalidnumber(str string) (bool, int) {
	res := str[:len(str)-1]
	if res == "" {
		return false, 0
	}
	ress := ""
	for i := 0; i < len(res); i++ {
		if res[i] == '+' && i == 0 {
			ress += "+"
		} else if res[i] >= '0' && res[i] <= '9' {
			ress += string(res[i])
		} else {
			return false, 0
		}
	}
	intger, _ := strconv.Atoi(ress)
	return true, intger
}
