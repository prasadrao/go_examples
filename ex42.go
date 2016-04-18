package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var arg []string
	if len(os.Args[1:]) == 0 {
		c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))
		fmt.Printf("%X\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
		return
	} else {
		arg = os.Args[1:]
	}
	if arg[0] == "SHA384" {
		c1 := sha512.Sum384([]byte("x"))
		c2 := sha512.Sum384([]byte("X"))
		fmt.Printf("%X\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	}
	if arg[0] == "SHA512" {
		c1 := sha512.Sum512([]byte("x"))
		c2 := sha512.Sum512([]byte("X"))
		fmt.Printf("%X\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	}
}
