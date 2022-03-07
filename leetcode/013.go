package leetcode

func RomanToInt(s string) int {
	m := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	res := m[string(s[len(s)-1])]

	for i := len(s) - 2; i >= 0; i-- {
		cur := m[string(s[i])]
		pre := m[string(s[i+1])]

		if pre > cur {
			res -= cur
		} else {
			res += cur
		}
	}

	return res
}
