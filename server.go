package main

import (
	"echo_sample/handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group level middleware
	// hot to enable basic auth?

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! From Go Echo3")
	})
	e.GET("/users/:id", handlers.GetUser)
	e.GET("/show", handlers.Show)
	e.POST("/save", handlers.Save)
	g.GET("/users", handlers.Auth, track)

	e.Logger.Fatal(e.Start(":80"))
}

// パスパラメータ
// e.GET("/users/:id", getUser)
// func getUser(c echo.Context) error {
// 	// User ID from path `users/:id` ex. users/name
// 	name := c.Param("id")
// 	return c.String(http.StatusOK, name)
// }

// // クエリパラメータ
// // e.GET("/show", show)
// func show(c echo.Context) error {
// 	// Get team and member from the query string
// 	// show?team=x-men&member=wolverine
// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")
// 	return c.String(http.StatusOK, "team:"+team+", member:"+member)
// }
