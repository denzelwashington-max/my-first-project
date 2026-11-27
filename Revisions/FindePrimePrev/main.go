package main

import "fmt"

func FindPrevPrime(nb int) int {
	for n := nb; n >= 2; n-- {
		if n == 2 {
			return 2
		}
		if n%2 == 0 {
			continue
		}
		prime := true
		for d := 3; d*d <= n; d += 2 {
			if n%d == 0 {
				prime = true
			}
		}
		if prime {
			return n
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
