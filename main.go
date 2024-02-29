package main

import (
	autopsy "matrix/Autopsy"
	al "matrix/algebra"
	"matrix/utils"
)

const count = 1000000

/*
	func polynomial(x complex128) complex128 {
		return 2*x*x - 4*x + 5
	}

	func evaluateCompiled() float64 {
		t := time.Now()
		for i := 0; i < count; i++ {
			r := polynomial(utils.RandomComplex())
			if rand.Int31()%count*10 == 0 {
				autopsy.Store(utils.FormatComplex(r))
			}
		}
		out := time.Now().Sub(t).Seconds()
		return out
	}

	func evaluateParsed(poly al.Polynomial) float64 {
		t := time.Now()
		for i := 0; i < count; i++ {
			r := poly.EvaluateComplex(utils.RandomComplex())
			if rand.Int31()%count*10 == 0 {
				autopsy.Store(utils.FormatComplex(r))
			}
		}
		out := time.Now().Sub(t).Seconds()
		return out
	}
*/
func main() {
	autopsy.Init()
	poly := al.PolynomialFromString("7x^5-2x^4+x-3")
	println(poly.ToString())
	roots := poly.FindZeros()
	for i := 0; i < len(roots); i++ {
		println("root:", utils.FormatComplex(roots[i]))
		println("value:", utils.FormatComplex(poly.EvaluateComplex(roots[i])))
		println()
	}
	return
}
