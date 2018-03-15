package main

import (
	"fmt"
)

type vertex struct {
	// 구조체는 필드들의 조합이며, type 선언으로 구조체의 이름을 지정할 수 있음
	// 사용자 정의 데이터 타입을 선언하기 위함
	X int
	Y int
}

func main() {
	v := vertex{1, 2}
	// 새로운 Vertex 구조체 생성
	// Go에서 연속된 데이터를 다루기 위해선 {}를 사용함
	// 나중에 slice나 map을 배울 때도 사용하므로 익숙해져야 함

	v.X = 3
	// 구조체에 속한 필드는 dot(.)으로 접근

	fmt.Println(v)

	w := vertex{Y: 10}
	// 구조체 리터럴은 필드와 값을 나열해 구조체를 할당하는 방법

	fmt.Println(w)

	x := new(vertex)
	// new(T)는 모든 필드가 zero value인 T 타입의 포인터를 반환
	// var x *vertex = new(vertex)와 동일
	fmt.Println(x)
}
