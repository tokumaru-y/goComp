package main

import (
	"bufio"
	"fmt"
	"math"
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

func ThreePartSearch(l, r, P float64) float64 {
	f := func(x float64) float64 {
		return x + (P / math.Pow(2, x/1.5))
	}
	for i := 0; i < 200; i++ {
		c1 := (l*2 + r) / 3
		c2 := (l + r*2) / 3
		if f(c1) > f(c2) {
			l = c1
		} else {
			r = c2
		}
	}
	return f(l)
}

func main() {
	s.Scan()
	p := s.Text()
	t, _ := strconv.ParseFloat(p, 64)
	P := float64(t)

	l := 0.0
	r := float64(200)
	ans := ThreePartSearch(l, r, P)
	fmt.Println(ans)
}
