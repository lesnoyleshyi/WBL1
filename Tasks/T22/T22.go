package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(9223372036854775807)
	b := big.NewInt(9223372036854775805)

	mul := big.NewInt(0).Mul(a, b)
	div := big.NewInt(0).Div(a, b)
	sum := big.NewInt(0).Add(a, b)
	diff := big.NewInt(0).Sub(a, b)

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("x * y = %d\n", mul)
	fmt.Printf("x / y = %d\n", div)
	fmt.Printf("x + y = %d\n", sum)
	fmt.Printf("x - y = %d\n", diff)
}
