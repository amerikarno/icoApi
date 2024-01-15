package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegularMiddleware struct{}

func NewRegularMiddleware() *RegularMiddleware {
	return &RegularMiddleware{}
}

func (RegularMiddleware) Authen(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization token missing"})
		}

		if tokenString[7:] == "fda-authen-key" {
			return next(c)
		}

		log.Printf("invalid token: %v", tokenString)

		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}
}
