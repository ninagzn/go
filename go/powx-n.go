package main

func myPow(x float64, n int) float64 {
	p := float64(1)
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			p = p * x
		}
		x *= x
	}
	return p
}
