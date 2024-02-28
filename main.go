package main

import (
	"fmt"
	"math/rand"
	autopsy "matrix/Autopsy"
	La "matrix/La"
	fr "matrix/fractions"
	"os"
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
	//v := LA.MatrixFromInts([][]int{{1, 1, 4}, {7, 9, 3}, {6, 4, 6}})
	//v := LA.RandomMatrix(3, 3)
	autopsy.Init()
	for i := 0; i < 100000; i++ {
		v := La.RandomMatrix(4, 4)
		v1 := v.Determinant()
		u := v.ToPolyMatrix()
		t := u.CharacteristicPolynomial()
		if !fr.Equals(t.ZeroCoef(), v1) {
			println("failed\n")
			println(v.ToString())
			println(t.ToString())
			println(v1.ToString())
			os.Exit(1)
		}
	}
	println("success")
	return
}
