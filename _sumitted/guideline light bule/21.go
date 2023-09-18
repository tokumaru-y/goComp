package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	timeTable := make([]int, t)
	for i := 0; i < t; i++ {
		h, s := hs[i][0], hs[i][1]
		if h > m {
			return false
		}
		timeTable[i] = (m - h) / s
	}

	sort.Slice(timeTable, func(i, j int) bool { return timeTable[i] < timeTable[j] })
	for i := 0; i < t; i++ {
		if timeTable[i] < i {
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
	r := 1000000000000000000 + 1
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
