package main

import "fmt"

func main() {
	const (
		Jan = iota + 1 //0+1
		Feb            // 1+1
		Mar            // 2+1
		Apr            // 3+1
		May            // 4+1
		Jun            // 5+1
	)

	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)
}
