package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	var words []string
	str := os.Args[1]
	start := 0
	inWord := false

	for i := 0; i < len(str); i++ {
		ch := str[i]
		if ch == ' ' || ch == '\f' {
			if inWord {
				words = append(words, str[start:i])
				inWord = false
			}
		} else {
			if !inWord {
				start = i
				inWord = true
			}
		}
	}
	if inWord {
		words = append(words, str[start:])
		inWord = false
	}
	if len(words) == 0 {
		return
	}
	for i, word := range words {
		if i > 0 {
			fmt.Print("   ")
		}
		fmt.Print(word)
	}
	fmt.Println()
}
