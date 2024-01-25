package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amerikarno/icoApi/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AdminAuthen struct{}

func NewAdminAuthen() *AdminAuthen {
	return &AdminAuthen{}
}

func (m *AdminAuthen) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		// Extracting token from the "Bearer token" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
		}

		tokenString := &parts[1]

		// Parsing and validating the token
		access := models.AccessClaims{}
		token, err := jwt.ParseWithClaims(*tokenString, &access, func(token *jwt.Token) (interface{}, error) {
			// Check token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the key for validation
			return []byte("secretKey"), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired JWT token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Token is valid, add it to the context and call the next handler
			c.Set("user", claims)
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token")
		}
	}
}
