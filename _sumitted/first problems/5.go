package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s = bufio.NewScanner(os.Stdin)

func isAcceptable(n, a, b int) bool {
	var res int
	for n > 0 {
		res += n % 10
		n /= 10
	}
	if a > res || res > b {
		return false
	}

	return true
}

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
	inputs := strings.Split(s.Text(), " ")
	inputInt := convIntFromStrings(inputs)
	var ans int
	for i := 1; i <= inputInt[0]; i++ {
		if isAcceptable(i, inputInt[1], inputInt[2]) {
			ans += i
		}
	}

	fmt.Println(ans)
}
