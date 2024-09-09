package functions

func Atoi(s string) int {
	negative := false
	res := 0
	start := 0
	new := []rune(s)
	if len(s) == 0 {
		return 1
	}
	if new[0] == '-' {
		negative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	for i := start; i < len(new); i++ {
		if new[i] < '0' || new[i] > '9' {
			return 0
		}
		res = res*10 + int(new[i]-'0')
	}
	if negative {
		return -res
	}
	return res
}