package ch01

// 后续子集算法
// 输入一个代表一个含有n个元素的子集的0-1串
// 算法计算下一个子集所对应的0-1串
func NextSubset(subset string) string {
	n := len(subset)
	next := make([]rune, n)
	copy(next, []rune(subset))
	k := n - 1

	// 找到串最右边的0（如果存在），并把它变成1
	for k >= 0 && subset[k] == '1' {
		k--
	}

	// 把它右边的所有数字变成0
	if k >= 0 {
		next[k] = '1'
		for i := k + 1; i < n; i++ {
			next[i] = '0'
		}
	}

	return string(next)
}
