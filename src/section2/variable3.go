package main

import "fmt"

func main() {
	// 짧은 선언
	// 함수 안에서만 사용(전역x), 선언 후 재할당하면 예외(에러) 발생
	// 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있음 -> 추천

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	//shortVar1 := 10  // 한 번 초기화가 되면 재할당 할 수 없음
	//Example
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
	}


	fmt.Println(shortVar1, shortVar2, shortVar3)
}