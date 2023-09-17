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
		fmt.Println(s)
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
func bSeach(x []int, t int) bool {
	l := -1
	r := len(x)
	for r-l > 1 {
		m := (l + r) / 2
		if x[m] > t {
			r = m
		} else if x[m] < t {
			l = m
		} else if x[m] == t {
			return true
		}
	}
	return false
}

func main() {
	buf := make([]byte, 10000)
	s.Buffer(buf, 100000000000)
	getIntInput()
	S := getIntInputFromOneLine()
	getIntInput()
	T := getIntInputFromOneLine()
	var ans int
	for _, t := range T {
		if bSeach(S, t) {
			ans++
		}
	}

	fmt.Println(ans)
}
