package main

import "fmt"

type EulerProblem interface {
	GetSolution() string
}

func main() {
	p := Problem54{}

	fmt.Println(p.GetSolution())
}
