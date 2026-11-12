package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	set1 := make(map[rune]bool)
	for _, r := range str1 {
		set1[r] = true
	}

	set2 := make(map[rune]bool)
	for _, r := range str2 {
		set2[r] = true
	}

	count := 0

	for r := range set1 {
		if !set2[r] {
			count++
		}
	}

	for r := range set2 {
		if !set1[r] {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(WeAreUnique("too", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
	fmt.Println(WeAreUnique("denzel", "washington"))
}
