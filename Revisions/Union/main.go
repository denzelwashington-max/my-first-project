package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	seen := make(map[rune]bool)

	var result []rune
	for _, r := range str1 {
		if !seen[r] {
			seen[r] = true
			result = append(result, r)
		}
	}

	for _, r := range str2 {
		if !seen[r] {
			seen[r] = true
			result = append(result, r)
		}
	}
	for _, r := range result {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}
