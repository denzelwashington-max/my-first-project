package main

import "fmt"

func IntToByte(n int) []byte {
	if n == 0 {
		return []byte{'0'}
	}
	var res []byte
	for n > 0 {
		d := n % 10

		res = append([]byte{byte('0' + d)}, res...)
		n /= 10
	}
	return res
}

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	var result []byte
	n := len(s)
	i := 0

	for i < n {
		ch := s[i]
		count := 1

		for i+1 < n && s[i+1] == ch {
			count++
			i++
		}
		countByte := IntToByte(count)
		result = append(result, countByte...)

		result = append(result, ch)
		i++
	}
	return string(result)
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
	fmt.Println(ZipString("Helloo There wwaassssshiiiinnnnggggtttoooonnnn"))
	fmt.Println(ZipString("hiiiii thheeerrreee smmmaaarrrttt dddeeeennnnzzzeeelll"))
}
