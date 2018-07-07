package main

func main() {
	if false {
		// For문과 비슷하게, 조건 표현을 위한 소괄호는 사용하지 않고 전체적인 구조는 C/Java와 비슷함
		println("false")
	} else if false {
		println("false!!")
	} else {
		println("else")
	}

	if val := 1 * 2; val == 2 {
		// for의 초기화 구문처럼, if에서도 조건문 앞에 statement를 실행할 수 있음
		// 해당 실행문을 통해 선언된 변수는 if 블럭 안쪽 scope나 동일한 level의 else 블럭에서만 사용할 수 있다
		// Go에서는 이런 특징을 응용하는 사례들이 매우 많다
		println("val is 2")
	}

	// if "abc" { }
	// Go는 타입들에 대한 bool 평가가 별도로 없기 때문에, expression 자체가 boolean이어야 한다
}
