package leetcode

/**
解题思路
滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。
一旦出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。
以此类推，每次移动需要计算当前长度，并判断是否需要更新最大长度，最终最大的值就是题目中的所求
*/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int

	bottom, top, total := 0, -1, 0
	for bottom < len(s) {
		if top+1 == len(s) {
			// 如果 top 走到最后一个，基本表示最长的子字符串已经出现了
			break
		}
		if top+1 < len(s) && freq[s[top+1]] == 0 {
			freq[s[top+1]]++
			top++
		} else {
			freq[s[bottom]]--
			bottom++
		}
		total = max(total, top-bottom+1)
	}
	return total
}
