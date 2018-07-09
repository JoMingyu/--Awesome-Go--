package main

func numberGenerator() func() int {
	u := 0
	return func() int {
		u++
		return u
	}
	// Go도 정적 스코핑(=lexical scoping) 언어기 때문에 클로저(함수와 해당 함수가 참조하는 비지역 변수를 포함한 지역)가 형성될 수 있다
	// lexical scoping에 의해 함수 정의 순간에 유효 범위가 잡히므로, numberGenerator가 내부 함수를 반환하고 종료되더라도
	// 내부 함수는 외부 범위인 numberGenerator 함수의 지역변수인 u에 대한 참조가 유지된다
	// 따라서 위의 경우 내부 함수와 변수 a가 클로저를 형성한다
}

func main() {
	gen := numberGenerator()

	println(gen())
	println(gen())
	println(gen())
}
