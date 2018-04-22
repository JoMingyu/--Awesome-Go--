package main

import (
	"net/http"

	// 고성능, 확장 가능한 고루틴 기반의 웹 프레임워크
	// https://github.com/labstack/echo
	// https://echo.labstack.com/
	// https://echo.labstack.com/guide

	// go get github.com/labstack/echo
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// Echo 인스턴스
	// Echo 구조체의 포인터를 반환하며, e *Echo = echo.New()와 같음

	e.GET("/", hello)
	// Routes

	e.GET("/2", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!")
	})

	e.Start(":5000")
	// Start server
}

// Handler
func hello(c echo.Context) error {
	// error를 리턴하고, echo.Context를 파라미터로 받는 함수를 통해 handler를 정의
	return c.String(http.StatusOK, "Hello, World!")
}
