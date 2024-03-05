package main

import (
	"fmt"
	autopsy "matrix/Autopsy"
	"matrix/La"
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
	for i := 0; i < 1; i++ {
		mat := La.RandomMatrix(2, 2)
		fmt.Printf("%s", mat.ToString())
		eigs := mat.EigenValues()
		for j := 0; j < len(eigs); j++ {
			fmt.Printf("%s,", utils.FormatComplex(eigs[j]))
		}
		println()
		vecs := mat.EigenVectors()
		for j := 0; j < len(vecs); j++ {

			println(vecs[j].ToString())
		}
	}

	//poly := al.PolynomialFromString("7x^5-2x^4+x-3")

	return
}
