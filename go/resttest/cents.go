package resttest

import (
	"fmt"
)

type Cents int64

func (cents Cents) String() string {
	sign := ""
	if cents < 0 {
		sign = "-"
	}
	unsignedCents := absCents(cents)
	dollars := unsignedCents / 100
	pennies := unsignedCents % 100
	pennyPad := ""
	if pennies < 10 {
		pennyPad = "0"
	}
	return fmt.Sprintf("%s%d.%s%d", sign, dollars, pennyPad, pennies)
}

func absCents(c Cents) Cents {
	if c < 0 {
		return -c
	}
	return c
}
