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

func bSeach(x []int, t int) int {
	l := -1
	r := len(x)
	for r-l > 1 {
		m := (l + r) / 2
		if x[m] > t {
			r = m
		} else if x[m] < t {
			l = m
		} else if x[m] == t {
			return 0
		}
	}
	i, j := getSubIntAbs(t, x[r]), getSubIntAbs(t, x[l])
	if i > j {
		return j
	} else {
		return i
	}
}
func main() {
	D := getIntInput()
	N := getIntInput()
	M := getIntInput()
	shopsPositions := []int{}
	shopsPositions = append(shopsPositions, 0)
	shopsPositions = append(shopsPositions, D)
	for i := 0; i < N-1; i++ {
		s := getIntInput()
		shopsPositions = append(shopsPositions, s)
	}

	sort.Slice(shopsPositions, func(i, j int) bool { return shopsPositions[i] < shopsPositions[j] })

	var ans int
	for i := 0; i < M; i++ {
		m := getIntInput()
		ans += bSeach(shopsPositions, m)
	}
	fmt.Println(ans)
}
