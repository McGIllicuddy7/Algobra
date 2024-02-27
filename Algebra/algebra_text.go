package algebra

import (
	"fmt"
	"matrix/utils"
	"strconv"
)

func (this *Polynomial) ToString() string {
	out := ""
	for i := 0; i < len(this.data); i++ {
		if imag(this.data[i].coef) != 0 {
			out += "("
		}
		if this.data[i].coef != 1 || this.data[i].pow == 0 {
			out += utils.FormatComplex(this.data[i].coef)
		}
		if imag(this.data[i].coef) != 0 {
			out += ")"
		}
		if this.data[i].pow != 0 {
			out += "x"
			if this.data[i].pow != 1 {
				out += "^"
				out += fmt.Sprintf("%d", this.data[i].pow)
			}
		}
		if i < len(this.data)-1 {
			if imag(this.data[i+1].coef) != 0 || real(this.data[i+1].coef) > 0 {
				out += "+"
			} else {
				out += "-"
			}
		}
	}
	return out
}

func trim_left(str *string) {
	for (*str) == " " && len(*str) > 0 {
		*str = (*str)[1:]
	}
}
func handle_power(str *string) int {
	right := ""
	rightv := 1
	if len(*str) > 0 {
		if (*str)[0] == '^' {
			*str = (*str)[1:]
			if len(*str) > 0 {
				minus := false
				if (*str)[0] == '-' {
					minus = true
					*str = (*str)[1:]
				}
				for (*str)[0] != '+' && (*str)[0] != '-' {
					right += string((*str)[0])
					*str = (*str)[1:]
					if len(*str) < 1 {
						break
					}
				}
				tmp, _ := strconv.ParseInt(right, 10, 32)
				if minus {
					tmp *= -1
				}
				rightv = int(tmp)
			}
		}
	}
	return rightv
}
func parsePolycule(str *string) polycule {
	var out polycule
	trim_left(str)
	left := ""
	leftv := complex128(1)
	if (*str)[0] == '(' {
		for (*str)[0] != ')' && len(*str) > 0 {
			left += string((*str)[0])
		}
		left = left[1:]
		left = left[:len(left)-1]
	} else {
		for true {
			c := (*str)[0]
			if c == '+' || c == '-' || c == 'x' {
				break
			}
			left += string(c)
			*str = (*str)[1:]
			if len(*str) < 1 {
				break
			}
		}
	}
	if len(left) == 0 {
		left = "1"
	}
	leftv, _ = strconv.ParseComplex(left, 64)
	rightv := 1
	if len(*str) < 1 {
		return polycule{leftv, 0}
	}
	if (*str)[0] == 'x' {
		*str = (*str)[1:]
		rightv = handle_power(str)
	} else {
		rightv = 0
	}
	//("left:", leftv)
	//println(leftv, rightv)
	out.coef = leftv
	out.pow = rightv
	return out
}
func ParsePolynomial(str string) Polynomial {
	tmp := str
	out := Polynomial{make([]polycule, 0)}
	for len(tmp) > 0 {
		//(tmp)
		sub := false
		if tmp[0] == '-' {
			sub = true
			tmp = tmp[1:]
		} else if tmp[0] == '+' {
			tmp = tmp[1:]
		}
		value := parsePolycule(&tmp)
		if sub {
			value.coef *= -1
		}
		out.data = append(out.data, value)
	}
	return out
}
