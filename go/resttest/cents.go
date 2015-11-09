package resttest

import (
	"fmt"
)

type Cents int64

func (cents Cents) String() string {
	sign := ""
	dollars := cents / 100
	pennies := cents % 100
	if cents < 0 {
		sign = "-"
	}
	return fmt.Sprintf("%s%d.%d", sign, dollars, pennies)
}
