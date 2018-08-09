package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	router := e.Router()
	// func (e *Echo) Router() *Router
	// 이름과 어울리게, 라우팅을 위한 구조체. Add와 Find만을 메소드로 가지고 있다

	router.Add(echo.GET, "/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8000")
}
