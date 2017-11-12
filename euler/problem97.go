package main

import "fmt"

const (
	tenMln int64 = 10000000000
)

type Problem97 struct{}

func (p *Problem97) GetSolution() string {
	last10Digits := (28433*getLastTenDigitsOfTwoAtPow(7830457) + 1) % tenMln
	return fmt.Sprint(last10Digits)
}

func getLastTenDigitsOfTwoAtPow(p int) int64 {
	result := int64(1)
	for ; p != 0; p-- {
		result = (result * 2) % tenMln
	}

	return result
}
