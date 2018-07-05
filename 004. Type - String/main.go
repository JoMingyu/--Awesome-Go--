package main

func main() {
	// 문자열은 텍스트를 표현하는 데 사용되는 길이가 정해진 문자의 나열
	// Go의 문자열은 개별 바이트로 구성돼 있으며, 보통 각 문자마다 한 byte를 차지한다
	println("Hello World")
	println(`Hello World`)
	// 문자열 리터럴을 표현하기 위해서는 위처럼 큰따옴표나 back quote를 사용한다
	// 큰따옴표로 만든 문자열은 줄바꿈을 표함할 수 없고, 대신 이스케이프 문자열을 사용할 수 있다

	a := "Hello World"

	println(a[3])
	println(a + "!")
	// 인덱싱, 덧셈 기호를 통한 concat이 기본적으로 내장되어 있다
}
