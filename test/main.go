package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func getClientIP(c echo.Context) error {
	// Attempt to retrieve the IP address from the X-Forwarded-For header first.
	// This header is set by proxies, load balancers, or other network devices.
	ip := c.Request().Header.Get("X-Forwarded-For")
	if ip == "" {
		// If the X-Forwarded-For header is empty, use the RemoteAddr from the request.
		// This will have the IP address and port number.
		ip = c.Request().RemoteAddr
	}

	// Send the IP address back to the client.
	return c.String(http.StatusOK, ip)
}

func main() {
	e := echo.New()

	// Define a route and use getClientIP as the handler function
	e.GET("/get-ip", getClientIP)

	// Start the Echo server on port 8080
	e.Start(":1333")
}
