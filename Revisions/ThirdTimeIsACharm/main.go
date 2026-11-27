package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return ""
	}
	var res []byte
	for i := 2; i < len(str); i += 3 {
		res = append(res, str[i])

	}
	res = append(res, '\n')
	return string(res)
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
