package stringkatas

// InPlaceReverseWords does an inplace reverse of the words delimted by a space " " in a run array
func InPlaceReverseWords(sentence []rune) []rune {

	reverseWords(sentence, 0, len(sentence)-1)

	var previousI int
	for i := 0; i < len(sentence); i++ {
		if string(sentence[i]) == " " {
			reverseWords(sentence, previousI, i-1)
			previousI = i + 1
		} else if i == len(sentence)-1 {
			reverseWords(sentence, previousI, i)
		}
	}

	return sentence
}

func reverseWords(unordered []rune, i int, j int) {
	for i < j {
		valOfI := unordered[i]
		unordered[i] = unordered[j]
		unordered[j] = valOfI
		i++
		j--
	}
}
