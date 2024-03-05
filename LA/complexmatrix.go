package La

import "matrix/utils"

type MatrixComplex struct {
	data   []complex128
	height int
	width  int
}

func (tmat *MatrixComplex) Get(x int, y int) complex128 {
	return tmat.data[y*tmat.width+x]
}
func (tmat *MatrixComplex) Set(x int, y int, v complex128) {
	tmat.data[y*tmat.width+x] = v
}
func (tmat *MatrixComplex) NumCols() int {
	return tmat.width
}
func (tmat *MatrixComplex) NumRows() int {
	return tmat.width
}
func (tmat *MatrixComplex) Clone() MatrixComplex {
	out := MatrixComplex{make([]complex128, len(tmat.data)), tmat.height, tmat.width}
	copy(out.data, tmat.data)
	return out
}
func (tmat *MatrixComplex) ToString() string {
	out := ""
	out_strs := make([]string, 0, tmat.height*tmat.width)
	for j := 0; j < tmat.height; j++ {
		for i := 0; i < tmat.width; i++ {
			out_strs = append(out_strs, utils.FormatComplex(tmat.Get(i, j)))
		}
	}
	max := 0
	for i := 0; i < len(out_strs); i++ {
		if len(out_strs[i]) > max {
			max = len(out_strs[i])
		}
	}
	for i := 0; i < len(out_strs); i++ {
		if out_strs[i][0] != '-' {
			out_strs[i] = " " + out_strs[i]
		}
		for len(out_strs[i]) < max {
			out_strs[i] += " "
		}
	}
	for j := 0; j < tmat.height; j++ {
		for i := 0; i < tmat.width; i++ {
			out += out_strs[j*tmat.width+i]
			if i < tmat.width-1 {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}
func (tmat *MatrixComplex) SwapRows(r0 int, r1 int) {
	for i := 0; i < tmat.width; i++ {
		tmp0 := tmat.Get(i, r0)
		tmp1 := tmat.Get(i, r1)
		tmat.Set(i, r1, tmp0)
		tmat.Set(i, r0, tmp1)
	}
}
