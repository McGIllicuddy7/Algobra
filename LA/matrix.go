package LA

import (
	"math/cmplx"
	"math/rand"
	"matrix/utils"
)

type Matrix struct {
	data   []complex128
	height int
	width  int
}
type Vector []complex128

func MatrixFromInts(data [][]int) Matrix {
	height := len(data)
	width := len(data[0])
	out := Matrix{make([]complex128, height*width), height, width}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out.Set(x, y, complex(float64(data[y][x]), 0))
		}
	}
	return out
}
func MatrixFromFloats(data [][]float64) Matrix {
	height := len(data)
	width := len(data[0])
	out := Matrix{make([]complex128, height*width), height, width}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out.Set(x, y, complex(data[y][x], 0))
		}
	}
	return out
}
func VectorAdd(v0 Vector, v1 Vector) Vector {
	out := make([]complex128, len(v0))
	for i := 0; i < len(v0); i++ {
		out[i] = v0[i] + v1[i]
	}
	return out
}
func VectorScale(v0 Vector, v1 Vector) Vector {
	out := make([]complex128, len(v0))
	for i := 0; i < len(v0); i++ {
		out[i] = v0[i] - v1[i]
	}
	return out
}
func VectorMlt(v0 Vector, s complex128) Vector {
	out := make(Vector, len(v0))
	for i := 0; i < len(v0); i++ {
		out[i] = v0[i] * s
	}
	return out
}
func VectorDot(v0 Vector, v1 Vector) complex128 {
	out := complex128(0)
	for i := 0; i < len(v0); i++ {
		out += v0[i] * v1[i]
	}
	return out
}
func vector_length(v0 Vector) complex128 {
	return cmplx.Sqrt(VectorDot(v0, v0))
}
func (this *Matrix) Get(x int, y int) complex128 {
	return this.data[y*this.width+x]
}
func (this *Matrix) Set(x int, y int, v complex128) {
	this.data[y*this.width+x] = v
}
func (this *Matrix) NumCols() int {
	return this.width
}
func (this *Matrix) NumRows() int {
	return this.height
}
func (this *Matrix) Clone() Matrix {
	out := Matrix{make([]complex128, len(this.data)), this.height, this.width}
	for i := 0; i < len(this.data); i++ {
		out.data[i] = this.data[i]
	}
	return out
}
func (this *Matrix) ToString() string {
	out := ""
	out_strs := make([]string, 0, this.height*this.width)
	for j := 0; j < this.height; j++ {
		for i := 0; i < this.width; i++ {
			out_strs = append(out_strs, utils.FormatComplex(this.Get(i, j)))
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
	for j := 0; j < this.height; j++ {
		for i := 0; i < this.width; i++ {
			out += out_strs[j*this.width+i]
			if i < this.width-1 {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}
func (this *Matrix) SwapRows(r0 int, r1 int) {
	for i := 0; i < this.width; i++ {
		tmp0 := this.Get(i, r0)
		tmp1 := this.Get(i, r1)
		this.Set(i, r1, tmp0)
		this.Set(i, r0, tmp1)
	}
}

// adds r0 to r1 scaled by s
func (this *Matrix) AddRows(r0 int, r1 int, s complex128) {
	for i := 0; i < this.width; i++ {
		tmp0 := this.Get(i, r0) * s
		this.Set(i, r1, tmp0+this.Get(i, r1))
	}
}
func (this *Matrix) SubRows(r0 int, r1 int, s complex128) {
	for i := 0; i < this.width; i++ {
		tmp0 := this.Get(i, r0) * s
		this.Set(i, r1, this.Get(i, r1)-tmp0)
	}
}
func (this *Matrix) ScaleRow(r0 int, s complex128) {
	for i := 0; i < this.width; i++ {
		tmp0 := this.Get(i, r0) * s
		this.Set(i, r0, tmp0)
	}
}
func Identity(n int) Matrix {
	out := Matrix{make([]complex128, n*n), n, n}
	for i := 0; i < n; i++ {
		out.Set(i, i, 1)
	}
	return out
}
func MatrixAdd(m0 Matrix, m1 Matrix) Matrix {
	if m0.height != m1.height || m0.width != m1.width {
		panic("error adding matrices without the same dimension")
	}
	out := Matrix{make([]complex128, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < m0.height; i++ {
		for j := 0; j < m0.width; j++ {
			out.Set(j, i, m0.Get(j, i)+m1.Get(j, i))
		}
	}
	return out
}
func MatrixSub(m0 Matrix, m1 Matrix) Matrix {
	if m0.height != m1.height || m0.width != m1.width {
		panic("error subtracting matrices without the same dimension")
	}
	out := Matrix{make([]complex128, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < m0.height; i++ {
		for j := 0; j < m0.width; j++ {
			out.Set(j, i, m1.Get(j, i)-m0.Get(j, i))
		}
	}
	return out
}
func MatrixScale(m0 Matrix, s complex128) Matrix {
	out := Matrix{make([]complex128, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < len(m0.data); i++ {
		out.data[i] = out.data[i] * s
	}
	return out
}
func MatrixRowReduce(matrx Matrix) Matrix {
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
func MatrixPairRowReduce(source Matrix, target Matrix) (Matrix, Matrix) {
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
func (this *Matrix) Determinant() complex128 {
	mtrx := Matrix{make([]complex128, len(this.data)), this.height, this.width}
	out := complex(1, 0)
	for i := 0; i < len(this.data); i++ {
		mtrx.data[i] = this.data[i]
	}
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
			out *= -1
		}
		v := mtrx.Get(i, i)
		mtrx.ScaleRow(i, 1/v)
		out *= v
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
	for i := 0; i < this.width; i++ {
		out *= mtrx.Get(i, i)
	}
	return out
}
func (this *Matrix) Inverse() Matrix {
	_, out := MatrixPairRowReduce(*this, Identity(this.width))
	return out
}
func RandomMatrix(height int, width int) Matrix {
	var out Matrix
	out.data = make([]complex128, height*width)
	out.height = height
	out.width = width
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out.Set(x, y, complex(float64(rand.Int31()%10+1), 0))
		}
	}
	return out
}
