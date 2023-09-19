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

func dfs(g [][]bool, h, w int) {
	lh := len(g)
	lw := len(g[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nh := h + i
			nw := w + j
			if !((0 <= nh && nh < lh) && (0 <= nw && nw < lw)) {
				continue
			}
			if !g[nh][nw] {
				continue
			}
			g[nh][nw] = false
			dfs(g, nh, nw)
		}
	}
}
func main() {
	for {
		wh := getIntInputFromOneLine()
		W := wh[0]
		H := wh[1]
		if H == 0 && W == 0 {
			return
		}
		grid := [][]bool{}
		for i := 0; i < H; i++ {
			t := make([]bool, W)
			grid = append(grid, t)
		}

		for i := 0; i < H; i++ {
			t := getIntInputFromOneLine()
			for j := 0; j < W; j++ {
				grid[i][j] = t[j] == 1
			}
		}

		var ans int
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if grid[i][j] {
					ans++
					dfs(grid, i, j)
				}
			}
		}
		fmt.Println(ans)
	}

}
