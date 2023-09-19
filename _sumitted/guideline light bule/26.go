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

func dfs(g [][]int, t *int, d *[200010]int, n int, p int) {
	o := d[n]
	*t += d[n]
	d[n] = *t
	for i := 0; i < len(g[n]); i++ {
		nn := g[n][i]
		if nn == p {
			continue
		}
		dfs(g, t, d, nn, n)
	}
	*t -= o
}

func main() {
	NQ := getIntInputFromOneLine()
	N, Q := NQ[0], NQ[1]
	g := [][]int{}
	for i := 0; i < N; i++ {
		g = append(g, make([]int, 0))
	}
	for i := 0; i < N-1; i++ {
		st := getIntInputFromOneLine()
		s, t := st[0]-1, st[1]-1
		g[s] = append(g[s], t)
		g[t] = append(g[t], s)
	}
	d := [200010]int{}
	for i := 0; i < Q; i++ {
		st := getIntInputFromOneLine()
		d[st[0]-1] += st[1]
	}

	t := 0
	dfs(g, &t, &d, 0, -1)

	for i := 0; i < N; i++ {
		if i == N-1 {
			fmt.Println(d[i])
		} else {
			fmt.Printf("%v ", d[i])
		}
	}
}
