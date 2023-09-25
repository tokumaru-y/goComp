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
	N := getIntInput()
	A := getIntInputFromOneLine()
	dp := [][]int{}
	for i := 0; i < N; i++ {
		tmp := []int{}
		for j := 0; j <= 20; j++ {
			tmp = append(tmp, 0)

		}
		dp = append(dp, tmp)
	}
	dp[1][A[0]] = 1
	for i := 1; i < N-1; i++ {
		for j := 0; j <= 20; j++ {
			if dp[i][j] > 0 {
				if j-A[i] >= 0 && j-A[i] <= 20 {
					dp[i+1][j-A[i]] += dp[i][j]
				}
				if j+A[i] >= 0 && j+A[i] <= 20 {
					dp[i+1][j+A[i]] += dp[i][j]
				}
			}
		}
	}

	fmt.Println(dp[N-1][A[N-1]])
}
