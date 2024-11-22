package index

import (
	"just-html/src/shared"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/")

	count := 0

	g.GET("", func(c echo.Context) error {
		return shared.Render(c, 200, Index(count))
	})

	g.GET("counter", func(c echo.Context) error {
		count++
		return shared.Render(c, 200, Counter(count))
	})
}
