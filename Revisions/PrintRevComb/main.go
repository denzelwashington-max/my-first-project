package main

import "fmt"

func main() {
	first := false

	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for h := j - 1; h >= 0; h-- {
				if !first {
					fmt.Print(",")
					fmt.Print(" ")
				}
				fmt.Printf("%d%d%d", i, j, h)
			}
		}
	}
	fmt.Println()
}
