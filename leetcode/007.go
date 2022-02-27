package leetcode

func Reverse(number int) int {

	result := 0

	for number != 0 {
		result = result*10 + number%10
		number = number / 10
	}

	if result > 1<<30 || result < -(1<<31) {
		return 0
	}

	return result
}
