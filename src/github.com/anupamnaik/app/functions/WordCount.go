package functions

import (
	"strings"
)

// WordCount count the number of words in the given string
func WordCount(s string) map[string]int {
	wcount := make(map[string]int)
	for _, v := range strings.Fields(s) {
		wcount[v] = wcount[v] + 1
	}
	return wcount
}
