package main

import "fmt"

const (
	tenMln int64 = 10000000000
)

func main() {
	fmt.Println(GetSolution97())
}

func GetSolution97() int64 {
	return (28433*getLastTenDigitsOfTwoAtPow(7830457) + 1) % tenMln
}

func getLastTenDigitsOfTwoAtPow(p int) int64 {
	result := int64(1)
	for ; p != 0; p-- {
		result = (result * 2) % tenMln
	}

	return result
}
