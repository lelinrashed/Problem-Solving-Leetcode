package main

import (
	"fmt"
)

func main() {
	var names []string = []string{"flower", "flow", "flight"}
	fmt.Println("prefix", longestCommonPrefix(names))
}

func longestCommonPrefix(strs []string) string {
	lenOfStrs := len(strs)

	if lenOfStrs == 0 {
		return ""
	}

	firstWord := strs[0]
	lenOfFirstWord := len(firstWord)
	match := true
	commonPrefix := ""

	for i := 0; i < lenOfFirstWord; i++ {
		firstStringChar := string(firstWord[i])
		for j := 1; j < lenOfStrs; j++ {
			if (len(strs[j]) - 1) < i {
				match = false
				break
			}

			if string(strs[j][i]) != firstStringChar {
				match = false
				break
			}
		}

		if match {
			commonPrefix += firstStringChar
		} else {
			break
		}
	}

	return commonPrefix
}
