package main

import "strings"

func wordsTyping(sentence []string, rows, cols int) int {
	joinedSentence := strings.Join(sentence, " ")
	sentenceNo, rowsUsed := computeParagraph(joinedSentence, rows, cols)

	completeParagraphSentences := (rows / rowsUsed) * sentenceNo
	var remainingSentences int
	remainingSentences, rowsUsed = computeParagraph(joinedSentence, rows%rowsUsed, cols)

	return completeParagraphSentences + remainingSentences
}

func computeParagraph(sentence string, maxRows, cols int) (sentenceNo, rowsUsed int) {
	cells := 0
	length := len(sentence)
	for i := 1; i <= maxRows; i++ {
		cells += cols
		remainingCells := cells % (length + 1)
		if remainingCells == 0 || remainingCells == length {
			rowsUsed = i
			sentenceNo = cells / (length + 1)
			if remainingCells == length {
				sentenceNo++
			}
			return
		}
		if sentence[remainingCells] == ' ' {
			cells++
		} else {
			for i := remainingCells - 1; i >= 0; i-- {
				if sentence[i] == ' ' {
					break
				}
				cells--
			}
		}
	}

	rowsUsed = maxRows
	sentenceNo = cells / (length + 1)
	return
}
