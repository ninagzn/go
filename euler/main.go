package main

import "fmt"

type EulerProblem interface {
	GetSolution() string
}

func main() {
	p := Problem98{}

	fmt.Println(p.GetSolution())
}
