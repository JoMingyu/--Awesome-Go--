package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// echo는 *Echo 타입의 구조체 포인터로 시작한다

	e.GET("/", func(c echo.Context) error {
		// func (e *Echo) GET(path string, h HandlerFunc)
		// Flask에서 API 로직을 처리하는 함수를 view function이라 하듯, Go에서는 handler function이라고 관례적으로 부르고 있다
		// handler function은 '현재 요청에 대한 context'를 의미하는 echo.Context 타입의 인자를 받고, error 타입의 반환값을 가지고 있어야 한다
		// Context 인터페이스에는 요청과 응답, handler, logger 등 echo 어플리케이션 로직 구성을 위한 대부분의 메소드들이 구성되어 있다

		fmt.Println("Request received!!")

		return c.String(200, "Hello")
		// Context 인터페이스에 정의된 String(code int, s string) 메소드
		// response를 위한 메소드는 HTML, String, JSON, JSONPretty, JSONP, XML, Blob, Stream, File, NoContent 등 정말 많은데
		// Get started에도 나오는 String 메소드가 가장 기본형이다
	})

	e.Start(":8000")
	// func (e *Echo) Start(address string)
}
