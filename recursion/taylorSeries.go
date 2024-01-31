package main

func getTaylor(x float32, n float32) func(float32, float32) float32 {
	var p float32
	var f float32
	p = 1
	f = 1
	var r float32
	r = float32(1)
	var TaylorClosure func(float32, float32) float32
	TaylorClosure = func(x, n float32) float32 {
		if n == 0 {
			return 1
		} else {
			r = TaylorClosure(x, n-1)
			p = p * x
			f = f * n
			println(r + float32(p/f))
			println(r + (p / f))
			return r + float32(p/f)
		}
	}
	return TaylorClosure
}
