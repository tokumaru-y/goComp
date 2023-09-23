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

func makeBlankSlice(W int) []int {
	res := []int{}
	for i := 0; i < W+2; i++ {
		res = append(res, 0)
	}
	return res
}

func main() {
	HW := getIntInputFromOneLine()
	H, W := HW[1], HW[0]
	g := [][]int{}
	d := [][]int{}
	fg, fd := makeBlankSlice(W), makeBlankSlice(W)
	g = append(g, fg)
	d = append(d, fd)

	for i := 0; i < H; i++ {
		l := getIntInputFromOneLine()
		t := []int{0}
		t = append(t, l...)
		t = append(t, 0)
		g = append(g, t)

		t = makeBlankSlice(W)
		d = append(d, t)
	}

	fg, fd = makeBlankSlice(W), makeBlankSlice(W)
	g = append(g, fg)
	d = append(d, fd)

	var ans int
	// 建物の周囲の壁の数をカウント
	for h := 1; h <= H; h++ {
		for w := 1; w <= W; w++ {
			if g[h][w] == 0 {
				continue
			}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if h%2 == 1 {
						if (i == -1 && j == -1) || (i == 1 && j == -1) || (i == 0 && j == 0) {
							continue
						}
					} else {
						if (i == -1 && j == 1) || (i == 1 && j == 1) || (i == 0 && j == 0) {
							continue
						}
					}

					nh, nw := h+i, w+j
					if g[nh][nw] == 0 {
						d[h][w]++
					}
				}
			}
		}
	}
	// 0,0から初めて外壁部分の要素を塗りつぶす
	seen := [][]bool{}
	for i := 0; i < len(g); i++ {
		t := []bool{}
		for j := 0; j < len(g[0]); j++ {
			t = append(t, false)
		}
		seen = append(seen, t)
	}
	q := [][]int{}
	q = append(q, []int{0, 0})
	seen[0][0] = true
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		h, w := n[0], n[1]
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if h%2 == 1 {
					if (i == -1 && j == -1) || (i == 1 && j == -1) || (i == 0 && j == 0) {
						continue
					}
				} else {
					if (i == -1 && j == 1) || (i == 1 && j == 1) || (i == 0 && j == 0) {
						continue
					}
				}
				nh, nw := h+i, w+j
				if !((0 <= nh && nh < len(g)) && (0 <= nw && nw < len(g[0]))) {
					continue
				}
				if seen[nh][nw] || g[nh][nw] == 1 {
					continue
				}
				seen[nh][nw] = true
				q = append(q, []int{nh, nw})
			}
		}

	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			ans += d[i][j]
		}
	}
	// 建物に囲まれた空地を計算し、ansから引く
	for h := 0; h < len(g); h++ {
		for w := 0; w < len(g[0]); w++ {
			if seen[h][w] || g[h][w] == 1 {
				continue
			}
			cnt := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if h%2 == 1 {
						if (i == -1 && j == -1) || (i == 1 && j == -1) || (i == 0 && j == 0) {
							continue
						}
					} else {
						if (i == -1 && j == 1) || (i == 1 && j == 1) || (i == 0 && j == 0) {
							continue
						}
					}
					nh, nw := h+i, w+j
					if !((0 <= nh && nh < len(g)) && (0 <= nw && nw < len(g[0]))) {
						continue
					}
					if g[nh][nw] == 1 {
						cnt++
					}
				}
			}
			ans -= cnt
			seen[h][w] = true
		}
	}
	fmt.Println(ans)
}
