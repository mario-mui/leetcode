package leetcode

func LongestCommonPrefix(strs []string) string {
	s1 := strs[0]
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(strs); j++ {
			if len(strs[i]) <= j || strs[j][i] != s1[i] {
				return s1[0:i]
			}
		}
	}
	return s1
}
