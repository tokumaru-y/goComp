// ref:https://qiita.com/ganyariya/items/1553ff2bf8d6d7789127
func ThreePartSearch(l, r, P float64) float64 {
	f := func(x float64) float64 {
		return x + (P / math.Pow(2, x/1.5))
	}
	for i := 0; i < 200; i++ {
		c1 := (l*2 + r) / 3
		c2 := (l + r*2) / 3
		if f(c1) > f(c2) {
			l = c1
		} else {
			r = c2
		}
	}
	return f(l)
}