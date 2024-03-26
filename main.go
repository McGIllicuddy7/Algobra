package main

import (
	autopsy "matrix/Autopsy"
	La "matrix/La"
)

func main() {
	autopsy.Init()
	for i := 0; i < 40000; i++ {
		mat := La.RandomComplexMatrix(3, 3)
		vs := mat.Solve(La.ZeroVector(3))
		v := mat.MultByVector(vs)
		if !La.VectorEqual(v, La.ZeroVector(3)) {
			println("i =", i)
			autopsy.Dump()
			//println(mat.ToString())
			println(vs.ToString())
			println(v.ToString())
			panic("failed\n")
		}
		autopsy.Reset()
	}
	println("Great Success :3")
}
