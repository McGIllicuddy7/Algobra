package main

import (
	"fmt"
	"math"
	"math/rand"
	"matrix/LA"
	"os"
	"strconv"
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
	for i := 0; i < 10000; i++ {
		v := LA.RandomMatrix(3, 3)
		v1 := v.Determinant()
		tmp := v.ToPolyMatrix().CharacteristicPolynomial()
		v2, err := strconv.ParseComplex(tmp.ToString(), 128)
		if err != nil {
			println("failed\n")
			os.Exit(1)
		}
		if math.Round(real(v2)) != math.Round(real(v1)) {
			println(v1, "!=", v2)
			println("failed\n")
			os.Exit(1)
		}
		if i%1000 == 0 {
			print(v.ToString())
			println("detertminant from guass jordan: ", real(v1))
			println("determinant from recursive: ", real(v2))
		}

	}
	println("success")
	return
}
