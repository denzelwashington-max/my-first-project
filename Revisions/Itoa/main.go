package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var result []byte
	for n > 0 {
		digit := n % 10

		result = append(result, byte('0'+digit))
		n /= 10
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	if isNegative {
		result = append( []byte{'-'},result...)
	}
	return string(result)
}
/*
func main() {
	fmt.Println(Itoa(1235))
	fmt.Println(Itoa(-89))
	fmt.Println(Itoa(786))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-9564))
}*/
