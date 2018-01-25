package main

type Trie struct {
	Value    byte
	Children []Trie
}

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	t := buildTrie(wordDict)

	return isSentence(s, 0, t, map[int]bool{})
}

func buildTrie(dict []string) Trie {
	t := Trie{
		' ',
		make([]Trie, 27)}
	for i := 0; i < len(dict); i++ {
		addWordToTrie(dict[i], &t)
	}

	return t
}

func addWordToTrie(w string, t *Trie) {
	node := *t

	for i := 0; i < len(w); i++ {
		index := int(w[i]) - 96
		if len(node.Children[index].Children) == 0 {
			node.Children[index] = Trie{w[i], make([]Trie, 27)}
		}
		node = node.Children[index]
	}

	node.Children[0] = Trie{'*', make([]Trie, 0)}
}

func isSentence(s string, startIndex int, t Trie, partialResults map[int]bool) bool {
	if startIndex == len(s) {
		return true
	}

	result, exists := partialResults[startIndex]
	if exists {
		return result
	}

	for i := startIndex; i < len(s); i++ {
		if isWord(s, startIndex, i, t) {
			result = isSentence(s, i+1, t, partialResults)
			partialResults[i+1] = result
			if result {
				return true
			}
		}
	}

	return false
}

func isWord(s string, startIndex, endIndex int, t Trie) bool {
	node := t
	for i := startIndex; i <= endIndex; i++ {
		index := int(s[i]) - 96
		if len(node.Children[index].Children) == 0 {
			return false
		}
		node = node.Children[index]
	}

	return node.Children[0].Value == '*'
}
