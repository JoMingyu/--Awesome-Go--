package main

func printHello() {
	// 0개의 input, 0개의 return
	println("Hello")
}

func printValue(value int) {
	// 1개의 input, 0개의 return
	println(value)
}

func returnZero() int {
	// 0개의 input, 1개의 return
	return 0
}

func sum(a, b int) int {
	// 2개의 input, 1개의 return
	return a + b
}

func main() {
	// 함수는 0개 이상의 input을 받아, 로직을 수행한 후 0개 이상의 결과를 반환하는 독립적인 코드 블럭
	printHello()
	printValue(3)
	println(returnZero())
	println(sum(1, 3))
}
