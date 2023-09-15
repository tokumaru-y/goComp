package main

import (
	"bufio"
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

}
