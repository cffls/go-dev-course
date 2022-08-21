package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	start := 0

	for i, v := range s {
		if v == ' ' {
			m[s[start:i]]++
			start = i + 1
		}
	}

	m[s[start:]]++

	return m
}

func main() {
	wc.Test(WordCount)
}
