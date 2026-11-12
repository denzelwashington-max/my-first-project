package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}

	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}
	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	fmt.Println(chunks)
}

func main() {
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 7)
	Chunk([]int{0, 1, 2, 3, 47, 8, 9}, 6)
	Chunk([]int{0, 1, 6, 7, 8, 9}, 4)
	Chunk([]int{0, 7, 8, 9}, 5)
	Chunk([]int{}, 7)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 8)
}
