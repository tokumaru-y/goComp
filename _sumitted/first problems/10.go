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
	var x, y, t int
	for i := 0; i < N; i++ {
		TXY := getIntInputFromOneLine()
		T, X, Y := TXY[0], TXY[1], TXY[2]
		dis_x, dis_y, dif_t := getSubIntAbs(x, X), getSubIntAbs(y, Y), getSubIntAbs(t, T)
		if dif_t < dis_x+dis_y {
			fmt.Println("No")
			os.Exit(0)
		}
		if (dif_t-(dis_x+dis_y))%2 == 1 {
			fmt.Println("No")
			os.Exit(0)
		}
		x, y, t = X, Y, T
	}

	fmt.Println("Yes")
}
