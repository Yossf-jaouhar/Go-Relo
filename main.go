package main

import (
	"goreloaded/project"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data2 := project.CleanStr(string(data))

	err = os.WriteFile("result.txt", []byte(data2), 0o666)
	if err != nil {
		log.Fatal(err)
	}
}
