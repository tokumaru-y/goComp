package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s = bufio.NewScanner(os.Stdin)

func convInt(s string) int {
	res, e := strconv.Atoi(s)
	if e != nil {
		panic("convInt failed")
	}
	return res
}

func getIntInput() int {
	s.Scan()
	st := s.Text()
	res := convInt(st)

	return res
}

func getIntInputFromOneLine() []int {
	s.Scan()
	st := strings.Split(s.Text(), " ")
	var res []int
	for _, s := range st {
		i := convInt(s)
		res = append(res, i)
	}

	return res
}

func main() {
	NY := getIntInputFromOneLine()
	N := NY[0]
	Y := NY[1]
	for i := 0; i <= N; i++ {
		for j := 0; j <= (N - i); j++ {
			k := N - i - j
			if 10000*i+5000*j+1000*k == Y {
				fmt.Println(i, j, k)
				os.Exit(0)
			}
		}
	}

	fmt.Println(-1, -1, -1)
}
