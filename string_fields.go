package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	string_map := make(map[string]int)

	string_fields := strings.Fields(s)

	for i := 0; i < len(string_fields); i++ {
		string_map[string_fields[i]]++
	}
	return string_map
}

func main() {
	var phrase string = "It was the best of times it was the worst of times"
	fmt.Println(phrase)
	fmt.Println(WordCount(phrase))
}
