package discoverpin

import (
	"fmt"
	"strings"
)

// GetPin returns a string of the given pin
func GetPin(keys []string) string {

	answer := make([]string, 0)

	for _, key := range keys {
		if len(key) == 3 {
			// Append single key if not already in answer
			for i := 0; i < len(key); i++ {
				if indexOf(answer, string(key[i])) == -1 {
					answer = append(answer, string(key[i]))
				}
			}

			for j := 0; j < 3; j++ {
				// Check position of each new key
				iOfA := indexOf(answer, string(key[0]))
				iOfB := indexOf(answer, string(key[1]))
				iOfC := indexOf(answer, string(key[2]))

				if iOfA > iOfB {
					answer = insert(answer, iOfA, iOfB)
					continue
				}
				if iOfB > iOfC {
					answer = insert(answer, iOfB, iOfC)
					continue
				}
				if iOfA > iOfC {
					answer = insert(answer, iOfA, iOfC)
					continue
				}
			}
		}
		fmt.Println(strings.Join(answer, ""))
	}
	return strings.Join(answer, "")
}

func insert(a []string, upper int, lower int) []string {
	val := a[upper]
	var removed []string
	if upper+1 == len(a) {
		removed = a[:upper]
	} else {
		removed = append(a[:upper], a[upper+1:]...)
	}

	newLower := a[:lower]
	newUpper := append([]string{val}, removed[lower:]...)
	updatedA := append(newLower, newUpper...)
	return updatedA
}

func indexOf(s []string, v string) int {
	for i := range s {
		if s[i] == v {
			return i
		}
	}
	return -1
}
