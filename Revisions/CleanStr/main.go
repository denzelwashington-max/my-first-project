package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
	}

	s := os.Args[1]
	wordStarted := false
	spacePending := false
	var res []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' {
			if wordStarted {
				spacePending = true
			}
		} else {
			if spacePending {
				res = append(res, ' ')
				spacePending = false

			}
			res = append(res, s[i])
			wordStarted = true
		}
	}
	fmt.Println(string(res))
}
