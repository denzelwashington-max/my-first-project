package main

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid"
	}
	start := 1

	if from > to {
		start = -1
	}
	var res []byte
	for i := from; ; i += start {
		if i != from {
			res = append(res, ',', ' ')
		}
		res = append(res, byte('0'+i/10), byte('0'+i%10))
		if i == to {
			break
		}
	}
	res = append(res, '\n')
	return string(res)
}
func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
