package main

func main() {
	defer func() {
		data := recover()
		// recover는 panic에 대한 제어를 회복한다
		// panic과 recover는 거의 사용되지 않는다
		println(data)
	}()

	panic("PANIC!")
	// panic은 런타임 오류를 일으킨다
	// 일반적으로 프로그램의 비정상적인 조건을 처리하기 위해 error 타입의 변수를 사용하는데,
	// 프로그램을 계속 실행할 수 없는 상황에는 panic을 사용해 프로그램을 종료한다
	// 함수가 panic을 만나면 실행이 중지되고, defer 함수가 실행된 다음 제어가 반환된다
}
