package main

func FirstWord(s string) string {
	if len(s) == 0 {
		return s
	}
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\n' || s[start] == '\f' || s[start] == '\t') {
		start++
	}
	if start == len(s) {
		return s
	}
	end := start

	for end < len(s) && !(s[end] == ' ' || s[end] == '\f' || s[end] == '\t' || s[end] == '\n') {
		end++
	}
	return s[start:end] + "\n"
}

/*
func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord("denzel washington"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("washington denzel"))
	fmt.Print(FirstWord("ugud?   .........  bye"))
}
*/
