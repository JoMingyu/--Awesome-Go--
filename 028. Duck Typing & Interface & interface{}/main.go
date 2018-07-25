package main

type shape interface {
	// 필드의 집합인 구조체와 다르게, 인터페이스는 메소드의 집합. Go의 인터페이스는 duck typing 방식으로 작동한다
	// -> square와 rectangle 구조체 모두 shape 인터페이스를 구현한다는 명시적 키워드가 없지만,
	// 단지 shape 인터페이스의 메소드를 구현한 것으로 이들은 shape 인터페이스의 타입으로 사용될 수 있다는 것이다
	// "어떤 새가 오리인지 검사하지 말고, 오리처럼 걷고, 헤엄치고, 꽥꽥거리는 소리를 낸다면 그 새를 오리라고 부를 것이다"라는 것이 duck typing의 이론이다
	area() float64
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type rectangle struct {
	r float64
}

func (r rectangle) area() float64 {
	return 3.14 * r.r * r.r
}

func showArea(s shape) {
	println(s.area())
}

func main() {
	s := square{3}
	showArea(s)

	// 인터페이스를 활용하는 예로 '빈 인터페이스(interface {})'가 있다
	// 흔히 인터페이스 타입으로도 불리는데, Go의 모든 타입은 0개 이상의 메소드를 구현하므로 '모든 타입'을 나타내기 위해 사용한다
	// 즉 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있으며,
	// 다른 언어에서 흔히 일컫는 Dynamic Type(Java의 object, C 계열의 void*)이라 볼 수 있다
	m4 := map[string]interface{}{"a": "hello", "b": 3}
	_ = m4
	// value type 자리에 interface{}를 두면, 여러 타입의 value를 사용할 수 있다
}
