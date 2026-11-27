package main

import "fmt"

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i++ {
		if i%4 != 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%02x", arr[i])
		if (i+1)%4 == 0 || i == len(arr)-1 {
			fmt.Println()
		}
	}
	for _, b := range arr {
		if b >= 33 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
