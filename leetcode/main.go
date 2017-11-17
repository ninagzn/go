package main

import "fmt"

func main() {
	sentence := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(sentence, wordDict))
}
