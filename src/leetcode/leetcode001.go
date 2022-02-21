package leetcode

func twoSum(list []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(list); i++ {
		t := target - list[i]
		if _, ok := m[t]; ok {
			return []int{m[t], i}
		}
		m[list[i]] = i
	}

	return nil
}
