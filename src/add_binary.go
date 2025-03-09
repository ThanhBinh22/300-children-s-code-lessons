package main

import (
	"fmt"
	"math/big"
)

// Bài 67
// https://leetcode.com/problems/add-binary/description/

func addBinary(a string, b string) string {
	decA := new(big.Int)
	decB := new(big.Int)
	decA.SetString(a, 2)
	decB.SetString(b, 2)

	sum := new(big.Int).Add(decA, decB)

	return sum.Text(2)
}
