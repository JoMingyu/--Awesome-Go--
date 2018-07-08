package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.DELETE("/", func(c echo.Context) error {
		// func (e *Echo) DELETE(path string, h HandlerFunc)
		return c.NoContent(http.StatusOK)
	})

	e.POST("/", func(c echo.Context) error {
		// func (e *Echo) POST(path string, h HandlerFunc)
		// 이처럼 echo 구조체의 메소드로 표준 HTTP 메소드들을 대응하여 라우팅할 수 있도록 지원하고 있다
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8000")
}
