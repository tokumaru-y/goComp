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

func main() {
	HW := getIntInputFromOneLine()
	H, W := HW[0], HW[1]
	sxy := getIntInputFromOneLine()
	exy := getIntInputFromOneLine()
	sx, sy := sxy[0]-1, sxy[1]-1
	ex, ey := exy[0]-1, exy[1]-1
	q := [][]int{}
	q = append(q, []int{sx, sy})
	dist := [][]int{}
	d := [][]string{}
	fv := 100000000
	for i := 0; i < H; i++ {
		s.Scan()
		c := s.Text()
		d = append(d, strings.Split(c, ""))
		t := []int{}
		for j := 0; j < W; j++ {
			t = append(t, fv)
		}
		dist = append(dist, t)
	}
	dist[sx][sy] = 0

	for len(q) > 0 {
		n := q[0]
		x, y := n[0], n[1]
		q = q[1:]
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == j || i+j == 0 {
					continue
				}
				nx, ny := x+i, y+j
				if !(0 <= nx && nx < H && 0 <= ny && ny < W) {
					continue
				}
				if dist[nx][ny] != fv || d[nx][ny] != "." {
					continue
				}
				dist[nx][ny] = dist[x][y] + 1
				q = append(q, []int{nx, ny})
			}
		}
	}

	fmt.Println(dist[ex][ey])
}
