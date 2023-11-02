package main

import (
	"main/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	Hoge()
	controllers.Test()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! From Go Echo3")
	})
	e.GET("/users/:id", controllers.GetUser)
	e.GET("/show", controllers.Show)
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
