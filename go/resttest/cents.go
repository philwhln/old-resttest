package resttest

import (
	"fmt"
)

// Money value as cents (eg. $12.34 == 1234c)
type Cents int64

// Stringify cents as dollars and pennies (eg. 1234 => "12.34")
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
