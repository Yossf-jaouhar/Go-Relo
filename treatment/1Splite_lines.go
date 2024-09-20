package treatment

// Split data line by line and return 2D array
func Split(data string) [][]string {
	var res string
	var newdata [][]string
	var myslice []string

	for _, char := range data {
		if char == '\n' {
			if len(res) > 0 {
				hh := Treatment(res)
				myslice = append(myslice, Treatment(hh))
				res = ""
			} else if len(res) == 0 {
				myslice = append(myslice, res)
			}
			if len(myslice) > 0 {

				newdata = append(newdata, myslice)
				myslice = nil
			}
		} else {
			res += string(char)
		}
	}

	if res != "" {
		hh := Treatment(res)
		myslice = append(myslice, Treatment(hh))
	}
	if len(myslice) > 0 {
		newdata = append(newdata, myslice)
	}

	return newdata
}
