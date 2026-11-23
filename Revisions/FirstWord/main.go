package main

import (
	"fmt"
)

func FirstWord(s string) string {
	start := 0

	for start < len(s) && (s[start] == ' ' || s[start] == '\n' || s[start] == '\f' || s[start] == '\t') {
		start++
	}
	if start == len(s) {
		return "\n"
	}
	end := start
	for end < len(s) && !(s[end] == '\n' || s[end] == '\f' || s[end] == ' ' || s[end] == '\t') {
		end++
	}
	return s[start:end] + "\n"
}
func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord("denzel washington"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("washington denzel"))
	fmt.Print(FirstWord("hello   .........  bye"))
}
