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

func main() {
	ALL := getIntInputFromOneLine()
	A, B, C, X, Y := ALL[0], ALL[1], ALL[2], ALL[3], ALL[4]
	ans := math.MaxInt64
	var l int
	if X > Y {
		l = X * 2
	} else {
		l = Y * 2
	}

	for i := 0; i <= l; i += 2 {
		t := 0
		a := X - i/2
		b := Y - i/2

		if a > 0 {
			t += a * A
		}
		if b > 0 {
			t += b * B
		}
		t += i * C

		if ans > t {
			ans = t
		}
	}

	fmt.Println(ans)
}
