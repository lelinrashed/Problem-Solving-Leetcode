package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"flow", "flower", "flight"}
	longestCommonPrefix(names)
}

func longestCommonPrefix(strs []string) string {
	var longestPrefix string = ""
	var endPrefix = false

	if len(strs) > 0 {
		sort.Strings(strs)
		fmt.Println("strs", strs)
		firstWord := string(strs[0])
		lastWord := string(strs[len(strs)-1])

		for i := 0; i < len(firstWord); i++ {
			if !endPrefix && string(lastWord[i]) == string(firstWord[i]) {
				longestPrefix += string(lastWord[i])
			} else {
				endPrefix = true
			}
		}
	}

	return longestPrefix
}
