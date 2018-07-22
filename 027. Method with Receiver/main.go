package main

type square struct {
	length int
}

func (s *square) area() int {
	// Go는 메소드를 지원하며, func 키워드와 함수명 사이에 ([name] [type]) 형태의 receiver를 추가하면 된다
	// 포인터 타입으로 정의해 두면, 런타임에서 메소드가 호출될 때 구조체가 포인터 타입으로 자동으로 변환되어 전달된다
	// 이런 식으로 함수를 생성하면, receiver 타입의 구조체에서 .(dot) 연산자를 이용해 호출할 수 있다
	return s.length * s.length
}

func main() {
	s := square{5}
	println(s.area())
}
