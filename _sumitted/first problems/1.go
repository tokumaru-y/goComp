package main

import "fmt"

func main() {
	var ab [2]int
	fmt.Scan(&ab[0], &ab[1])
	x := ab[0] * ab[1]
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
