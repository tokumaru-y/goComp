package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	a := s.Text()
	n := 0
	for i := 0; i < 3; i++ {
		if a[i] == '1' {
			n++
		}
	}

	fmt.Println(n)

}
