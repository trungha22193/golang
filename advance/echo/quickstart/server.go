package main 
import (
	"net/http"
	"github.com/labstack/echo"
	)

func main() {
	e := echo.New()

	//Form application/x-www-form-urlencoded
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Echo Framework v1.0")
		})
	e.GET("/users/:id", getUserDetail)
	e.POST("/users", postUser)
	e.PUT("/users/:id", putUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1234"))
}

func getUserDetail(c echo.Context) error {
	id := c.Param("id")
	team := c.QueryParam("team")
	return c.String(http.StatusOK, "Get detail ID := " + id + "\nTeam := " + team)
}

func postUser(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, "Post Name:" + name)
}

func putUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "PUT ID := " + id)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Delete Id:" + id)
}