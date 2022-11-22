package main

import (
	"fmt"
)

func main() {
	// for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 무한루프
	//for {
	//	fmt.Println(1)
	//}

	//range
	loc := []string{"Seoul", "Busan", "Incheon"}
	for i, name := range loc {
		fmt.Println(i, name)
	}

	li := []string{"A", "B", "C"}
	for i, _ := range li {
		fmt.Println(li[i])
	}


}
