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

func dfs(g [][]int, c, ans *int, h, w int) {
	g[h][w] = 0
	*c++
	if *ans < *c {
		*ans = *c
	}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == j || i+j == 0 {
				continue
			}
			nh := h + i
			nw := w + j
			if !((0 <= nh && nh < len(g)) && (0 <= nw && nw < len(g[0]))) {
				continue
			}
			if g[nh][nw] == 0 {
				continue
			}
			dfs(g, c, ans, nh, nw)
		}
	}
	g[h][w] = 1
	*c--
}

func main() {
	W := getIntInput()
	H := getIntInput()
	g := [][]int{}
	for i := 0; i < H; i++ {
		t := getIntInputFromOneLine()
		g = append(g, t)
	}
	var ans int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if g[i][j] == 1 {
				var c int
				dfs(g, &c, &ans, i, j)
			}
		}
	}

	fmt.Println(ans)
}
