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
	NK := getIntInputFromOneLine()
	N, K := NK[0], NK[1]
	A := getIntInputFromOneLine()

	ans := math.MaxInt64
	for b := 0; b < (1 << N); b++ {
		var sum int
		var cnt int
		var mh int
		for i := 0; i < N; i++ {
			if (b & (1 << i)) > 0 {
				cnt++
				if mh >= A[i] {
					sum += mh - A[i] + 1
					mh += 1
				}
			}

			if mh < A[i] {
				mh = A[i]
			}
		}

		if cnt < K {
			continue
		}

		if ans > sum {
			ans = sum
		}
	}

	fmt.Println(ans)
}
