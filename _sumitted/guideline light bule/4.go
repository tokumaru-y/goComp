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
	A := [][]int{}
	N, M := NM[0], NM[1]

	for i := 0; i < N; i++ {
		a := getIntInputFromOneLine()
		A = append(A, a)
	}

	var ans int
	for i := 0; i < M; i++ {
		for j := i + 1; j < M; j++ {
			var s int
			for k := 0; k < N; k++ {
				if A[k][i] > A[k][j] {
					s += A[k][i]
				} else {
					s += A[k][j]
				}
			}

			if ans < s {
				ans = s
			}
		}
	}

	fmt.Println(ans)
}
