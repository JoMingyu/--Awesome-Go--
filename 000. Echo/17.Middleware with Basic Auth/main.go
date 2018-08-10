package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// BasicAuth는 요청의 Authorization 헤더에 "basic ***" 형태의 credential을 받아
	// username과 password를 분리하고, echo.Context와 함께 handler function 호출의 인자로 전달하며 호출한다
	// ***에 들어가는 값은 "username:password" 포맷의 데이터가 base64 인코딩된 문자열이다
	// 그 예는 다음과 같다 : basic am9lOnNlY3JldGU=
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// func BasicAuth(fn BasicAuthValidator) echo.MiddlewareFunc
		// 사용되는 BasicAuthConfig(DefaultBasicAuthConfig)는 아래와 같다
		/*
			{
				Skipper: DefaultSkipper,
				Realm:   defaultRealm,
			}
		*/

		if username == "joe" && password == "secret" {
			// username과 password가 인가된 사용자로 평가된다면, true를 반환하면 된다
			return true, nil
		}
		return false, nil
		// BasicAuth 미들웨어에서 false가 반환되면 handler function을 처리하지 않는다
	}))

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8000")
}
