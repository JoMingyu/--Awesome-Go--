package main

import (
	"fmt"
)

type combiner interface {
	// 구조체가 필드의 집합이었던 것처럼, 인터페이스는 메소드의 집합으로 정의됨
	// 일반적으로 인터페이스 타입의 변수에 구조체의 주소를 assign해서 implement하는 방식으로 구현
	sum() int
}

type vertex struct {
	X, Y int
}

func (v *vertex) sum() int {
	return v.X + v.Y
}

func sum(a, b int) int {
	// Duck typing이라고 부르는, 동적 타이핑의 한 종류
	// 메소드 구현체를 만들기 위해 구조체의 주소를 인터페이스 타입에 assign하지 않고 바로 구현하는 방식
	return a + b
}

func main() {
	var a combiner

	v := vertex{3, 4}

	a = &v
	// *vertex implements combiner를 뜻함
	// 인터페이스와, 인터페이스의 특정 메소드가 정의된 구조체를 엮어주는 역할

	fmt.Println(a.sum())
	fmt.Println(sum(1, 3))
}
