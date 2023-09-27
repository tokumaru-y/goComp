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
	NM := getIntInputFromOneLine()
	N, M := NM[0], NM[1]
	dist := []int{}
	cost := []int{}
	for i := 0; i < N; i++ {
		t := getIntInput()
		dist = append(dist, t)
	}
	for i := 0; i < M; i++ {
		t := getIntInput()
		cost = append(cost, t)
	}

	dp := [1010][1010]int{}
	for i := 0; i < 1010; i++ {
		for j := 0; j < 1010; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if dp[i][j] == math.MaxInt {
				continue
			}
			d := dist[i]
			c := cost[j]
			n := d * c
			if dp[i+1][j+1] > n+dp[i][j] {
				dp[i+1][j+1] = n + dp[i][j]
			}
			if dp[i][j+1] > dp[i][j] {
				dp[i][j+1] = dp[i][j]
			}
		}
	}

	ans := math.MaxInt
	for i := 1; i <= M; i++ {
		if ans > dp[N][i] {
			ans = dp[N][i]
		}
	}

	fmt.Println(ans)
}
