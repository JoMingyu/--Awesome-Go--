package main

func divmod(a, b int) (int, int) {
	// Go의 함수는 여러 개의 값을 반환할 수 있으며
	// 함수의 시그니처 중 반환 값을 명시하는 부분에 (type, type, ...) 형태로 타입을 명시해 주면 된다
	return a / b, a % b
}

func main() {
	quotient, remainder := divmod(3, 2)
	// 반환 값을 순서대로 변수에 대입

	println(quotient, remainder)
}
