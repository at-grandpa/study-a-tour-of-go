package main

import "golang.org/x/tour/wc"
import "strings"

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, char := range strings.Fields(s) {
		_, ok := result[char]
		if !ok {
			result[char] = 1
			continue
		}
		result[char]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
