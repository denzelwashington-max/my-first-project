package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("options:abcdefghijklmnopqrstuvwxyz")
		return
	}

	var flats uint32 = 0
	for _, arg := range args {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		if arg[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
		for i := 1; i < len(arg); i++ {
			ch := arg[i]
			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			start := ch - 'a'
			flats |= (1 << start)
		}
	}

	for i := 31; i >= 0; i-- {
		bit := (flats >> uint(i)) & 1
		fmt.Printf("%d", bit)
		if i == 23 || i == 15 || i == 7 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
