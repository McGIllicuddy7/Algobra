package main

import (
	"fmt"
	"math/rand"
	algebra "matrix/Algebra"
)

func make_ints(amnt int) []int {
	out := make([]int, amnt)
	for i := 0; i < amnt; i++ {
		out[i] = int(rand.Int31()) % (amnt * 2)
	}
	return out
}
func print_ints(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
		if i != len(arr)-1 {
			print(",")
		}
	}
}
func main() {
	v := algebra.ParsePolynomial("x^-2")
	q := algebra.ParsePolynomial("x")
	v = algebra.PolynomialMult(v, q)
	println(v.ToString())
	err, v := algebra.PolynomialIntegrate(v)
	if err != nil {
		println(err.Error())
	}
	println(v.ToString())
	v = algebra.PolynonialDerivitive(v)
	println(v.ToString())
	return
}
