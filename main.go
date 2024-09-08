package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"goreloaded/functions"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data2 := Change(string(data))
	result := ""
	for _, char := range data2 {
		result += string(char)
	}
	err = os.WriteFile("result.txt", []byte(result), 0o666)
	if err != nil {
		log.Fatal(err)
	}
}

func Change(data string) []string {
	var result []string
	res := ""

	for _, char := range data {
		if char == ')' {
			if res != "" {
				res += string(char)

				if len(strings.Split(res, " ")) != 1{
					result = append(result, edit(res))
				}
				
				res = ""

			}
		} else {
			res += string(char)
		}
	}
	if res != "" {
		result = append(result, edit(res))
	}
	fmt.Println([]byte(strings.Join(result, "")))
	return result 
}

func edit(res string) string {
	data3 := strings.Split(res, " ")
	data4 := ""
	for i, char := range data3 {
		if char == "(cap)" {
			data4 += functions.Capitalize(data3[i-1])
			data4 = strings.Replace(data4, data3[i-1], "", -1)
		} else if char == "(low)" {
			data4 += strings.ToLower(data3[i-1])
			data4 = strings.Replace(data4, data3[i-1], "", -1)
		} else if char == "(up)" {
			data4 += strings.ToUpper(data3[i-1])
			data4 = strings.Replace(data4, data3[i-1], "", -1)
		} else {
			data4 += char + " "
		}
	}
	return data4
}
