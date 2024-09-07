package functions

func isPrint(r rune) bool {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}
func Capitalize(s string) string {
	mySlice := []rune(s)
	next := true
	for i := 0; i < len(s); i++ {
		if isPrint(mySlice[i]) && next {
			if mySlice[i] >= 'a' && mySlice[i] <= 'z' {
				mySlice[i] = mySlice[i] - ('a' - 'A')
			}
			next = false
		} else if mySlice[i] >= 'A' && mySlice[i] <= 'Z' {
			mySlice[i] = mySlice[i] + ('a' - 'A')
		} else if !isPrint(mySlice[i]) {
			next = true
		}
	}
	return string(mySlice)
}