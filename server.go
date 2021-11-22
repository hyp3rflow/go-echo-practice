package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	e.GET("/show", func(c echo.Context) error {
		team := c.QueryParam("team")
		member := c.QueryParam("member")
		return c.String(http.StatusOK, "team:" + team + ", member:" + member)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
