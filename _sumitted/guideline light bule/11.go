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

func getAns(N int, M int, sm [][]int, P []int) int {
	var res int
	for b := 0; b < (1 << N); b++ {
		isOk := true

		for i := 0; i < M; i++ {
			var t int
			for _, j := range sm[i] {
				if b&(1<<j) > 0 {
					t++
				}
			}

			if t%2 != P[i] {
				isOk = false
			}
		}

		if isOk {
			res++
		}
	}

	return res
}

func main() {
	NM := getIntInputFromOneLine()
	N := NM[0]
	M := NM[1]
	sm := [][]int{}
	for i := 0; i < M; i++ {
		l := getIntInputFromOneLine()
		for i := 0; i < l[0]; i++ {
			l[i+1]--
		}
		sm = append(sm, l[1:])
	}
	P := getIntInputFromOneLine()

	ans := getAns(N, M, sm, P)
	fmt.Println(ans)
}
