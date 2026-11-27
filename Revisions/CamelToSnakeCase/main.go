package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	ls := s[len(s)-1]
	if ls >= 'A' && ls <= 'Z' {
		return s
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if i < len(s)-1 {
			nx := s[i+1]
			if ch >= 'A' && ch <= 'Z' && nx >= 'A' && nx <= 'Z' {
				return s
			}
		}
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			return s
		}
	}
	var res []byte
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			if i > 0 {
				res = append(res, '_')
			}
			res = append(res, ch+32)
		} else {
			res = append(res, ch)
		}

	}
	return string(res)
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
