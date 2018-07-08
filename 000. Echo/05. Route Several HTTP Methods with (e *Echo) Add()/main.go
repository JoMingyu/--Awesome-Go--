package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Echo 구조체에서 라우팅을 돕는 몇가지 메소드들은, 내부적으로 (e *Echo) Add() 메소드를 호출한다
	e.Add(echo.GET, "/", func(c echo.Context) error {
		// func (e *Echo) Add(method, path string, handler HandlerFunc)
		// echo 패키지에 정의되어 있는, HTTP 메소드를 나타내는 상수를 전달한다
		return c.NoContent(http.StatusOK)
	})

	// 위에서 사용한 HTTP 메소드 상수들은, 상수 이름과 동일한 문자열을 담고 있다. 따라서 아래와 같은 표현이 가능하다
	e.Add("GET", "/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// 라우팅 방식은 취향으로 두는 분위기인 것 같다
	e.Start(":8000")
}
