package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type Problem96 struct{}

func (p *Problem96) GetSolution() string {
	input := readInput("problem96.txt")
	sum := 0

	for i := 0; i < len(input); i++ {
		s := input[i]
		x := buildInitialMatrix(s)
		y, _ := completeSolution(x, 0, 0)
		sum = sum + y[0][0]*100 + y[0][1]*10 + y[0][2]

		fmt.Println()
		fmt.Println(i)
		printPairMatrix(x, y)
	}

	return fmt.Sprint(sum)
}

func printPairMatrix(x [9][9]int, y [9][9]int) {
	space := "      "
	hr := "-----------------------"
	verticalLine := " | "
	fmt.Println()
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println(hr+space, hr)
		}

		for j := 0; j < 9; j++ {
			if j > 0 && j%3 == 0 {
				fmt.Print(verticalLine)
			}
			fmt.Print(x[i][j], " ")
		}

		fmt.Print(space)

		for j := 0; j < 9; j++ {
			if j > 0 && j%3 == 0 {
				fmt.Print(verticalLine)
			}
			fmt.Print(y[i][j], " ")
		}
		fmt.Println()
	}
}

func isValidSudoku(x [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			m := getOptions(x, i, j)
			for o := 0; o < 9; o++ {
				if o+1 != x[i][j] && m[o] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func completeSolution(x [9][9]int, i, j int) ([9][9]int, bool) {
	if i >= 9 {
		return x, isValidSudoku(x)
	}
	if j >= 9 {
		return completeSolution(x, i+1, 0)
	}

	if x[i][j] != 0 {
		return completeSolution(x, i, j+1)
	}

	options := getOptions(x, i, j)
	k := -1
	for l := 0; l < 9; l++ {
		if options[l] != 0 {
			k = l
			break
		}
	}
	if k < 0 {
		return x, false
	}

	for l := k; l < 9; l++ {
		if options[l] != 0 {
			x[i][j] = l + 1
			s, valid := completeSolution(x, i, j+1)
			if valid {
				return s, true
			}
		}
	}

	return x, false
}

func readInput(filePath string) []string {
	content, _ := ioutil.ReadFile(filePath)
	re := regexp.MustCompile("Grid [0-9]{2}")
	lines := re.Split(string(content), -1)

	return lines[1:]
}

func getOptions(x [9][9]int, i, j int) [9]int {
	var options = [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for a := 0; a < 9; a++ {
		if x[a][j] != 0 {
			options[x[a][j]-1] = 0
		}
		if x[i][a] != 0 {
			options[x[i][a]-1] = 0
		}
	}

	for a := i - (i % 3); a < i-(i%3)+3; a++ {
		for b := j - (j % 3); b < j-(j%3)+3; b++ {
			if x[a][b] != 0 {
				options[x[a][b]-1] = 0
			}
		}
	}
	return options
}

func buildInitialMatrix(input string) [9][9]int {
	var x [9][9]int
	k := 0

	reg := regexp.MustCompile("[0-9]")
	s := reg.FindAllString(input, -1)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			d, _ := strconv.ParseInt(s[k], 0, 8)
			x[i][j] = int(d)
			k++
		}
	}

	return x
}
