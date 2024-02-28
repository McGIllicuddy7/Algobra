package algebra

import (
	fr "matrix/fractions"
	"matrix/utils"
)

type polycule struct {
	coef fr.Fraction
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
	return polycule{fr.Mult(a.coef, b.coef), a.pow + b.pow}
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
	if len(this.data) < 2 {
		return
	}
	new_data := make([]polycule, 0)
	powers := make([]int, 0)
	for i := 0; i < len(this.data); i++ {
		if !slice_contains(powers, this.data[i].pow) && !fr.Equals(this.data[i].coef, fr.FromFloat(0)) {
			powers = append(powers, this.data[i].pow)
		}
	}
	for i := 0; i < len(powers); i++ {
		p := polycule{fr.NewFrac(0, 1), powers[i]}
		for j := 0; j < len(this.data); j++ {
			if this.data[j].pow == powers[i] {
				p.coef = fr.Add(p.coef, this.data[j].coef)
			}
		}
		new_data = append(new_data, p)
	}
	this.data = new_data
}
func (this Polynomial) ZeroCoef() fr.Fraction {
	return this.data[0].coef
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
	tmp := PolynomialScale(b, fr.NewFrac(-1, 1))
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
func PolynomialScale(a Polynomial, s fr.Fraction) Polynomial {
	out := a.Clone()
	for i := 0; i < len(a.data); i++ {
		out.data[i].coef = fr.Mult(a.data[i].coef, s)
	}
	return out
}
func PolynomialScaleInplace(a Polynomial, s fr.Fraction) {
	for i := 0; i < len(a.data); i++ {
		a.data[i].coef = fr.Mult(a.data[i].coef, s)
	}
}
func NewPoly(coef fr.Fraction, pow int) Polynomial {
	return Polynomial{[]polycule{{coef, pow}}}
}
func CompPoly(coef fr.Fraction, ceof2 fr.Fraction, pow int) Polynomial {
	return Polynomial{[]polycule{{coef, 0}, {ceof2, pow}}}
}
func (this *polycule) Evaluate(x fr.Fraction) fr.Fraction {
	return fr.Mult(this.coef, fr.Pow(x, this.pow))
}
func (this Polynomial) Evaluate(x fr.Fraction) fr.Fraction {
	out := fr.FromInt(0)
	for i := 0; i < len(this.data); i++ {
		addr := this.data[i].Evaluate(x)
		out = fr.Add(out, addr)
	}
	return out
}
func (this Polynomial) FindZeros() []fr.Fraction {
	return make([]fr.Fraction, 0)
}
