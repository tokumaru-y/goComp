package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

func main() {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	s.Scan()
	b, _ := strconv.Atoi(s.Text())
	s.Scan()
	c, _ := strconv.Atoi(s.Text())
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	var ans int
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for l := 0; l <= c; l++ {
				if 500*i+100*j+50*l == x {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
