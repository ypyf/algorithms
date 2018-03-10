package ch01

// 计算Fib数列中的第n个数
// Fib(0) == 1
func Fibonacci(n uint) int {
	a := 1
	b := 1
	k := uint(1)
	for k < n {
		c := a + b
		a = b
		b = c
		k++
	}
	return b
}
