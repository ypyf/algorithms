package ch01

// 扩展欧几里得算法
// ax+bx=gcd(a,b)
// 输入a,b，输出x,y,gcd(a,b)
func ExtendedGcd(a, b int) (x, y, gcd int) {
	if b == 0 {
		return 1, 0, a
	}
	// b*x' + (a%b)*y' = ax + bx = gcd
	x, y, gcd = ExtendedGcd(b, a%b)
	return y, x - (a/b)*y, gcd
}
