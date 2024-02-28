package La

import (
	"math/rand"
	fr "matrix/fractions"
)

type Matrix struct {
	data   []fr.Fraction
	height int
	width  int
}

func (this *Matrix) Get(x int, y int) fr.Fraction {
	return this.data[y*this.width+x]
}
func (this *Matrix) Set(x int, y int, v fr.Fraction) {
	this.data[y*this.width+x] = v
}
func (this *Matrix) NumCols() int {
	return this.width
}
func (this *Matrix) NumRows() int {
	return this.height
}
func (this *Matrix) Clone() Matrix {
	out := Matrix{make([]fr.Fraction, len(this.data)), this.height, this.width}
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
			out_strs = append(out_strs, this.Get(i, j).ToString())
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
func (this *Matrix) AddRows(r0 int, r1 int, s fr.Fraction) {
	for i := 0; i < this.width; i++ {
		tmp0 := fr.Mult(this.Get(i, r0), s)
		this.Set(i, r1, fr.Add(tmp0, this.Get(i, r1)))
	}
}
func (this *Matrix) SubRows(r0 int, r1 int, s fr.Fraction) {
	for i := 0; i < this.width; i++ {
		tmp0 := fr.Mult(this.Get(i, r0), s)
		this.Set(i, r1, fr.Sub(this.Get(i, r1), tmp0))
	}
}
func (this *Matrix) ScaleRow(r0 int, s fr.Fraction) {
	for i := 0; i < this.width; i++ {
		tmp0 := fr.Mult(this.Get(i, r0), s)
		this.Set(i, r0, tmp0)
	}
}
func Identity(n int) Matrix {
	out := Matrix{make([]fr.Fraction, n*n), n, n}
	for i := 0; i < n; i++ {
		out.Set(i, i, fr.NewFrac(1, 1))
	}
	return out
}
func MatrixAdd(m0 Matrix, m1 Matrix) Matrix {
	if m0.height != m1.height || m0.width != m1.width {
		panic("error adding matrices without the same dimension")
	}
	out := Matrix{make([]fr.Fraction, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < m0.height; i++ {
		for j := 0; j < m0.width; j++ {
			out.Set(j, i, fr.Add(m0.Get(j, i), m1.Get(j, i)))
		}
	}
	return out
}
func MatrixSub(m0 Matrix, m1 Matrix) Matrix {
	if m0.height != m1.height || m0.width != m1.width {
		panic("error subtracting matrices without the same dimension")
	}
	out := Matrix{make([]fr.Fraction, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < m0.height; i++ {
		for j := 0; j < m0.width; j++ {
			out.Set(j, i, fr.Sub(m1.Get(j, i), m0.Get(j, i)))
		}
	}
	return out
}
func MatrixScale(m0 Matrix, s fr.Fraction) Matrix {
	out := Matrix{make([]fr.Fraction, m0.height*m0.width), m0.height, m0.width}
	for i := 0; i < len(m0.data); i++ {
		out.data[i] = fr.Mult(m0.data[i], s)
	}
	return out
}
func MatrixRowReduce(matrx Matrix) Matrix {
	mtrx := matrx.Clone()
	for i := 0; i < mtrx.width; i++ {
		r := i
		degen := false
		for fr.Equals(mtrx.Get(i, r), fr.NewFrac(0, 1)) {
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
		mtrx.ScaleRow(i, fr.Recip(v))
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
		for fr.Equals(mtrx.Get(i, r), fr.NewFrac(0, 1)) {
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
		mtrx.ScaleRow(i, fr.Recip(v))
		out.ScaleRow(1, fr.Recip(v))
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
func (this *Matrix) Determinant() fr.Fraction {
	mtrx := Matrix{make([]fr.Fraction, len(this.data)), this.height, this.width}
	out := fr.NewFrac(1, 1)
	for i := 0; i < len(this.data); i++ {
		mtrx.data[i] = this.data[i]
	}
	for i := 0; i < mtrx.width; i++ {
		r := i
		degen := false
		for fr.Equals(mtrx.Get(i, r), fr.NewFrac(0, 1)) {
			r++
			if r >= mtrx.height {
				degen = true
				break
			}
		}
		if degen {
			return fr.NewFrac(0, 1)
		}
		if r != i {
			mtrx.SwapRows(r, i)
			out = fr.Scale(out, -1)
		}
		v := mtrx.Get(i, i)
		mtrx.ScaleRow(i, fr.Recip(v))
		out = fr.Mult(out, v)
		for j := 0; j < mtrx.height; j++ {
			if j == i {
				continue
			}
			mlt := mtrx.Get(i, j)
			if fr.Equals(mlt, fr.NewFrac(0, 1)) {
				continue
			}
			mtrx.SubRows(i, j, mlt)
		}
	}
	return out
}
func (this *Matrix) Inverse() Matrix {
	_, out := MatrixPairRowReduce(*this, Identity(this.width))
	return out
}
func RandomMatrix(height int, width int) Matrix {
	var out Matrix
	out.data = make([]fr.Fraction, height*width)
	out.height = height
	out.width = width
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out.Set(x, y, fr.NewFrac(int(rand.Int31()%10-5), 1))
		}
	}
	return out
}
