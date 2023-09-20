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

func bfs(ans []int, g [][]int, n int) {
	q := make([]int, 110)
	q = append(q, n)
	ans[0] = 0
	for len(q) > 0 {
		nx := q[0]
		for _, i := range g[nx] {
			if ans[i] > -1 {
				continue
			}
			ans[i] = ans[nx] + 1
			q = append(q, i)
		}
		q = q[1:]
	}
}

func main() {
	N := getIntInput()
	g := [][]int{}
	for i := 0; i < N; i++ {
		t := getIntInputFromOneLine()
		if t[1] == 0 {
			g = append(g, []int{})
		} else {
			for j := 0; j < t[1]; j++ {
				t[2+j]--
			}
			g = append(g, t[2:])

		}
	}

	var ans []int
	for i := 0; i < N; i++ {
		ans = append(ans, -1)
	}
	bfs(ans, g, 0)
	for i := 0; i < N; i++ {
		fmt.Println(i+1, ans[i])
	}
}
