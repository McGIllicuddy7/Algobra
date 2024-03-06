package main

import (
	"math/rand"
	autopsy "matrix/Autopsy"
	"matrix/La"
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
func new_int() int {
	return int(rand.Int31()%10 - 5)
}
func main() {
	autopsy.Init()
	mat := La.ComplexMatrixFromInts([][]int{{new_int(), new_int()}, {0, 0}})
	vs := mat.Solve(La.ZeroVector(2))
	println(vs.ToString())
}
