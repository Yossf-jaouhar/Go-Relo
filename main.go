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
	linebyline := ""
	for _, row := range data2 {

		for _, cell := range row {
			linebyline += cell
		}
		linebyline += "\n"

	}
	err = os.WriteFile("result.txt", []byte(linebyline), 0o666)
	if err != nil {
		panic(err)
	}
}

func div(data string) [][]string {
	var res string
	var bigarray [][]string
	var myslice []string

	for _, char := range data {
		if char == ')' {
			res += string(char)

			myslice = append(myslice, res)
			res = ""
		} else if char == '\n' {
			if len(res) > 0 {
				myslice = append(myslice, res)
				res = ""
			} else if len(res) == 0 {

				myslice = append(myslice, res)
				res = ""
			}
			if len(myslice) > 0 {
				bigarray = append(bigarray, treatment(myslice))
				myslice = nil
			}
		} else {
			res += string(char)
		}
	}

	if res != "" {
		myslice = append(myslice, res)
	}

	if len(myslice) > 0 {
		bigarray = append(bigarray, myslice)
	}
	return bigarray
}

func treatment(myslice []string) []string {
	var mynewslice []string
	for _, char := range myslice {
		mynewslice = append(mynewslice, Change(char))
	}
	return mynewslice
}

func Change(char string) string {
	
}
