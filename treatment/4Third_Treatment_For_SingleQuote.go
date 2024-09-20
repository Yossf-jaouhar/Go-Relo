package treatment

import (
	"strings"
)

// // Third treatment of the data line by line and apply the singleQuote
func Third_Treatment(res []string) []string {
	var result string
	for i := 0; i < len(res); i++ {
		result += res[i] + " "
	}

	Changement := Single_Cote(result)
	mynewslice := strings.Fields(Changement)
	return mynewslice
}

// Search for the single quote in the string and replace it with a double quote
func Single_Cote(content string) string {
	new_content := Split_Single_Cote(content)

	if len(new_content) == 0 {
		return content
	}

	str := ""
	for _, arg := range new_content {
		if arg[0] == '\'' && arg[len(arg)-1] == '\'' && arg != "'" {
			str += "'" + strings.TrimSpace(arg[1:len(arg)-1]) + "' "
		} else {
			str += arg + " "
		}
	}

	return str
}

// Split the string by single quote
func Split_Single_Cote(Data string) []string {
	var result []string
	firstindex := 0
	lastindex := 0
	Helper := false
	counter := 0

	for i, char := range Data {
		if char == '\'' {
			if i > 0 && i < len(Data)-1 && isAlpha(rune(Data[i-1])) && isAlpha(rune(Data[i+1])) {
				continue
			}

			if Helper {
				Helper = false
				lastindex = i + 1
				result = append(result, Data[firstindex:lastindex])
				counter++
			} else {

				Helper = true
				firstindex = i
				if lastindex < firstindex {
					result = append(result, Data[lastindex:firstindex])
				}
			}
		}
	}

	if Helper {
		result = append(result, Data[firstindex:])
	} else if len(Data) > lastindex {
		result = append(result, Data[lastindex:])
	}

	return result
}

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
