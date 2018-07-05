package main

// go 파일은 .go를 확장자로 하며 주석은 이처럼 slash 두개를 사용하거나
/*
	slash와 asterisk를 조합해 범위 주석을 표현할 수 있다
*/

// 패키지는 Go에서 코드를 조직화하고 재사용하는 수단이므로, 모든 Go 프로그램은 반드시 패키지 선언으로 시작해야 한다
// entry point 용도의 go 파일의 경우 패키지 이름은 main으로 두도록 한다

func main() {
	// 함수는 func 키워드로 선언하며, 해당 함수(main 패키지의 main 함수)가 Go의 entry point다
	println("Hi!")
	print("Hello.")
	// 전역 네임스페이스에 빌트인 함수들을 제공한다
}
