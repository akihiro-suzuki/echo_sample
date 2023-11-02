package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// パスパラメータ
// e.GET("/users/:id", getUser)
func GetUser(c echo.Context) error {
	// User ID from path `users/:id` ex. users/name
	name := c.Param("id")
	return c.String(http.StatusOK, name)
}

// クエリパラメータ
// e.GET("/show", show)
func Show(c echo.Context) error {
	// Get team and member from the query string
	// show?team=x-men&member=wolverine
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
