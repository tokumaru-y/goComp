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
	N, M := NM[0], NM[1]
	a := [][]int{}
	for i := 0; i < N; i++ {
		tmp := getIntInputFromOneLine()
		a = append(a, tmp)
	}

	var ans int

	for b := 0; b < (1 << N); b++ {
		ts := []int{}
		for i := 0; i < M; i++ {
			ts = append(ts, 0)
		}

		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if (b & (1 << i)) > 0 {
					ts[j] += (a[i][j] + 1) % 2
				} else {
					ts[j] += a[i][j]
				}
			}
		}
		var s int
		for i := 0; i < M; i++ {
			if ts[i] > N-ts[i] {
				s += ts[i]
			} else {
				s += N - ts[i]
			}
		}

		if s > ans {
			ans = s
		}

	}

	fmt.Println(ans)
}
