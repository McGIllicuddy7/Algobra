package main

import (
	"fmt"
	"math/rand"
	La "matrix/LA"
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
	for i := 1; i < 5; i++ {
		mat := La.RandomMatrix(i, i)
		//println(i)
		poly := mat.Determinant()
		println(poly.ToString())
	}
	return
}
