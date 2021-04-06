package algorithm

import (
	"fmt"
	"testing"
)

func TestGcd(t *testing.T) {
	gcd := Gcd(20, 8)
	fmt.Println("20 and 8, get gcd = ", gcd)
}
