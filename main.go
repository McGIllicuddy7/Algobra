package main

import (
	"fmt"
	"math/rand"
	"matrix/LA"
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
	v := LA.RandomMatrix(2, 2)
	q := v.ToEigenMatrix()
	u := q.CharacteristicPolynomial()
	println(q.ToString())
	println(u.ToString())
	return
}
