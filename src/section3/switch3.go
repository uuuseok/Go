package main

import "fmt"

func main() {
	switch a := 3; a {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		if a == 1 {
			fmt.Println(a, "a == 1")
		} else {
			fmt.Println(a, "a != 1")
		}
		fmt.Println("홀수")
	}

	// fallthrough
	// 조건에 맞는 case에 들어갔을 때, fallthrough가 있다면, 다음 case로 넘어가서 조건에 맞지 않아도 실행
	switch e := "go"; e {
	case "java":
		fmt.Println("java")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
		//fmt.Println("go2") // fallthrough 보다 뒤에 작성할 수 없음
	case "ruby":
		fmt.Println("ruby")
		fallthrough
	case "c":
		fmt.Println("c")
		//fallthrough // 마지막 case에 작성할 수 없음

		// 조건에 맞는 go case에 들어가 명령어 실행하고, fallthrough가 있기 때문에
		// 다음 case인 ruby로 가서 명령어를 실행
		// 마찬가지로 fallthrough가 있기 때문에 다음 케이스인 c로 가서 명령어를 실행

	}

}
