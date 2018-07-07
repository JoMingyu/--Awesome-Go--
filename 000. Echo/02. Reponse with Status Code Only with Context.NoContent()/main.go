package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		// (c *echo.Context) String(code int, s string) 메소드는 s로 전달된 문자열과 함께 response하는 역할을 하는데
		// 여기에 빈 문자열을 전달하더라도 response는 string 타입의 empty string으로 전달된다
		// Context 인터페이스의 NoContent(code int) 메소드를 이용해 response하면, response header에 content type이 명시되지 않아 완전히 빈 response body를 표현할 수 있다
		return c.NoContent(200)
		// HTTP status code 204 No content와 연관이 1도 없음에 주의하자
	})

	e.Start(":8000")
}
