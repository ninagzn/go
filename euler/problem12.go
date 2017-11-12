package main

import (
	"math"
)

type Problem12 struct{}

func (p *Problem12) GetSolution() string {
	n := int64(500)
	sum := n * (n - 1) / 2
	var primes []int64
	addNextPrime(&primes)
	for i := n; true; i++ {
		sum += i
		if isTriangle(sum, &primes) {
			break
		}
	}
	return string(sum)
}

func addNextPrime(primes *[]int64) {
	if len(*primes) == 0 {
		*primes = append(*primes, 2)
		return
	}
	for i := (*primes)[len(*primes)-1]; true; i++ {
		if isPrime(i, *primes) {
			*primes = append(*primes, i)
			break
		}
	}
}

func isPrime(n int64, primes []int64) bool {
	for i := int64(0); i < int64(len(primes)); i++ {
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}

func isTriangle(n int64, primes *[]int64) bool {
	p := int64(1)
	divisorsNo := 1
	max := int64(math.Sqrt(float64(n)))
	for i := 0; (*primes)[i] <= max && divisorsNo < 500; i++ {
		p = (*primes)[i]
		if n%p == 0 {
			n0 := n
			k := 1
			for n0%p == 0 {
				k++
				n0 = n0 / p
			}
			divisorsNo *= k
		}
		if i == len(*primes)-1 {
			addNextPrime(primes)
		}
	}
	return divisorsNo >= 500
}
