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
	N := NM[0]
	C := getIntInputFromOneLine()
	dp := [50010]int{}
	for i := 0; i < 50010; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	for _, c := range C {
		for i := 0; i <= 50000; i++ {
			ni := i - c
			if ni >= 0 && dp[ni] != math.MaxInt64 && dp[i] > dp[ni]+1 {
				dp[i] = dp[ni] + 1
			}
		}
	}
	for i := N; i >= 0; i-- {
		if dp[i] != math.MaxInt64 {
			ans := dp[i] + (N - i)
			fmt.Println(ans)
			break
		}
	}

}
