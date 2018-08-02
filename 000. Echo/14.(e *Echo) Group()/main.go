package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	group := e.Group("/pref")
	// func (e *Echo) Group(prefix string, m ...MiddlewareFunc) (g *Group)
	// 전달된 prefix를 기본 URI prefix로 가지는 Group 구조체 포인터 반환

	group.GET("/something", func(c echo.Context) error {
		// Echo 구조체와 매우 유사하게 구성되어 있음
		// 특정 API들에 한정하여 URL prefix나 미들웨어 등 특별한 속성을 한 번에 부여하기 위한 그룹핑 용도로 사용됨
		return c.NoContent(http.StatusOK)
	})
	// Flask의 Blueprint와 비슷하게, Echo 어플리케이션을 작은 단위로 쪼개는 데에 유용하게 사용될 수 있다

	e.Start(":8000")
}
