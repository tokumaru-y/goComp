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
	NW := getIntInputFromOneLine()
	N, W := NW[0], NW[1]
	dp := [10010]int{}
	g := [][]int{}
	for i := 0; i < N; i++ {
		vw := getIntInputFromOneLine()
		g = append(g, vw)
	}
	for i := 0; i < N; i++ {
		v, w := g[i][0], g[i][1]
		for j := 0; j <= 10000; j++ {
			nj := j - w
			if nj >= 0 && dp[j] < dp[nj]+v {
				dp[j] = dp[nj] + v
			}
		}
	}

	var ans int
	for i := 0; i <= W; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	fmt.Println(ans)
}
