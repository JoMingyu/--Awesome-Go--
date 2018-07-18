package main

func main() {
	first := func() {
		println("first")
	}

	second := func() {
		println("second")
	}
	// Go에는 defer라는 특별한 statement가 있다
	defer second()
	// defer [function call]
	first()

	// 함수가 실행을 완료하고 제어가 반환되기 직전에 defer에 해당하는 함수를 실행한다
	// 위의 경우, second()를 defer하고 first()를 실행했으므로 first() -> second() 순서로 함수가 호출된 후 main 함수가 종료된다

	// 1. 함수에 반환 지역이 여러 개 있더라도, 함수가 종료되면 무조건 호출될 것이며
	// 2. 런타임에 문제가 생기더라도 defer는 실행된다
	// 3. 가독성에도 도움된다(변수를 생성하고, 바로 밑에 해당 변수를 정리하는 close()같은 구문을 defer)
}
