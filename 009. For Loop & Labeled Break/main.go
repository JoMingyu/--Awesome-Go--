package main

func main() {
	sum := 0

	for i := 1; i <= 100; i++ {
		// 소괄호가 없는 걸 제외하면 알고 있던 for와 별반 다를 것 없다
		sum += i
	}

	println(sum)

	i := 1
	for i < 100 {
		// 전/후 처리를 제외하고 조건식만 있는 for loop
		// 다른 언어의 while 루프와 같이 쓰이도록 함
		sum += i

		break
		// continue와 break는 다른 언어와 동일하게 사용 가능
	}

	println(sum)

L1:
	for {
		for {
			break L1
			// 특정 level의 for loop를 바로 break할 수 있음
		}
	}
}
