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

func isOk(m int, hs [][]int) bool {
	t := len(hs)
	h := make([]int, t)
	for i := 0; i < t; i++ {
		h[i] = hs[i][0]
		if h[i] > m {
			return false
		}
	}
	for i := 1; i <= t; i++ {
		var cnt int
		for j := 0; j < t; j++ {
			h[j] += hs[j][1]
			if h[j] > m {
				cnt++
			}
		}
		if cnt > i {
			return false
		}
	}
	return true
}

func main() {
	N := getIntInput()
	HS := [][]int{}
	for i := 0; i < N; i++ {
		hs := getIntInputFromOneLine()
		HS = append(HS, hs)
	}
	l := 0
	r := 1000000000 + 1
	for m := (l + r) / 2; r-l > 1; {
		if isOk(m, HS) {
			r = m
		} else {
			l = m
		}
		m = (l + r) / 2
	}

	fmt.Println(r)
}
