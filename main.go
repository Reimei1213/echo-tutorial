package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/users", saveUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}