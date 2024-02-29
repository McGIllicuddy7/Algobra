package algebra

import (
	"errors"
	"math/cmplx"
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
	new_data := make([]polycule, 0)
	powers := make([]int, 0)
	for i := 0; i < len(this.data); i++ {
		if !slice_contains(powers, this.data[i].pow) && this.data[i].coef.ToInt() != 0 {
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
	this.data = this.data[:len(powers)]
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
func (this *polycule) evaluate(x fr.Fraction) fr.Fraction {
	return fr.Mult(this.coef, fr.Pow(x, this.pow))
}
func (this Polynomial) Evaluate(x fr.Fraction) fr.Fraction {
	out := fr.FromInt(0)
	for i := 0; i < len(this.data); i++ {
		addr := this.data[i].evaluate(x)
		out = fr.Add(out, addr)
	}
	return out
}
func (this *polycule) evaluateComplex(x complex128) complex128 {
	v := this.coef.ToComplex()
	p := complex(float64(this.pow), 0)
	y := cmplx.Pow(x, p)
	return v * y
}
func (this Polynomial) EvaluateComplex(x complex128) complex128 {
	out := complex128(0)
	for i := 0; i < len(this.data); i++ {
		out += this.data[i].evaluateComplex(x)
	}
	return out
}
func (this *polycule) Derivitive() {
	if this.pow == 0 {
		this.coef = fr.FromInt(0)
		this.pow = 0
		return
	}
	if this.pow == 1 {
		this.pow = 0
		return
	}
	this.coef = fr.Scale(this.coef, int64(this.pow))
	this.pow--
}
func (this *polycule) Integral() error {
	if this.pow == -1 {
		return errors.New("error unspported function \"natural log\"")
	}
	this.coef = fr.Scale(this.coef, fr.Recip(fr.FromInt(this.pow+1)).ToInt())
	this.pow++
	return nil
}
func (this Polynomial) Derivitive() Polynomial {
	out := this.Clone()
	for i := 0; i < len(this.data); i++ {
		out.data[i].Derivitive()
	}
	this.compress()
	return out
}
func (this Polynomial) Integral() (error, Polynomial) {
	out := this.Clone()
	for i := 0; i < len(this.data); i++ {
		err := out.data[i].Integral()
		if err != nil {
			return err, NewPoly(fr.FromInt(0), 0)
		}
	}
	return error(nil), out
}
func (this Polynomial) MinMaxPowers() (int, int) {
	lowest := this.data[0].pow
	highest := this.data[0].pow
	for i := 0; i < len(this.data); i++ {
		if this.data[i].pow > highest {
			highest = this.data[i].pow
		}
		if this.data[i].pow < lowest {
			lowest = this.data[i].pow
		}
	}
	return lowest, highest
}
func (this Polynomial) GetPowerCoefficient(power int) fr.Fraction {
	for i := 0; i < len(this.data); i++ {
		if this.data[i].pow == power {
			return this.data[i].coef
		}
	}
	return fr.FromInt(0)
}
func (this Polynomial) FindZero(seed complex128) complex128 {
	der := this.Derivitive()
	value := seed
	failsafe := 0
restart:
	for i := 0; i < 10000; i++ {
		current := this.EvaluateComplex(value)
		delta := der.EvaluateComplex(value)
		value -= current / delta
		if i%1000 == 0 {
			//cstr := utils.FormatComplex(current)
			//vstr := utils.FormatComplex(value)
			//fmt.Printf("current:%s value: %s\n", cstr, vstr)
		}
	}
	if failsafe < 10 {
		failsafe++
		value = utils.RandomComplex()
		goto restart
	}
	return value
}
func cmplxContains(slice []complex128, v complex128) bool {
	for i := 0; i < len(slice); i++ {
		if utils.ComplexNearlyEqual(slice[i], v) {
			return true
		}
	}
	return false
}
func (this Polynomial) FindZeros() []complex128 {
	minp, maxp := this.MinMaxPowers()
	utils.SortInplace[polycule](this.data, polycule_cmp)
	if minp >= 0 && maxp == 2 {
		a := this.GetPowerCoefficient(2).ToComplex()
		b := this.GetPowerCoefficient(1).ToComplex()
		c := this.GetPowerCoefficient(0).ToComplex()
		out := make([]complex128, 2)
		out[0] = (-b - cmplx.Sqrt(b*b-4*a*c)) / (2 * a)
		out[1] = (-b + cmplx.Sqrt(b*b-4*a*c)) / (2 * a)
		return out
	}
	zeros := make([]complex128, 0)
	num_zeros := maxp - minp
	for len(zeros) < num_zeros {
		tmp := this.FindZero(utils.RandomComplex())
		if !cmplxContains(zeros, tmp) {
			zeros = append(zeros, tmp)
		}
	}
	return zeros
}
