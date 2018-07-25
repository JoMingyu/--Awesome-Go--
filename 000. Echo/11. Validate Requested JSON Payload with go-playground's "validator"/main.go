package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	// go get gopkg.in/go-playground/validator.v9
	// Echo 공식 문서에서는 JSON payload의 validator를 위해 별도의 validator 라이브러리를 이용하는 것을 권고하고 있다
)

type binder struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
	// JSON payload binding을 위한 구조체에 validate 태그를 추가하자
	// required, length, max, min, gt, gte 등의 validation rule들을 지원하고 있으며, 이들은 라이브러리 공식 문서에 설명되어 있다
	// https://godoc.org/gopkg.in/go-playground/validator.v9#hdr-Required
}

var validate = validator.New()

// New() *Validate 함수를 이용해 Validate 구조체 포인터 변수를 정의

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		payload := &binder{}

		if err := c.Bind(payload); err != nil {
			return err
		}

		if err := validate.Struct(payload); err != nil {
			// (v *Validate) Struct(s interface{}) error 함수는 인자로 전달된 구조체에 대해 validation하며
			// 해당 구조체가 태그에 정의된 rule에 맞지 않는 경우 에러를 반환한다
			return c.NoContent(http.StatusBadRequest)
		}

		fmt.Println(payload.Name, payload.Age)

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", payload.Name))
	})
	// 번거롭지만 어쩔 수 없다. 빠르니까...

	e.Start(":8000")
}
