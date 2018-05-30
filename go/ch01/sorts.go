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

// 插入排序
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		j := i - 1
		// 将大于待插入元素的数据移动到右侧
		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
		// 插入元素
		arr[j+1] = v
	}
	return arr
}
