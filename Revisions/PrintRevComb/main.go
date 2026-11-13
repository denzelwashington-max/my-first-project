package main

import "fmt"

func main() {
	first := true

	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				if !first {
					fmt.Print(",")
				}
				fmt.Printf("%d%d%d", i, j, k)
				first = false
			}

		}
	}
	fmt.Println()
}
