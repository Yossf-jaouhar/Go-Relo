package functions

import (
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
// Check the word
func Capitalizeslice(res []string) []string {
	ff := false 
	for i:= len(res)-1 ; i >= 0; i-- {
		if !(ff) {
			for _, char :=  range res[i] {
				if char >= 'a'  && char <= 'z' ||  char >= 'A' && char  <= 'Z' {
					res[i] = Capitalize(res[i])
					ff  = true
					
				} 
			}

		} else {
			break
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

func Upperslice(res []string) []string {
	ff := false 
	for i:= len(res)-1 ; i >= 0; i-- {
		if !(ff) {
			for _, char :=  range res[i] {
				if char >= 'a'  && char <= 'z' ||  char >= 'A' && char  <= 'Z' {
					res[i] = Upper(res[i])
					ff  = true
					
				} 
			}

		} else {
			break
		}
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

func Lowerslice(res []string) []string {
	ff := false 
	for i:= len(res)-1 ; i >= 0; i-- {
		if !(ff) {
			for _, char :=  range res[i] {
				if char >= 'a'  && char <= 'z' ||  char >= 'A' && char  <= 'Z' {
					res[i] = Lower(res[i])
					ff  = true
					
				} 
			}

		} else {
			break
		}
	}	
	return res
}

func Binary(binaryStr string) string {
	decimalValue, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		// fmt.Println("you have a >>non-binary<< number in your text, I will write t= i -2his number as it is.")
		return binaryStr
	}
	return strconv.FormatInt(decimalValue, 10)
}

func Hexadecimal(hexadecimalStr string) string {
	decimalValue, err := strconv.ParseInt(hexadecimalStr, 16, 64)
	if err != nil {
		// fmt.Println("you have a non-hexadecimal number in your text, I will write this number as it is.")
		return hexadecimalStr
	}
	return strconv.FormatInt(decimalValue, 10)
}
