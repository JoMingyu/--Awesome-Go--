package main

type position struct {
	// type [name] struct
	// 이름이 지정된(name) 필드들로 이루어진 타입
	x int
	y int
	// 각 필드들은 이름과 타입을 가진다
}

func main() {
	var p position
	// 구조체는 값 타입이기 때문에, nil로 초기화되지 않고, 필드들이 zero value로 채워진 구조체로 초기화됨
	println(p.x, p.y)
	// 필드 접근

	p2 := position{x: 1, y: -3}
	_ = p2
	// 필드에 대해 값을 할당하며 변수 생성

	p3 := position{1, 3}
	_ = p3
	// 필드에 순서대로 할당하는 경우, 필드명 생략

	p4 := position{}
	_ = p4
	// zero value로 초기화하며 생성
	// var p4 poisition과 동일

	p5 := new(position)
	_ = p5
	// 모든 필드에 대해 메모리가 할당되고, 각 필드는 zero value로 설정된 후 포인터(*position 타입)가 반환
	// var p5 *position, p4 := &position{}과 동일한 구문
}
