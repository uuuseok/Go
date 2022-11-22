package main

import "fmt"

func main() {
	//제어문(조건문)
	//if문은 반드시 Boolean 형으로 검사( 0,1 이 아닌 false, true )

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15이상")
	}

	if b >= 25 {
		fmt.Println("25이상")
	}

	if b == 20 {
		fmt.Println("20")
		b = 10
		fmt.Println(b)
	}

	if c := true; c {
		fmt.Println("True")
	}

	if d := 10; d < 11 {
		fmt.Println("d")
	}

}
