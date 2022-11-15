package main

import "fmt"

func main() {
	// 상수
	// const 선언과 동시에 초기화 돼야 함, 한 번 선언 후 값 변경 금지, 고정된 값 관리 용

	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	const d = 34.6
	const e = false


	//const f = getValue()  // 함수의 리턴값을 상수의 값으로 초기화 할 수 없음


	/*
		에러 발생: 상수는 선언과 동시에 초기화 돼야 함
		const g string
		g = "Test3"
	*/


	fmt.Println(a,b,c,d,e)


}