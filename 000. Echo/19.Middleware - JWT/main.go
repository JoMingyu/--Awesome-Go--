package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	group := e.Group("/restricted")
	// JWT 인증이 필요한 API를 별도로 관리하기 위해 group 생성

	// JWT는 JWT authentication을 제공하는 미들웨어
	// 요청의 토큰이 유효하지 않는 경우 401을 반환하고, Authorization 헤더가 존재하지 않는 경우 400을 반환
	group.Use(middleware.JWT([]byte("secret")))
	// func JWT(key interface{}) echo.MiddlewareFunc
	// 사용되는 JWTConfig(DefaultJWTConfig)는 다음과 같다
	/*
		JWTConfig{
			Skipper:       DefaultSkipper,
			SigningMethod: AlgorithmHS256,
			ContextKey:    "user",
			TokenLookup:   "header:" + echo.HeaderAuthorization,
			AuthScheme:    "Bearer",
			Claims:        jwt.MapClaims{},
		}
	*/
	// JWT()는 전달된 인자인 key를 JWTConfig 구조체에 SigningKey라는 필드로 설정한다
	// Token의 기본 prefix는 "Bearer"다

	e.GET("/token", func(c echo.Context) error {
		type binder struct {
			Username string `json:"username"`
		}

		payload := &binder{}
		if err := c.Bind(payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		token := jwt.New(jwt.SigningMethodHS256)
		// func New(method SigningMethod) *Token

		claims := token.Claims.(jwt.MapClaims)
		// map[string]interface{} 타입으로 캐스팅
		claims["identity"] = payload.Username
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		// func (t *Token) SignedString(key interface{}) (string, error)
		// 전달된 key로 토큰을 인코딩하여 반환
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	})

	group.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		identity := claims["identity"].(string)

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", identity))
	})

	e.Start(":8000")
}
