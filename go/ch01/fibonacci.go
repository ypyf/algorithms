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

// 返回一个返回 int 的函数。
func Fibonacci2() func() int {
	a := 0
	b := 1
	f := func() int {
		r := a + b
		a = b
		b = r
		return a
	}
	return f
}
