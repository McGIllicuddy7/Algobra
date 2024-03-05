package La

import (
	"math/rand"
	"matrix/utils"
)

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
func (tmat *MatrixComplex) AddRows(r0 int, r1 int, s complex128) {
	for i := 0; i < tmat.width; i++ {
		tmp0 := tmat.Get(i, r0) * s
		tmat.Set(i, r1, tmp0+tmat.Get(i, r1))
	}
}
func (tmat *MatrixComplex) SubRows(r0 int, r1 int, s complex128) {
	for i := 0; i < tmat.width; i++ {
		tmp0 := tmat.Get(i, r0) * s
		tmat.Set(i, r1, tmat.Get(i, r1)-tmp0)
	}
}
func (tmat *MatrixComplex) ScaleRow(r0 int, s complex128) {
	for i := 0; i < tmat.width; i++ {
		tmp0 := tmat.Get(i, r0) * s
		tmat.Set(i, r0, tmp0)
	}
}
func ComplexIdentity(n int) MatrixComplex {
	out := MatrixComplex{make([]complex128, n*n), n, n}
	for i := 0; i < n; i++ {
		out.Set(i, i, 1)
	}
	return out
}
func ComplexMatrixAdd(m0 MatrixComplex, m1 MatrixComplex) MatrixComplex {
	if m0.height != m1.height || m0.width != m1.width {
		panic("error adding matrices without the same dimension")
	}
	out := MatrixComplex{make([]complex128, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < m0.height; i++ {
		for j := 0; j < m0.width; j++ {
			out.Set(j, i, (m0.Get(j, i) + m1.Get(j, i)))
		}
	}
	return out
}
func ComplexMatrixSub(m0 MatrixComplex, m1 MatrixComplex) MatrixComplex {
	if m0.height != m1.height || m0.width != m1.width {
		panic("error adding matrices without the same dimension")
	}
	out := MatrixComplex{make([]complex128, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < m0.height; i++ {
		for j := 0; j < m0.width; j++ {
			out.Set(j, i, (m1.Get(j, i) - m0.Get(j, i)))
		}
	}
	return out
}
func ComplexMatrixScale(m0 MatrixComplex, s complex128) MatrixComplex {
	out := MatrixComplex{make([]complex128, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < len(m0.data); i++ {
		out.data[i] = m0.data[i] * s
	}
	return out
}
func ComplexMatrixRowReduce(matrx MatrixComplex) MatrixComplex {
	mtrx := matrx.Clone()
	for i := 0; i < mtrx.width; i++ {
		r := i
		degen := false
		for mtrx.Get(i, r) == 0 {
			r++
			if r >= mtrx.height {
				degen = true
				break
			}
		}
		if degen {
			continue
		}
		if r != i {
			mtrx.SwapRows(r, i)
		}
		v := mtrx.Get(i, i)
		mtrx.ScaleRow(i, 1/v)
		for j := 0; j < mtrx.height; j++ {
			if j == i {
				continue
			}
			mlt := mtrx.Get(i, j)
			mtrx.SubRows(i, j, mlt)
		}
	}
	return mtrx
}
func ComplexMatrixFromInts(slice [][]int) MatrixComplex {
	height := len(slice)
	width := len(slice[0])
	out := MatrixComplex{make([]complex128, height*width), height, width}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out.Set(x, y, complex(float64(slice[y][x]), 0))
		}
	}
	return out
}
func ComplexMatrixPairRowReduce(source MatrixComplex, target MatrixComplex) (MatrixComplex, MatrixComplex) {
	mtrx := source.Clone()
	out := target.Clone()
	for i := 0; i < mtrx.width; i++ {
		r := i
		degen := false
		for mtrx.Get(i, r) == 0 {
			r++
			if r >= mtrx.height {
				degen = true
				break
			}
		}
		if degen {
			continue
		}
		if r != i {
			out.SwapRows(r, i)
			mtrx.SwapRows(r, i)
		}
		v := mtrx.Get(i, i)
		mtrx.ScaleRow(i, 1/v)
		out.ScaleRow(1, 1/v)
		for j := 0; j < mtrx.height; j++ {
			if j == i {
				continue
			}
			mlt := mtrx.Get(i, j)
			mtrx.SubRows(i, j, mlt)
			out.SubRows(i, j, mlt)
		}
	}
	return mtrx, out
}
func (tmat *MatrixComplex) Determinant() complex128 {
	mtrx := MatrixComplex{make([]complex128, len(tmat.data)), tmat.height, tmat.width}
	out := complex128(1)
	copy(mtrx.data, tmat.data)
	for i := 0; i < mtrx.width; i++ {
		r := i
		degen := false
		for mtrx.Get(i, r) == 0 {
			r++
			if r >= mtrx.height {
				degen = true
				break
			}
		}
		if degen {
			return 0
		}
		if r != i {
			mtrx.SwapRows(r, i)
			out *= -1
		}
		v := mtrx.Get(i, i)
		mtrx.ScaleRow(i, 1/v)
		out = v * out
		for j := 0; j < mtrx.height; j++ {
			if j == i {
				continue
			}
			mlt := mtrx.Get(i, j)
			if mlt == 0 {
				continue
			}
			mtrx.SubRows(i, j, mlt)
		}
	}
	return out
}
func (tmat *MatrixComplex) Inverse() MatrixComplex {
	_, out := ComplexMatrixPairRowReduce(*tmat, ComplexIdentity(tmat.width))
	return out
}
func RandomComplexMatrix(height int, width int) MatrixComplex {
	var out MatrixComplex
	out.data = make([]complex128, height*width)
	out.height = height
	out.width = width
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out.Set(x, y, complex(float64(rand.Int31()%10-5), 0))
		}
	}
	return out
}
func (tmat *MatrixComplex) ToUpperTriangular() MatrixComplex {
	mtrx := tmat.Clone()
	for i := 0; i < mtrx.width; i++ {
		r := i
		degen := false
		for mtrx.Get(i, r) == 0 {
			r++
			if r >= mtrx.height {
				degen = true
				break
			}
		}
		if degen {
			continue
		}
		if r != i {
			mtrx.SwapRows(r, i)
		}
		v := mtrx.Get(i, i)
		mtrx.ScaleRow(i, 1/v)
		for j := r; j < mtrx.height; j++ {
			if j == i {
				continue
			}
			mlt := mtrx.Get(i, j)
			mtrx.SubRows(i, j, mlt)
		}
	}
	return mtrx
}
func (tmat *MatrixComplex) Solve(values Vector) Vector {
	mtrx := tmat.Clone()
	vals := values.Clone()
	for i := 0; i < mtrx.width; i++ {
		r := i
		degen := false
		for mtrx.Get(i, r) == 0 {
			r++
			if r >= mtrx.height {
				degen = true
				break
			}
		}
		if degen {
			continue
		}
		if r != i {
			mtrx.SwapRows(r, i)
			vals.Swap(r, i)
		}
		v := mtrx.Get(i, i)
		recip := 1 / v
		mtrx.ScaleRow(i, recip)
		vals[i] *= recip
		for j := r; j < mtrx.height; j++ {
			if j == i {
				continue
			}
			mlt := mtrx.Get(i, j)
			mtrx.SubRows(i, j, mlt)
			vals[j] -= vals[i] * mlt
		}
	}
	symbolTable := make(Vector, tmat.width)
	definedSymbols := make([]bool, tmat.width)
	for i := 0; i < len(symbolTable); i++ {
		symbolTable[i] = 0
		definedSymbols[i] = false
	}
	for y := tmat.height - 1; y >= 0; y-- {
		syms := make([]int, 0)
		for x := 0; x < tmat.width; x++ {
			if mtrx.Get(x, y) == 0 {
				syms = append(syms, x)
			}
		}
		undefined := make([]int, 0)
		for i := 0; i < len(syms); i++ {
			if !definedSymbols[syms[i]] {
				undefined = append(undefined, syms[i])
			}
		}
		for i := len(undefined) - 1; i > 0; i-- {
			idx := undefined[i]
			symbolTable[idx] = 1
			definedSymbols[idx] = true
		}
		if len(undefined) > 0 {
			//println(undefined[0])
			newSym := complex128(0)
			for i := 1; i < len(syms); i++ {
				newSym -= symbolTable[syms[i]] * tmat.Get(syms[i], y)
			}
			newSym += vals[y]
			newSym /= mtrx.Get(undefined[0], y)
			symbolTable[undefined[0]] = newSym
			definedSymbols[undefined[0]] = true
		}
	}
	return symbolTable
}
func (tmat *MatrixComplex) MultByVector(v Vector) Vector {
	out := make(Vector, len(v))
	for i := 0; i < tmat.height; i++ {
		total := complex128(0)
		for j := 0; j < tmat.width; j++ {
			total += tmat.Get(j, i) * v[j]
		}
		out[i] = total
	}
	return out
}
