package main

import (
//	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

var wordMap map[string]int

func WordCount(s string) map[string]int {

	wordMap = make(map[string]int)
	wordList := strings.Fields(s)

	for _, v := range wordList {
		wordMap[v] += 1
	}
	// return map[string]int{"x": 1}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
