package main

import (
	"os"

	"goreloaded/treatment"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}

	data2 := treatment.Div(string(data))

	linebyline := ""
	for i, row := range data2 {
		for _, cell := range row {
			if i != len(data2)-1 {
				linebyline += cell + "\n"
			} else {
				linebyline += cell
			}
		}
	}

	err = os.WriteFile("result.txt", []byte(linebyline), 0o666)
	if err != nil {
		panic(err)
	}
}
