package algebra

import (
	"fmt"
)

func (this *Polynomial) ToString() string {
	out := ""
	for i := 0; i < len(this.data); i++ {
		out += this.data[i].coef.ToString()
		out += fmt.Sprintf("x^%d", this.data[i].pow)
	}
	return out
}
