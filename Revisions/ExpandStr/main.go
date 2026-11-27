package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	inWord := false
	printedWord := false

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == ' ' || ch == '\t' {
			inWord = false
			continue
		}
		if !inWord {
			if printedWord {
				fmt.Print("   ")
			}
			inWord = true
			printedWord = true
		}
		fmt.Printf("%c", ch)
		printedWord = true
	}
	fmt.Println()
}
