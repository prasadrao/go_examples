package main

import "fmt"

func reverse(s *[6]int) *[6]int {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func main() {
	var a *[6]int
	a = &[6]int{9, 8, 7, 6, 5, 4}
	fmt.Println(*reverse(a))
}
