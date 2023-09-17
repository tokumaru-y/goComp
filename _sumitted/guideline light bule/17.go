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

type Grid [8][8]bool

func (g Grid) checkGrid() bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if g[i][j] {
				for h := -1; h+i >= 0; h-- {
					if g[i+h][j] {
						return false
					}
					if j+h >= 0 {
						if g[i+h][j+h] {
							return false
						}
					}
					if j-h < 8 {
						if g[i+h][j-h] {
							return false
						}
					}
				}
				for h := 1; h+i < 8; h++ {
					if g[i+h][j] {
						return false
					}
					if j+h < 8 {
						if g[i+h][j+h] {
							return false
						}
					}
					if j-h >= 0 {
						if g[i+h][j-h] {
							return false
						}
					}
				}

				for w := -1; w+j >= 0; w-- {
					if g[i][w+j] {
						return false
					}
				}
				for w := 1; w+j < 8; w++ {
					if g[i][w+j] {
						return false
					}
				}
			}
		}
	}
	return true
}

func dfs(g *Grid, h int, q map[int]int) bool {
	if h == 8 {
		return true
	}

	if _, e := q[h]; e {
		return dfs(g, h+1, q)
	} else {
		for w := 0; w < 8; w++ {
			g[h][w] = true
			f := g.checkGrid()
			if f {
				if dfs(g, h+1, q) {
					return true
				}
			}
			g[h][w] = false
		}
	}
	return false
}

func makeGrid(q map[int]int) Grid {
	var g Grid
	for k, v := range q {
		g[k][v] = true
	}
	dfs(&g, 0, q)

	return g
}

func main() {
	N := getIntInput()
	q := map[int]int{}
	for i := 0; i < N; i++ {
		xy := getIntInputFromOneLine()
		q[xy[0]] = xy[1]
	}
	g := makeGrid(q)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if g[i][j] {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
