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
	s.Scan()
	_S := s.Text()
	S := []rune(_S)

	var ans int
	for i := 0; i < 1000; i++ {
		t := fmt.Sprintf("%03d", i)
		rt := []rune(t)
		var cnt int
		for j := 0; j < N; j++ {
			if rt[cnt] == S[j] {
				cnt++
			}

			if cnt == 3 {
				ans++
				break

			}
		}
	}

	fmt.Println(ans)
}
