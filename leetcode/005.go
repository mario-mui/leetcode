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
// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0005.Longest-Palindromic-Substring/
func LongestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	newS := make([]rune, 0)
	newS = append(newS, '#')
	for _, c := range s {
		newS = append(newS, c)
		newS = append(newS, '#')
	}
	// dp[i]:    以预处理字符串下标 i 为中心的回文半径(奇数长度时不包括中心)
	// maxRight: 通过中心扩散的方式能够扩散的最右边的下标
	// center:   与 maxRight 对应的中心字符的下标
	// maxLen:   记录最长回文串的半径
	// begin:    记录最长回文串在起始串 s 中的起始下标
	dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
	for i := 0; i < len(newS); i++ {
		if i < maxRight {
			// 这一行代码是 Manacher 算法的关键所在
			dp[i] = min(maxRight-i, dp[2*center-i])
		}
		// 中心扩散法更新 dp[i]
		left, right := i-(1+dp[i]), i+(1+dp[i])
		for left >= 0 && right < len(newS) && newS[left] == newS[right] {
			dp[i]++
			left--
			right++
		}
		// 更新 maxRight，它是遍历过的 i 的 i + dp[i] 的最大者
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}
		// 记录最长回文子串的长度和相应它在原始字符串中的起点
		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - maxLen) / 2 // 这里要除以 2 因为有我们插入的辅助字符 #
		}
	}
	return s[begin : begin+maxLen]
}
