package main

import "fmt"

func main() {
	//선언
	var a int
	var b string
	var c bool
	var d, e float32
	//선언 및 초기화
	var f,g,h int = 1,2,3
	var i float32 = 11.4
	var j string = "hi!"
	// 선언 동시 초기화
	var k = 4.75
	var l = "Hello"
	var m,n = true, 7.6

	// 선언이 된 변수이기 때문에, 값만 초기화
	a = 4
	b = "bbbbb"
	c = true
	d,e = 3.5,6.6

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)


}