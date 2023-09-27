package main

import (
	"bufio"
	"fmt"
	"math"
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
	N := getIntInput()
	table := map[int]string{0: "R", 1: "B", 2: "W"}
	t_grid := [][]string{}
	for i := 0; i < 5; i++ {
		s.Scan()
		t := strings.Split(s.Text(), "")
		t_grid = append(t_grid, t)
	}

	dp := [][]int{}
	for i := 0; i <= N; i++ {
		t := []int{}
		for j := 0; j < 3; j++ {
			t = append(t, math.MaxInt)
		}
		dp = append(dp, t)
	}
	dp[0][0] = 0
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			js := table[j]
			c := 0
			for k := 0; k < 5; k++ {
				if t_grid[k][i] != js {
					c++
				}
			}
			for s := 0; s < 3; s++ {
				if j == s || dp[i][s] == math.MaxInt {
					continue
				}
				n := dp[i][s] + c
				if dp[i+1][j] > n {
					dp[i+1][j] = n
				}
			}
		}
	}
	ans := math.MaxInt
	for i := 0; i < 3; i++ {
		if ans > dp[N][i] {
			ans = dp[N][i]
		}
	}

	fmt.Println(ans)
}
