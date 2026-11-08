package main

import "fmt"

func HashCode(dec string) string {
	size := len(dec)
	var result []byte

	for i := 0; i < len(dec); i++ {
		hash := (int(dec[i]) + size) % 127
		if hash <= 32 {
			hash += 33
		}
		result = append(result, byte(hash))
	}
	return string(result)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("DG"))
	fmt.Println(HashCode("BFT"))
	fmt.Println(HashCode("Done"))
	fmt.Println(HashCode("Denzel"))
}
