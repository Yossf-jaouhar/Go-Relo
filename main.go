package main

import (
	"os"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}
	data2 := div(string(data))

	err = os.WriteFile("result.txt", []byte(data2), 0o666)
	if err != nil {
		panic(err)
	}
}

func div(data string) string {
	res := ""
	var result []string
	for _, char := range data {
		if char == ')' {
			res += string(char)
			result = append(result, tretement(res)+"\n")

		} else if char == '\n' {
			res = ""
		} else {
			res += string(char)
		}
	}
	if res != "" {
		result = append(result, tretement(res)+"\n")
	}
	resultt := ""
	for _, char := range result {
		resultt += char
	}
	return resultt
}

func tretement(balance string) string {
	return balance
}
