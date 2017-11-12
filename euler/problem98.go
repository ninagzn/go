package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strings"
)

var (
	digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

type Problem98 struct{}

func (p *Problem98) GetSolution() string {
	words := readWords("problem98.txt")
	maxSquare := int64(0)
	perms := make(map[int][][]int, 0)
	for key, v := range words {
		if len(v) == 1 {
			continue
		}

		distinctCharacters := getDistinctCharacters(key)
		lettersCount := len(distinctCharacters)
		p := getPermutations(lettersCount, &perms)

		for _, prm := range p {
			for i := 0; i < len(v); i++ {
				nr1 := createNumber(prm, distinctCharacters, v[i])
				if !isSquare(nr1) {
					continue
				}

				for j := i + 1; j < len(v); j++ {
					if len(v[i]) != len(v[j]) {
						continue
					}

					nr2 := createNumber(prm, distinctCharacters, v[j])
					if nr2 > 0 && int64(math.Sqrt(float64(nr2)))*int64(math.Sqrt(float64(nr2))) == nr2 {
						if nr2 > maxSquare {
							maxSquare = nr2
						}
						if nr1 > maxSquare {
							maxSquare = nr1
						}

					}
				}
			}
		}
	}
	return fmt.Sprint(maxSquare)
}

func getPermutations(length int, perms *map[int][][]int) [][]int {
	p, ok := (*perms)[length]
	if !ok {
		combs := getAllCombinations(length, digits, 0)
		m := make([][]int, 0)
		for _, cmb := range combs {
			p := generatePermutations(cmb)
			for _, i := range p {
				m = append(m, i)
			}
		}
		(*perms)[length] = m
	}

	return p
}

func isSquare(n int64) bool {
	return n > 0 && math.Pow(float64(int64(math.Sqrt(float64(n)))), 2) == float64(n)
}

func getDistinctCharacters(word string) string {
	letters := make(map[string]bool)
	s := ""
	for _, l := range strings.Split(word, "") {
		contains, _ := letters[l]
		if !contains {
			s += l
			letters[l] = true
		}
	}

	return s
}

func createNumber(cmb []int, k string, word string) int64 {
	letterMap := make(map[string]int)

	for i, l := range strings.Split(k, "") {
		letterMap[l] = cmb[i]
	}
	n := int64(0)
	v := strings.Split(word, "")
	if letterMap[v[0]] == 0 {
		return -1
	}

	for i := 0; i < len(word); i++ {
		n = n*10 + int64(letterMap[v[i]])
	}
	return n
}

func getAllCombinations(length int, elements []int, start int) [][]int {
	if start >= len(elements) {
		return [][]int{}
	}
	n := [][]int{}
	if length == 1 {
		for i := start; i < len(elements); i++ {
			n = append(n, []int{elements[i]})
		}
	} else {
		n = getAllCombinations(length-1, elements, start+1)
		for i := 0; i < len(n); i++ {
			n[i] = append(n[i], elements[start])
		}
		m := getAllCombinations(length, elements, start+1)
		n = append(n, m...)
	}
	return n
}

func generatePermutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {

					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(arr, len(arr))

	return res
}

func getWordKey(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func readWords(filePath string) map[string][]string {
	content, _ := ioutil.ReadFile(filePath)
	re := regexp.MustCompile("\"[a-zA-Z]+\"")
	lines := re.FindAllString(string(content), -1)
	words := make(map[string][]string)

	for _, word := range lines {
		w := strings.Replace(word, "\"", "", -1)
		key := getWordKey(w)
		words[key] = append(words[key], w)
	}

	return words
}
