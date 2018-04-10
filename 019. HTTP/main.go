package main

import (
	"fmt"
	"net/http"
)

// HTTP Server를 구현하기 위한 빌트인 패키지. 특정 경로의 handler를 호출하는 것을 'Serve'라는 개념으로 봄

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type hello2 struct{}

func (h hello2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello2!")
}

// 요청을 핸들링하기 위해 http.Handler 인터페이스의 ServeHTTP() 메소드를 구현한 두 개의 구조체

func main() {
	var handler http.Handler
	var h hello

	handler = &h
	http.Handle("/", handler)
	// 1. http.Handler 인터페이스의 ServeHTTP() 메소드를 구현한 구조체를 사용하여 요청을 핸들링할 수 있음

	var h2 hello2
	http.Handle("/2", h2)
	// 2. Duck typing되어 ServeHTTP() 메소드가 구현된 구조체를 사용해 요청을 핸들링할 수 있음

	http.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello3!!")
	})
	// 3. HandleFunc() 함수를 이용해 익명 함수로 핸들러를 즉시 전달

	http.ListenAndServe("localhost:4000", nil)
}
