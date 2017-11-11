package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strings"
)

func main() {

	quotedWords := readWords("problem98.in")
	words, _ := getWords(quotedWords)
	maxSquare := int64(0)

	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	maxWord := ""
	for key, v := range words {

		if len(v) == 1 {
			continue
		}
		digitMaps := getAllCombinations(len(key), digits, 0)
		//fmt.Println(digitMaps)
		for _, cmb := range digitMaps {
			p := getPermutations(cmb)
			//fmt.Println(p)
			for _, prm := range p {
				for i := 0; i < len(v); i++ {
					nr := getNumber(prm, key, v[i])
					if nr < 0 || int64(math.Sqrt(float64(nr)))*int64(math.Sqrt(float64(nr))) != nr {
						continue
					}

					for j := i + 1; j < len(v); j++ {
						if len(v[i]) != len(v[j]) {
							continue
						}

						nr2 := getNumber(prm, key, v[j])

						if nr2 != 2 && int64(math.Sqrt(float64(nr2)))*int64(math.Sqrt(float64(nr2))) == nr2 {
							fmt.Println(nr, nr2)
							if nr2 > maxSquare {
								fmt.Println(v[i], v[j], nr, nr2)
								maxWord = v[i] + " " + v[j]
								maxSquare = nr2
							}
							if nr > maxSquare {
								fmt.Println(v[i], v[j], nr, nr2)
								maxWord = v[i] + " " + v[j]

								maxSquare = nr
							}
							fmt.Println(key, cmb, prm, v[i], v[j], nr, nr2)

						}
					}

				}
			}
		}
	}
	fmt.Println(maxSquare, maxWord)

}

func getNumber(cmb []int, k string, word string) int64 {
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

// func getPermutations(digits []int) [][]int {
// 	if len(digits) == 1 {
// 		return [][]int{{digits[0]}}
// 	}
// 	t := [][]int{}
// 	p := getPermutations(digits[1:])

// 	for i := 0; i < len(p); i++ {
// 		k := make([]int, len(p[i]))
// 		copy(k, p[i])
// 		t = append(t, append([]int{digits[0]}, k...))

// 		for j := 0; j < len(p[i]); j++ {
// 			m := make([]int, len(p[i]))
// 			copy(m, p[i])
// 			m = m[:j+1]
// 			a := append(m, digits[0])
// 			t = append(t, append(a, m[j+1:]...))
// 		}
// 	}

// 	return t
// }

func getPermutations(arr []int) [][]int {
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

func getWords(quotedWords []string) (map[string][]string, int) {
	maxWordLength := 1

	words := make(map[string][]string)

	for _, word := range quotedWords {

		if maxWordLength < len(word) {
			maxWordLength = len(word)
		}
		w := strings.Replace(word, "\"", "", -1)
		key := getWordKey(w)
		words[key] = append(words[key], w)
	}

	return words, maxWordLength
}

func getWordKey(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func readWords(filePath string) []string {
	content, _ := ioutil.ReadFile(filePath)
	re := regexp.MustCompile("\"[a-zA-Z]+\"")
	lines := re.FindAllString(string(content), -1)

	return lines
}
