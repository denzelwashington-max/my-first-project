package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	step1 := make(map[rune]bool)
	for _, r := range str1 {
		step1[r] = true
	}

	step2 := make(map[rune]bool)
	for _, r := range str2 {
		step2[r] = true
	}

	count := 0

	for r := range step1 {
		if !step2[r] {
			count++
		}
	}
	for r := range step2 {
		if !step1[r] {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
