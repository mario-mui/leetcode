package leetcode

func PalindromeNumner(num int) bool {
	if num < 0 {
		return false
	}
	if num == 0 {
		return true
	}
	if num%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)

	for num != 0 {
		arr = append(arr, num%10)
		num = num / 10
	}

	j := len(arr) - 1

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[j] {
			return false
		}
		j--
	}
	return true
}
