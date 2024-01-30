package common

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// DO NOT use in production
// Store secure key in vault instead
const RefreshSecret = "refresh secret"
const AccessSecret = "access secret"

// Extract JWT from gin context.
// Header: Bearer XXX
// return "" if not found
func ExtractJWT(c echo.Context) string {
	// authHeader := c.Request.Header.Get("Authorization")
	authHeader := c.Request().Header.Get("Authorization")

	return strings.TrimPrefix(authHeader, "Bearer ")
}
