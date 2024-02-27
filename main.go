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
	v := algebra.ParsePolynomial("x-2")
	q := algebra.ParsePolynomial("x+3")
	v = algebra.PolynomialMult(v, q)
	str := v.ToString()
	println("(x-2)(x+3) =", str)
	err, v := algebra.PolynomialIntegrate(v)
	if err != nil {
		println(err.Error())
	}
	str2 := v.ToString()
	println("integral(", str, ") = ", str2)
	v = algebra.PolynonialDerivitive(v)
	println("derivitive(", str2, ") = ", v.ToString())
	return
}
