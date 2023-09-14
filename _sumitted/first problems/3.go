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

func main() {
	s.Scan()

	ans := math.MaxInt64
	s.Scan()
	for _, a := range strings.Split(s.Text(), " ") {
		num, _ := strconv.Atoi(a)
		tmp := 0
		for num%2 == 0 && num > 0 {
			tmp++
			num /= 2
		}
		if ans > tmp {
			ans = tmp
		}
	}

	if ans == math.MaxInt64 {
		ans = 0
	}
	fmt.Println(ans)
}
