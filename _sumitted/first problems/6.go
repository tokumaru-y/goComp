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
	s.Scan()
	A := convIntFromStrings(strings.Split(s.Text(), " "))
	sort.Slice(A, func(i, j int) bool { return A[i] > A[j] })

	var alice, bob int
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			alice += A[i]
		} else {
			bob += A[i]
		}
	}

	fmt.Println(alice - bob)
}
