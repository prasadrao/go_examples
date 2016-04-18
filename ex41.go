package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.

func PopCount(x byte) int {
	return int(pc[x>>(0*8)] +
		pc[x>>(1*8)] +
		pc[x>>(2*8)] +
		pc[x>>(3*8)] +
		pc[x>>(4*8)] +
		pc[x>>(5*8)] +
		pc[x>>(6*8)] +
		pc[x>>(7*8)])
}

func count_ones(x, y [32]byte) int {
	var count int
	for i, j := range x {
		count = count + PopCount(j^y[i])
	}
	return count
}
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%X\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(count_ones(c1, c2))

	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8
	//
	// false
	//[32]uint8
}
