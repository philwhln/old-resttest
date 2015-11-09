package resttest

import (
	"log"
	"testing"
)

func TestCentsString(t *testing.T) {
	testCentsString(1, "0.01")
	testCentsString(-1, "-0.01")
	testCentsString(12, "0.12")
	testCentsString(-12, "-0.12")
	testCentsString(123, "1.23")
	testCentsString(-123, "-1.23")
	testCentsString(1234, "12.34")
	testCentsString(-1234, "-12.34")
	testCentsString(12345, "123.45")
	testCentsString(-12345, "-123.45")
}

func testCentsString(cents Cents, expectString string) {
	if s := cents.String(); s != expectString {
		log.Fatalf("cents:%d expected:%s got:%s", cents, expectString, s)
	}
}
