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
	NK := getIntInputFromOneLine()
	N, K := NK[0], NK[1]
	dp := [][][]int{}
	days := [101]int{}
	dayA := [101]bool{}
	for i := 0; i < K; i++ {
		t := getIntInputFromOneLine()
		dayA[t[0]] = true
		days[t[0]] = t[1] - 1
	}

	for i := 0; i <= N; i++ {
		t := [][]int{}
		for j := 0; j < 3; j++ {
			s := []int{}
			for k := 0; k < 2; k++ {
				s = append(s, 0)
			}
			t = append(t, s)
		}
		dp = append(dp, t)
	}
	if dayA[1] {
		dp[1][days[1]][0] = 1
	} else {
		for i := 0; i < 3; i++ {
			dp[1][i][0] = 1
		}
	}
	d := 10000
	for i := 1; i < N; i++ {
		for j := 0; j < 3; j++ {
			if dayA[i+1] {
				k := days[i+1]
				if j == k {
					dp[i+1][k][1] += dp[i][j][0]
				} else {
					dp[i+1][k][0] += dp[i][j][0] + dp[i][j][1]
				}
				dp[i+1][k][0] %= d
				dp[i+1][k][1] %= d
				continue
			}
			for k := 0; k < 3; k++ {
				if j == k {
					dp[i+1][k][1] += dp[i][j][0]
				} else {
					dp[i+1][k][0] += dp[i][j][0] + dp[i][j][1]
				}
				dp[i+1][k][0] %= d
				dp[i+1][k][1] %= d
			}
		}
	}
	var ans int
	for j := 0; j < 3; j++ {
		for k := 0; k < 2; k++ {
			ans += dp[N][j][k]
		}
		ans %= 10000
	}

	fmt.Println(ans)
}
