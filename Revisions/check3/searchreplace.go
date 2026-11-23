package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Print()
		return
	}
	str := os.Args[1]
	old := []rune(os.Args[2])
	new := []rune(os.Args[3])

	if len(old) != 1 || len(new) != 1 {
		fmt.Println()
		return
	}
	oldch := old[0]
	newch := new[0]

	var result []rune
	for _, r := range str {
		if r == oldch {
			result = append(result, newch)
		} else {
			result = append(result, r)
		}
	}
	fmt.Println(string(result))
}
