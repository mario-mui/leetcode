package leetcode

func MaxArea(height []int) int {
	len := len(height)
	i, j := 0, len-1
	max := min(height[i], height[j]) * (j - i)
	for i < j-1 { // 应该是可以算出下一次循环的面积，然后提前终止循环
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		area := min(height[i], height[j]) * (j - i)
		if max < area {
			max = area
		}
	}
	return max
}
