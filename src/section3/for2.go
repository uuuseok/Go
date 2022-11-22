package main

import "fmt"

func main() {
	// 'for' example 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	// 'for' example 2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		//j := i++ //Go에서 후치연산은 반환 값이 없음.
	}
	fmt.Println(sum2)

	// while과 유사
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println(sum3)

	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println(i, j)
	}

}
