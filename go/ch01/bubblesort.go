package ch01

// 冒泡排序算法
// 输入一个可排序数组
// 输出升序排列的数组
func BubbleSort(arr []int) []int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)

	for i := 0; i < n; i++ {
		for k := n - 2; k >= i; k-- {
			// 交换两个逆序元素
			if sorted[k+1] < sorted[k] {
				sorted[k+1], sorted[k] = sorted[k], sorted[k+1]
			}
		}
	}
	return sorted
}
