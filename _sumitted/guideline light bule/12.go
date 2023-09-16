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
	NM := getIntInputFromOneLine()
	N := NM[0]
	M := NM[1]
	XY := [][]int{}
	for i := 0; i < M; i++ {
		xy := getIntInputFromOneLine()
		xy = []int{xy[0] - 1, xy[1] - 1}
		XY = append(XY, xy)
	}

	var ans int
	f := [][]bool{}
	t := []bool{}
	for i := 0; i < N; i++ {
		t = append(t, false)
	}
	for i := 0; i < N; i++ {
		tmp := append([]bool{}, t...)
		f = append(f, tmp)
	}

	for i := 0; i < M; i++ {
		x, y := XY[i][0], XY[i][1]
		f[x][y] = true
		f[y][x] = true
	}
	for b := 0; b < (1 << N); b++ {
		var t int
		l := []int{}
		for i := 0; i < N; i++ {
			if (b & (1 << i)) > 0 {
				l = append(l, i)
				t++
			}
		}

		isOk := true
		for i := 0; i < len(l); i++ {
			for j := 0; j < len(l); j++ {
				if i == j {
					continue
				}
				if !f[l[i]][l[j]] {
					isOk = false
				}
			}
		}

		if isOk && ans < t {
			ans = t
		}
	}

	fmt.Println(ans)
}
