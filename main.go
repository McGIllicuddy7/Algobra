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
	a := algebra.PolynomialFromString("2+3x^2")
	println(a.ToString())
	b := algebra.PolynomialFromString("2-3x^2")
	println(b.ToString())
	c := algebra.PolynomialMult(a, b)
	println(c.ToString())
	return
}
