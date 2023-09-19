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
func dfs(g [][]int, d *[110]int, f *[110]int, t *int, n int) {
	if d[n] != 0 {
		return
	}
	d[n] = *t
	for i := 0; i < len(g[n]); i++ {
		nn := g[n][i]
		if d[nn] != 0 {
			continue
		}
		*t += 1
		dfs(g, d, f, t, nn)
	}
	*t += 1
	f[n] = *t
}

func main() {
	N := getIntInput()
	g := [][]int{}
	for i := 0; i < N; i++ {
		g = append(g, make([]int, 0))
	}
	for i := 0; i < N; i++ {
		d := getIntInputFromOneLine()
		for j := 0; j < d[1]; j++ {
			g[i] = append(g[i], d[j+2]-1)
		}
	}

	for l := 0; l < N; l++ {
		sort.Slice(g[l], func(i, j int) bool { return g[l][i] < g[l][j] })
	}
	d := [110]int{}
	f := [110]int{}
	t := 0
	for i := 0; i < N; i++ {
		if d[i] != 0 {
			continue
		}
		t += 1
		dfs(g, &d, &f, &t, i)
	}

	for i := 0; i < N; i++ {
		fmt.Println(i+1, d[i], f[i])
	}
}
