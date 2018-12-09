package stringkatas

import (
	"fmt"
)

func LongestConsec(strarr []string, k int) string {

	var longestConsec string

	for i := 0; i < len(strarr); i++ {
		var curWord string
		for j := 0; j < k; j++ {
			fmt.Println(i)
			if i+j < len(strarr) && k < len(strarr) {
				curWord += strarr[i+j]
			}
		}

		fmt.Println(curWord)
		if len(longestConsec) < len(curWord) {
			longestConsec = curWord
		}
	}

	return longestConsec
}
