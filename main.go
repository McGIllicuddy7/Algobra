package main

import (
	"math/rand"
	autopsy "matrix/Autopsy"
	La "matrix/La"
)

/*
#include <stdio.h>
void Read(char * start, size_t size){
	fread(stdin, start, size);
}
*/

//const count = 1000000

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
