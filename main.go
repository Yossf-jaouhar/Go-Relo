package main

import (
	"os"
	"strings"

	"goreloaded/functions"
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
	var res2 string
	var newdata [][]string
	var olddata [][]string
	var myslice []string
	var mystring []string
	// old data
	for _, char := range data {
		if char == '\n' {
			if len(res2) > 0 {
				mystring = append(mystring, res2)
				res2 = ""
			} else if len(res2) == 0 {
				mystring = append(mystring, res2)
			}
			if len(mystring) > 0 {

				olddata = append(olddata, mystring)
				mystring = nil
			}
		} else {
			res2 += string(char)
		}
	}
	if res2 != "" {
		mystring = append(mystring, res2)
	}
	if len(mystring) > 0 {
		olddata = append(olddata, div2(mystring))
	}

	// new data
	for _, char := range data {
		if char == '\n' {
			if len(res) > 0 {
				myslice = append(myslice, res)
				res = ""
			} else if len(res) == 0 {
				myslice = append(myslice, res)
			}
			if len(myslice) > 0 {

				newdata = append(newdata, div2(myslice))
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
		newdata = append(newdata, div2(myslice))
	}

	newdata = Comp(olddata, newdata)

	return newdata
}

func div2(myslice []string) []string {
	var res []string
	for _, cell := range myslice {
		res = append(res, treatment(cell))
	}
	return res
}

func treatment(cell string) string {
	mynewslice := strings.Fields(cell)
	var result []string

	for i, word := range mynewslice {
		if word == "(cap)" && len(result) > 0 {
			result[len(result)-1] = functions.Capitalize(result[len(result)-1])
		} else if word == "(cap)" && len(result) == 0 {
		} else if word == "(up)" && len(result) > 0 {
			result[len(result)-1] = functions.Upper(result[len(result)-1])
		} else if word == "(up)" && len(result) == 0 {
		} else if word == "(low)" && len(result) > 0 {
			result[len(result)-1] = functions.Lower(result[len(result)-1])
		} else if word == "(low)" && len(result) == 0 {
		} else if word == "(bin)" && len(result) > 0 {
			result[len(result)-1] = functions.Binary(result[len(result)-1], i)
		} else if word == "(bin)" && len(result) == 0 {
		} else {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

func Comp(old2d, new2d [][]string) [][]string {
	var res [][]string
	var temp []string
	for i := 0; i < len(old2d); i++ {
		if old2d[i][0] != "" && new2d[i][0] == "" {
			continue
		} else {
			temp = append(temp, new2d[i][0])
			res = append(res, temp)
			temp = nil
		}
	}
	return res
}
