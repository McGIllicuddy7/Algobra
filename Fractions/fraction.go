package fractions

import (
	"fmt"
	"math"
)

type Fraction struct {
	num int
	den int
}

func NewFrac(num int, denum int) Fraction {
	out := Fraction{num, denum}
	out.den = denum
	out.num = num
	out.simplify()
	return out
}
func gcf(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcf(b, a%b)
}
func (this *Fraction) simplify() {
	if this.num == 0 {
		this.den = 1
		return
	}
	if this.den < 0 {
		this.num *= -1
		this.den *= -1
	}
	gc := gcf(this.num, this.den)
	this.num /= gc
	this.den /= gc
}
func (f Fraction) ToString() string {
	if f.den == 1 {
		return fmt.Sprintf("%d", f.num)
	}
	return fmt.Sprintf("%d/%d", f.num, f.den)
}
func Mult(f0 Fraction, f1 Fraction) Fraction {
	out := Fraction{(f0.num * f1.num), (f0.den * f1.den)}
	out.simplify()
	return out
}
func Add(f0 Fraction, f1 Fraction) Fraction {
	out := Fraction{(f0.num*f1.den + f1.num*f0.den), (f0.den * f1.den)}
	out.simplify()
	return out
}
func Scale(f0 Fraction, scalar int) Fraction {
	out := Fraction{(f0.num * scalar), (f0.den)}
	out.simplify()
	return out
}
func Sub(f0 Fraction, f1 Fraction) Fraction {
	out := Add(f0, Scale(f1, -1))
	return out
}
func Recip(f0 Fraction) Fraction {
	return Fraction{f0.den, f0.num}
}
func Equals(f0 Fraction, f1 Fraction) bool {
	return f0.num == f1.num && f1.den == f0.den
}
func (this Fraction) ToFloat() float64 {
	return float64(this.num) / float64(this.den)
}
func (this Fraction) ToComplex() complex128 {
	return complex(this.ToFloat(), 0)
}
func (this Fraction) ToInt() int {
	return this.num / this.den
}
func FromFloat(v float64) Fraction {
	f := v
	count := 1
	for f != math.Floor(f) {
		f *= 10
		count *= 10
	}
	return NewFrac(int(f), count)
}
func FromInt(v int) Fraction {
	return NewFrac(v, 1)
}
func Pow(frac Fraction, power int) Fraction {
	out := FromInt(1)
	pow := power
	frc := frac
	if pow < 0 {
		frc = Recip(frc)
		pow *= -1
	}
	for i := 0; i < power; i++ {
		out = Mult(out, frc)
	}
	return out
}
