package ch01

// 筛法求素数
// 输入一个自然数n
// 输出一个不大于n的素数列表
func sieve(n int) []int {
	p := 2
	var primes []int

	// 没有小于2的素数
	if n < p {
		return primes
	}

	// 构造自然数列表
	var arr []int
	for i := p; i <= n; i++ {
		arr = append(arr, i)
	}

	// 假定小于p*p的p的倍数都已经被筛去了
	// 也可以使用条件 p <= sqrt(n)
	for p*p <= n {
		// 筛去p的倍数
		if arr[p-2] != 0 {
			j := p * p
			for j <= n {
				arr[j-2] = 0
				j += p
			}
		}
		p++
	}

	for _, x := range arr {
		if x != 0 {
			primes = append(primes, x)
		}
	}

	return primes
}
