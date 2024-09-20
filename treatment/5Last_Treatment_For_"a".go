package treatment

import (
	"strings"
)

// Last treatment of the data line by line and apply "a" or "A".
func Last_Treatment(res []string) []string {
	var result []string
	if len(res) < 2 {
		return res
	}

	for i := 0; i < len(res)-1; i++ {
		if (res[i] == "a" || res[i] == "'a" ) && isvalid_STRING(res[i+1]) {
			result = append(result, res[i]+"n")
		} else if (res[i] == "A" || res[i] == "'A") && isvalid_STRING(res[i+1]) {
			result = append(result, res[i]+"n")
		} else {
			result = append(result, res[i])
		}
	}
	result = append(result, res[len(res)-1])

	return result
}

// check the string to append "n".
func isvalid_STRING(str string) bool {
	ValidCharacters := "aAeEiIoOuUhH"
	
	if len(str) == 0 {
		return false
	}
	
	return strings.Contains(ValidCharacters, string(str[0]))
}
