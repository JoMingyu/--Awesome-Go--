package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type user struct {
	// echo.Context.Bind()를 이용해 Request body를 Go의 구조체로 bind할 수 있음
	// 1. 요청 payload의 타입(json, form, query 등)에 관계 없이, 필드에 태깅만 해 두고 구조체로 바로 바인딩
	// 2. 요청의 JSON Payload를 불러와 구조체로 역직렬화
	// 3. Form validator로서의 역할 등이 가능
	Email string `json:"email" form:"email" query:"email"`
	Pw    string `json:"pw" form:"pw" query:"pw"`
	// json, form, query에 대해 tag(type:keyName)
	// 해당 구조체의 Email, Pw 모두 요청의 json, form, query에 대응하여 바인딩됨
}

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		if err := c.Bind(new(user)); err != nil {
			// c.Bind()로 zero value인 구조체를 전달하면
			// Echo에서 해당 구조체의 필드들을 태그 기반으로 리플렉션 접근하여 필드들을 채워줌
			return err
		}

		return c.String(http.StatusOK, "Hello!")
	})

	e.Start(":5000")
}
