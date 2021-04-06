package algorithm

// Greatest Common Divisor最大公约数，使用欧几里得算法可求
func Gcd(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return Gcd(q, r)
}
