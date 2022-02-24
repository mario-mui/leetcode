package leetcode

// 时间复杂度 O(n^2)，空间复杂度 O(1)
func LongestPalindrome(s string) string {
	subString := ""

	for i := 0; i < len(s); i++ {
		//回文是奇数
		subString = getPalindrome(s, i, i, subString)
		//回文是偶数
		subString = getPalindrome(s, i, i+1, subString)
	}
	return subString
}

func getPalindrome(s string, j, k int, sub string) string {
	subString := ""
	for j >= 0 && k < len(s) && s[j] == s[k] {
		subString = s[j : k+1]
		j--
		k++
	}
	if len(sub) < len(subString) {
		return subString
	}
	return sub
}

// 待学习 马拉车算法 将时间复杂度从 O(n^2) 优化到了 O(n)，空间换了时间，空间复杂度增加到 O(n)
