package main

import (
	"fmt"
	"os"
	"strings"

	"goreloaded/treatment"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(">>Error in Arguments.<<")
		return
	} else if len(args) == 2 && !((strings.HasSuffix(args[0], ".txt")) && (strings.HasSuffix(args[1], ".txt"))) {
		fmt.Println(">>Wrong file name!!<<")
		return
	}
	Myfile := os.Args[1]
	data, err := os.ReadFile(Myfile)
	if err != nil {
		fmt.Println(err)
	}

	data2 := treatment.Split(string(data))

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

	err = os.WriteFile(args[1], []byte(linebyline), 0o666)
	if err != nil {
		fmt.Println(err)
	}
}
