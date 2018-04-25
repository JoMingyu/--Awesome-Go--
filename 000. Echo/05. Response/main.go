package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type user struct {
	Email string
	Pw    string
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		u := user{"some@email.com", "pw"}

		c.Response().Header().Add("Cache-Control", "no-cache")
		c.SetCookie(&http.Cookie{Name: "Some", Value: "Cookie"})
		// 헤더와 쿠키 설정

		return c.JSON(http.StatusOK, u)
		// echo.Context.JSON() 함수에 구조체를 전달. Echo 측에서 해당 구조체를 JSON string으로 직렬화
	})

	e.Start(":5000")
}
