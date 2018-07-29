package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusNotFound)
		// func NewHTTPError(code int, message ...interface{}) *HTTPError
		// HTTP Error code(400번 이상)를 반환해야 하는 경우엔, 해당 함수를 통해서도 response 가능하다
		// Non-error code들과 대조되므로 가독성 면에서 유리할 가능성이 있다(개발자 따라 케바케)
	})

	e.Start(":8000")
}
