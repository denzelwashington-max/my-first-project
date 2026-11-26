package main

import "fmt"

func RepeatAlpha(s string) string {
	var res []byte
	repeat := 1
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			repeat = int(ch - 'a')
		}
		if ch >= 'A' && ch <= 'Z' {
			repeat = int(ch - 'A')
		}

		for j := 0; j < repeat; j++ {
			res = append(res, ch)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
