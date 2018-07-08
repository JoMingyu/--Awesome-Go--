package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
		// Go에서 HTTP status code는 정수 리터럴을 사용하지 않고, net/http에 정의된 상수를 사용하는 관례가 있다
	})

	e.Start(":8000")
}
