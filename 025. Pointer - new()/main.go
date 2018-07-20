package main

func main() {
	swap := func(a *int, b *int) {
		*a, *b = *b, *a
	}

	a := new(int)
	b := new(int)
	// new(type)은 type에 해당하는 충분한 메모리를 할당하고, zero value로 초기화한 후 그에 대한 포인터를 반환한다
	// 별도의 초기화 없이 포인터형 변수를 선언(var a *int)하면 nil로 초기화되므로, zero value로 초기화된 포인터 변수를 선언하기 위해 사용하기 좋다
	*a = 3
	*b = 5

	swap(a, b)
	println(*a, *b)
	// 빌트인 타입에 대한 포인터는 사용되는 일이 드물지만 구조체와 함께 사용할 때 꽤 유용하다
}
