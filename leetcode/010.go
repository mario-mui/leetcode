package leetcode

func IsMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
		dp[i][0] = false
	}
	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			if j > 1 {
				dp[0][j] = dp[0][j-2]
			} else {
				dp[0][j] = true
			}
		}
	}
	for i := 1; i <= len(s); i++ {
		chartA := s[i-1]
		for j := 1; j <= len(p); j++ {
			chartB := p[j-1]
			if chartA == chartB || chartB == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if chartB == '*' {
					if j > 1 {
						if dp[i][j-2] {
							dp[i][j] = true
						} else {
							if chartA == p[j-2] || p[j-2] == '.' {
								dp[i][j] = dp[i-1][j]
							} else {
								dp[i][j] = false
							}
						}
					} else {
						dp[i][j] = true
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
