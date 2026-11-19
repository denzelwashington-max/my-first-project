package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print(" ")
		return
	}

	var result []byte
	s := os.Args[1]
	wordStarted := false
	spacePending := false

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == ' ' || ch == '\f' {
			if wordStarted {
				spacePending = true
			}
		} else {
			if spacePending {
				result = append(result, ' ')
				spacePending = false
			}
			result = append(result, ch)
			wordStarted = true
		}
	}
	fmt.Println(string(result))
}
