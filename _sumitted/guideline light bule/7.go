package main

import (
	"bufio"
	"errors"
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

func calcPoint(A, B Point, pm map[Point]bool) (int, error) {
	x := A.y - B.y
	y := A.x - B.x
	C := Point{A.x + x, A.y - y}
	D := Point{B.x + x, B.y - y}
	_, cok := pm[C]
	_, dok := pm[D]
	if cok && dok {
		return (A.x-B.x)*(A.x-B.x) + (A.y-B.y)*(A.y-B.y), nil
	} else {
		return 0, errors.New("not exist")
	}
}

func main() {
	N := getIntInput()
	AB := []Point{}
	pointMap := map[Point]bool{}

	for i := 0; i < N; i++ {
		ab := getIntInputFromOneLine()
		p := Point{ab[0], ab[1]}
		AB = append(AB, p)
		pointMap[p] = true
	}

	var ans int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			a, e := calcPoint(AB[i], AB[j], pointMap)
			if e == nil && ans < a {
				ans = a
			}
		}
	}

	fmt.Println(ans)
}
