package main

import "fmt"

func main() {

	//레이블 지정
Loop1:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == 1 && y == 1 {
				break Loop1
			}
			fmt.Println(x, y)
		}
	}
	fmt.Println("end")

	for k := 0; k < 3; k++ {
	Loop2:
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if k == 1 && i == 1 && j == 1 {
					break Loop2
				}
				fmt.Println(k, i, j)
			}
		}
	}
	fmt.Println("END")

	//반복문에서 continue를 만났을 때, 다음 loop로 넘어감
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//break 또는 continue 뒤에 레이블을 명시적으로 적어주면, 해당 레이블로 즉시 이동

Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue Loop
			}
			fmt.Println(i, j)
		}
	}

	for i := 0; i < 3; i++ {
	Loop4:
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if i == 1 && j == 1 && k == 1 {
					continue Loop4
				}
				fmt.Println(i, j, k)
			}
		}
	}

}
