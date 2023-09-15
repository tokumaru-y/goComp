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
	s.Scan()
	S := []rune(s.Text())
	c := []rune{'A', 'C', 'G', 'T'}
	var ans int
	var tmp_cnt int
	for i := 0; i < len(S); i++ {
		isOk := false
		for j := 0; j < 4; j++ {
			if S[i] == c[j] {
				tmp_cnt++
				isOk = true
				break
			}
		}
		if !isOk {
			tmp_cnt = 0
		}
		if ans < tmp_cnt {
			ans = tmp_cnt
		}
	}

	fmt.Println(ans)
}
