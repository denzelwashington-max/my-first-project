package main

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

		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			return s
		}
		if i+1 < len(s) {
			nx := s[i+1]
			if ch >= 'A' && ch <= 'Z' && nx >= 'A' && nx <= 'Z' {
				return s
			}
		}
	}
	var result []byte
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, (ch + 32))
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}

/*
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
*/
