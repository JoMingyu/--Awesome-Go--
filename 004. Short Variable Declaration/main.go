package main

func main() {
	// 아래는 기본적인 변수 선언
	var a int = 3

	// 짧게
	var b = 3

	// Go는 타입 추론을 통한 변수 선언 문법이 하나 더 있다
	c := 3

	println(a, b, c)

	d, e := 3, 5

	println(d, e)
}
