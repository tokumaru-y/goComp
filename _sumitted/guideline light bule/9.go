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

type Point struct {
	x int
	y int
}

func main() {
	M := getIntInput()
	MAB := []Point{}
	for i := 0; i < M; i++ {
		a := getIntInputFromOneLine()
		MAB = append(MAB, Point{a[0], a[1]})
	}
	N := getIntInput()
	NAB := []Point{}
	Nm := map[Point]bool{}
	for i := 0; i < N; i++ {
		a := getIntInputFromOneLine()
		NAB = append(NAB, Point{a[0], a[1]})
		Nm[Point{a[0], a[1]}] = true
	}

	for k := 0; k < M; k++ {
		m := MAB[k]
		for i := 0; i < N; i++ {
			dx := NAB[i].x - m.x
			dy := NAB[i].y - m.y
			isOk := true
			for j := 0; j < M; j++ {
				if j == k {
					continue
				}
				p := MAB[j]
				a := dx + p.x
				b := dy + p.y
				_, ok := Nm[Point{a, b}]
				if !ok {
					isOk = false
					break
				}
			}

			if isOk {
				fmt.Println(dx, dy)
				os.Exit(0)
			}
		}
	}
}
