package LA

import (
	al "matrix/Algebra"
	"matrix/utils"
)

type PolyMatrix struct {
	data   []al.Polynomial
	height int
	width  int
}

func (this *PolyMatrix) Get(x int, y int) al.Polynomial {
	return this.data[y*this.width+x]
}
func (this *PolyMatrix) Set(x int, y int, v al.Polynomial) {
	this.data[y*this.height+x] = v
}
func (this *PolyMatrix) ToString() string {
	strings := make([]string, this.height*this.width)
	for i := 0; i < len(this.data); i++ {
		strings[i] = this.data[i].ToString()
	}
	strings = utils.NormalizeStrlens(strings)
	out := ""
	for y := 0; y < this.height; y++ {
		for x := 0; x < this.width; x++ {
			out += strings[y*this.height+x]
			if x < this.width-1 {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}
func (this *Matrix) ToEigenMatrix() PolyMatrix {
	out := PolyMatrix{make([]al.Polynomial, this.height*this.width), this.height, this.width}
	for y := 0; y < this.height; y++ {
		for x := 0; x < this.width; x++ {
			var v al.Polynomial
			if x == y {
				v = al.CompPoly(this.Get(x, y), -1, 1)
			} else {
				v = al.CompPoly(this.Get(x, y), 0, 0)
			}
			out.Set(x, y, v)
		}
	}
	return out
}
func (this *Matrix) ToPolyMatrix() PolyMatrix {
	out := PolyMatrix{make([]al.Polynomial, this.height*this.width), this.height, this.width}
	for y := 0; y < this.height; y++ {
		for x := 0; x < this.width; x++ {
			v := al.CompPoly(this.Get(x, y), 0, 0)
			out.Set(x, y, v)
		}
	}
	return out
}
func (this *PolyMatrix) elimRowCollumn(idx int) PolyMatrix {
	out := PolyMatrix{make([]al.Polynomial, (this.width-1)*(this.height-1)), this.height - 1, this.width - 1}
	for y := 1; y < this.height; y++ {
		dy := y - 1
		for x := 0; x < this.width; x++ {
			dx := x
			if dx == idx {
				continue
			}
			if x > idx {
				dx--
			}
			out.Set(dx, dy, this.Get(x, y))
		}
	}
	return out
}
func (this PolyMatrix) CharacteristicPolynomial() al.Polynomial {
	if this.width == 2 && this.height == 2 {
		//print(this.ToString())
		a := this.Get(0, 0)
		b := this.Get(0, 1)
		c := this.Get(1, 0)
		d := this.Get(1, 1)
		ad := al.PolynomialMult(a, d)
		bc := al.PolynomialMult(b, c)
		ret := al.PolynomialSub(ad, bc)
		//println(ad.ToString(), bc.ToString(), ret.ToString())
		//print("\n")
		return ret
	}
	var out al.Polynomial
	for i := 0; i < this.width; i++ {
		tmp := this.elimRowCollumn(i)
		m := this.Get(i, 0)
		mdet := al.PolynomialMult(m, tmp.CharacteristicPolynomial())
		if i%2 == 0 {
			out = al.PolynomialAdd(out, mdet)
		} else {
			mdet = al.PolynomialScale(mdet, -1)
			out = al.PolynomialAdd(out, mdet)
		}
	}
	return out
}
