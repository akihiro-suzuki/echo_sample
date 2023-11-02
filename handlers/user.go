package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

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

func Save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func Users(c echo.Context) error {
	println("handlers.Users")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
func Auth(c echo.Context) error {
	println("Auth")
	team := c.QueryParam("team")
	auth := c.QueryParam("auth")
	return c.String(http.StatusOK, "team:"+team+", auth:"+auth)
}
