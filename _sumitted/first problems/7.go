package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

func convInt(st string) int {
	res, e := strconv.Atoi(st)
	if e != nil {
		panic("fail")
	}
	return res
}

func convIntFromStrings(st []string) []int {
	var res []int
	for _, s := range st {
		res = append(res, convInt(s))
	}

	return res
}

func main() {
	s.Scan()
	N := convInt(s.Text())
	mochi := map[int]bool{}
	for i := 0; i < N; i++ {
		s.Scan()
		mochi[convInt(s.Text())] = true
	}

	fmt.Println(len(mochi))
}
