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

func divisor(n int) int {
	var cnt int
	l := int(math.Sqrt(float64(n)))
	for i := 1; i <= l; i++ {
		if n%i == 0 {
			cnt++
			if n/i != i {
				cnt++
			}

			n /= i
		}
	}

	return cnt
}

func main() {
	N := getIntInput()

	var ans int
	for i := 1; i <= N; i += 2 {
		cnt := divisor(i)
		if cnt == 8 {
			ans++
		}
	}

	fmt.Println(ans)
}
