package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func absSubInt(a, b int) int {
	if b > a {
		return b - a
	} else {
		return a - b
	}
}

func main() {
	N := getIntInput()
	A := []int{}
	B := []int{}

	for i := 0; i < N; i++ {
		ab := getIntInputFromOneLine()
		A = append(A, ab[0])
		B = append(B, ab[1])
	}

	sort.Slice(A, func(i, j int) bool { return A[i] > A[j] })
	sort.Slice(B, func(i, j int) bool { return B[i] > B[j] })

	var ans int
	var ca int
	var cb int
	if N%2 == 0 {
		ca = (A[N/2] + A[N/2-1]) / 2
		cb = (B[N/2] + B[N/2-1]) / 2
	} else {
		ca = A[N/2]
		cb = B[N/2]
	}

	for i := 0; i < N; i++ {
		ans += absSubInt(ca, A[i]) + absSubInt(A[i], B[i]) + absSubInt(B[i], cb)
	}

	fmt.Println(ans)
}
