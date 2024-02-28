package algebra

import (
	"errors"
	"math/cmplx"
	"matrix/utils"
)

type polycule struct {
	coef complex128
	pow  int
}
type Polynomial struct {
	data []polycule
}

func (this *Polynomial) Clone() Polynomial {
	out := Polynomial{make([]polycule, len(this.data))}
	copy(out.data, this.data)
	return out
}
func polycule_cmp(a polycule, b polycule) int {
	if a.pow > b.pow {
		return -1
	}
	if a.pow < b.pow {
		return 1
	}
	return 0
}
func polyculeMlt(a polycule, b polycule) polycule {
	return polycule{a.coef * b.coef, a.pow + b.pow}
}
func (this *Polynomial) addPolycule(p polycule) {
	this.data = append(this.data, p)
}
func slice_contains(data []int, value int) bool {
	for i := 0; i < len(data); i++ {
		if data[i] == value {
			return true
		}
	}
	return false
}
func (this *Polynomial) compress() {
	utils.SortInplace[polycule](this.data, polycule_cmp)
	v := this.data[0].pow
	new_data := make([]polycule, 0)
	counter := 0
	new_data = append(new_data, polycule{0, this.data[counter].pow})
	for counter < len(this.data) {
		if this.data[counter].pow != v {
			v = this.data[counter].pow
			new_data = append(new_data, polycule{0, this.data[counter].pow})
		}
		new_data[len(new_data)-1].coef += this.data[counter].coef
		counter++
	}
	this.data = new_data
}
func PolynomialAdd(a Polynomial, b Polynomial) Polynomial {
	var out Polynomial
	out.data = make([]polycule, 0)
	out.data = append(out.data, a.data...)
	out.data = append(out.data, b.data...)
	out.compress()
	return out
}
func PolynomialSub(a Polynomial, b Polynomial) Polynomial {
	var out Polynomial
	tmp := PolynomialScale(b, -1)
	out.data = make([]polycule, 0)
	out.data = append(out.data, a.data...)
	out.data = append(out.data, tmp.data...)
	out.compress()
	return out
}
func polyculeMultByPolynomial(a Polynomial, b polycule) Polynomial {
	out := a.Clone()
	for i := 0; i < len(out.data); i++ {
		out.data[i] = polyculeMlt(out.data[i], b)
	}
	out.compress()
	return out
}
func PolynomialMult(a Polynomial, b Polynomial) Polynomial {
	out := Polynomial{make([]polycule, 0)}
	for i := 0; i < len(b.data); i++ {
		out = PolynomialAdd(out, polyculeMultByPolynomial(a, b.data[i]))
	}
	out.compress()
	return out
}
func PolynomialScale(a Polynomial, s complex128) Polynomial {
	out := a.Clone()
	for i := 0; i < len(a.data); i++ {
		out.data[i].coef = a.data[i].coef * s
	}
	return out
}
func PolynomialScaleInplace(a Polynomial, s complex128) {
	for i := 0; i < len(a.data); i++ {
		a.data[i].coef *= s
	}
}
func PolynonialDerivitive(a Polynomial) Polynomial {
	out := a.Clone()
	for i := 0; i < len(a.data); i++ {
		if a.data[i].pow != 0 {
			out.data[i].coef = a.data[i].coef * complex(float64(a.data[i].pow), 0)
			out.data[i].pow = a.data[i].pow - 1
		} else {
			out.data[i].coef = 0
			out.data[i].pow = 0
		}
	}
	return out
}
func PolynomialIntegrate(a Polynomial) (error, Polynomial) {
	out := a.Clone()
	err := error(nil)
	for i := 0; i < len(a.data); i++ {
		if a.data[i].pow != -1 {
			out.data[i].coef = a.data[i].coef / complex(float64(a.data[i].pow+1), 0)
			out.data[i].pow = a.data[i].pow + 1
		} else {
			out.data[i].coef = 0
			out.data[i].pow = 0
			err = errors.New("error unsupported function type \"natural log\"")
		}
	}
	return err, out
}
func polycule_evaluate(p polycule, x complex128) complex128 {
	return p.coef * cmplx.Pow(x, complex(float64(p.pow), 0))
}
func PolynomialEvalulate(a Polynomial, x complex128) complex128 {
	total := complex128(0)
	for i := 0; i < len(a.data); i++ {
		total += polycule_evaluate(a.data[i], x)
	}
	return total
}
func RealPoly(addval float64, coef float64, pow int) Polynomial {
	out := Polynomial{make([]polycule, 0)}
	v1 := polycule{complex(addval, 0), 0}
	v2 := polycule{complex(coef, 0), pow}
	out.addPolycule(v1)
	out.addPolycule(v2)
	out.compress()
	return out
}
func CompPoly(addval complex128, coef complex128, pow int) Polynomial {
	out := Polynomial{make([]polycule, 0)}
	v1 := polycule{addval, 0}
	v2 := polycule{coef, pow}
	out.addPolycule(v1)
	out.addPolycule(v2)
	out.compress()
	return out
}
