package main

import "fmt"

func ThirdTimeIsACharm(s string) string {
	if len(s) < 3 {
		return "\n"
	}

	var result []byte
	for i := 2; i < len(s); i += 3 {
		result = append(result, s[i])
	}
	result = append(result, '\n')

	return string(result)
}
func main() {
	fmt.Println(ThirdTimeIsACharm("123456789"))
	fmt.Println(ThirdTimeIsACharm(""))
	fmt.Println(ThirdTimeIsACharm("a b c d e f"))
	fmt.Println(ThirdTimeIsACharm("12"))
	fmt.Println(ThirdTimeIsACharm("denzel"))
	fmt.Println(ThirdTimeIsACharm("w a s h i n g t o n"))
}
