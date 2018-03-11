package ch01

// 选择排序算法
// 输入一个可排序数组
// 输出升序排列的数组
func SelectionSort(arr []int) []int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if sorted[j] < sorted[min] {
				min = j
			}
		}
		sorted[min], sorted[i] = sorted[i], sorted[min]
	}
	return sorted
}
