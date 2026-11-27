package main

func LastWord(s string) string {
	if len(s) == 0 {
		return ""
	}

	last := len(s) - 1

	for last >= 0 && (s[last] == '\n' || s[last] == ' ' || s[last] == '\t' || s[last] == '\f') {
		last--
	}

	if len(s) < 0 {
		return "\n"
	}
	first := last

	for first >= 0 && !(s[first] == '\n' || s[first] == ' ' || s[first] == '\f' || s[first] == '\t') {
		first--
	}
	return s[first+1:last+1] + "\n"
}

/*
func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
*/
