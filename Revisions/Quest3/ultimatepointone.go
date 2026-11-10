package main

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	var a int = 0
	
	var p1 *int = &a     
	var p2 **int = &p1  
	var p3 ***int = &p2  

	UltimatePointOne(p3)

	fmt.Println(a) 
}