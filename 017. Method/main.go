package main

import (
	"fmt"
)

type vertex struct {
	X, Y int
}

// Go에는 클래스가 없지만, 메소드를 구조체에 붙일 수 있음

func (v *vertex) sum() int {
	// 구조체에 메소드를 붙이려면, func 키워드와 메소드의 이름 사이에 인자 형태로 method receiver를 넣어 주어야 함
	// 리시버는 위와 같은 포인터 리시버 와 값 타입 리시버로 나뉨
	// 메소드는 구조체 뿐만 아니라 아무 타입에나 붙을 수 있음. 그러나 다른 패키지에 있는 타입이나 기본 타입들에 메소드를 붙이는 것은 불가능
	return v.X + v.Y
}

func main() {
	v := vertex{1, 2}

	fmt.Println(v.sum())
}
