package stringkatas

import "fmt"

func HasUniqueChar(str string) bool {

	var charMap map[rune]int
	charMap = make(map[rune]int)

	for _, c := range str {
		val := charMap[c]
		if val == 0 {
			charMap[c] = 1
		} else {
			return false
		}
	}

	fmt.Println(charMap)

	return true
}
