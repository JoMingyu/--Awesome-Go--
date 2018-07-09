package main

func sum(a, b int, callback func(int)) {
	// Go의 함수는 일급 객체로서, 변수에 할당하고 값으로 전달할 수 있다
	callback(a + b)
	// 인자로 받은 함수에 a + b를 전달하며 호출
}

func callback(result int) {
	println(result)
}

func main() {
	s := sum

	s(1, 3, callback)
	// 인자에 a와 b에 각각 1과 3, 인자 callback에 callback이라는 함수를 전달
	// 제어는 main() -> sum() -> callback() 순서로 흘러간다

	inner := func() {

	}
	_ = inner
	// 단, Go에서 inner function에는 위처럼 함수 선언식만을 사용할 수 있다
}
