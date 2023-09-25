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
	Q := getIntInput()
	for ; Q > 0; Q-- {
		s.Scan()
		S := s.Text()
		s.Scan()
		T := s.Text()
		ls := len(S)
		lt := len(T)
		dp := [][]int{}
		for i := 0; i <= ls; i++ {
			tmp := []int{}
			for j := 0; j <= lt; j++ {
				tmp = append(tmp, 0)
			}
			dp = append(dp, tmp)
		}

		for i := 0; i < ls; i++ {
			for j := 0; j < lt; j++ {
				if S[i] == T[j] {
					dp[i+1][j+1] = dp[i][j] + 1
				}
				if dp[i+1][j+1] < dp[i][j+1] {
					dp[i+1][j+1] = dp[i][j+1]
				}
				if dp[i+1][j+1] < dp[i+1][j] {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}

		fmt.Println(dp[ls][lt])
	}
}
