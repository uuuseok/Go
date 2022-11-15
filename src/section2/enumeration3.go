package main

import "fmt"

func main() {
	const (
		_ = iota // 0 제외
		A // 1
		B // 2
		_ // 3 제외
		C // 4
	)

	const(
		_ = (iota + 0.75) *2
		D
		E
		_
		F
	)

	fmt.Println(A,B,C)
	fmt.Println(D,E,F)

}

