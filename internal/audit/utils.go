package audit

import (
	"github.com/labstack/echo/v4"
)

// GetClientIP extracts the client IP from an Echo context.
// Uses Echo's built-in RealIP() which handles X-Forwarded-For, X-Real-IP, and RemoteAddr.
func GetClientIP(c echo.Context) string {
	ip := c.RealIP()
	if ip == "" {
		return "unknown"
	}
	return ip
}
