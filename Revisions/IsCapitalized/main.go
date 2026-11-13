package main

import "fmt"

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == 0 || s[i-1] == ' ' {
			if ch >= 'a' && ch <= 'z' {
				return false
			}
		}
	}
	return true
}
func main(){
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello! How Are You"))
	fmt.Println(IsCapitalized("Denzel Washington"))
	fmt.Println(IsCapitalized("Whats 4this 100k?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!! Whatsthis4"))
	fmt.Println(IsCapitalized(""))

}