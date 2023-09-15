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

func reveseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func getSubIntAbs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func main() {
	for {
		NX := getIntInputFromOneLine()
		N, X := NX[0], NX[1]
		if N == 0 && X == 0 {
			break
		}

		var ans int
		for i := 1; i <= N; i++ {
			for j := i + 1; j <= N; j++ {
				k := X - (i + j)
				if i == j || j == k || k == i || k <= j {
					continue
				}
				if k >= 0 && k <= N {
					ans++
				}
			}
		}

		fmt.Println(ans)
	}

}
