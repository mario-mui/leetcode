package leetcode

func Convert(s string, rowNum int) string {
	if rowNum < 2 {
		return s
	}

	array := make([][]byte, rowNum)
	// down 表示方向，如果为 true 表示向下，如果为 false 表示向上
	index, down := 0, true

	for i := 0; i < len(s); i++ {
		array[index] = append(array[index], byte(s[i]))
		if index+1 == rowNum {
			down = false
		}
		if index == 0 {
			down = true
		}
		if down {
			index++
		} else {
			index--
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range array {
		solution = append(solution, row...)
	}
	return string(solution)
}
