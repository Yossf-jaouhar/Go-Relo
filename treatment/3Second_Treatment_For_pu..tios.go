package treatment

import (
	"strings"
)

// Second treatment of the data line by line and apply punctuations.
func SecondTreatement(res []string) []string {
	var result string

	var word string
	listponc := ",;:!.?"
	for i := 0; i < len(res); i++ {
		result += res[i] + " "
	}

	for j := 0; j < len(result); j++ {
		if j != 0 && strings.ContainsAny(string(result[j]), listponc) && result[j-1] == ' ' && result[j+1] == ' ' {
			result = result[:j-1] + result[j:]
		} else if j != 0 && strings.ContainsAny(string(result[j]), listponc) && result[j-1] == ' ' && result[j+1] != ' ' {
			result = result[:j-1] + string(result[j]) + " " + result[j+1:]
		} else if j != 0 && strings.ContainsAny(string(result[j]), listponc) && result[j-1] != ' ' && result[j+1] != ' ' {
			result = result[:j] + string(result[j]) + " " + result[j+1:]
		}
	}

	var res1 []string

	for _, char := range result {
		if char == ' ' {
			res1 = append(res1, word)
			word = ""
		} else {
			word += string(char)
		}
	}

	return res1
}
