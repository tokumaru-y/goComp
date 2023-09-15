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

func main() {
	candidate := []string{"dream", "dreamer", "erase", "eraser"}
	for i := 0; i < 4; i++ {
		candidate[i] = reveseString(candidate[i])
	}

	s.Scan()
	S := s.Text()
	S = reveseString(S)
	isAble := true
	for i := 0; i < len(S); {
		isOk := false
		for _, st := range candidate {
			if i+len(st) <= len(S) && S[i:i+len(st)] == st {
				isOk = true
				i += len(st)
			}
		}

		if !isOk {
			isAble = false
			break
		}
	}

	if isAble {
		fmt.Println("Yes")
	} else {
		fmt.Println("NO")
	}

}
