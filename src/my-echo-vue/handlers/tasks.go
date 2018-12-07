package handlers
import (
	"database/sql"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	)

type H map[string]interface{}

// They all take a db connection as an argument
// Proper handler used by the Echo router, the function needs to implement the Echo.HandlerFunc interface
// We accomplish this by returning an anonymous function that matches the interface signature


func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"tasks": "",
			})
	}
} 

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": 123,
			})
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"deleted": id,
			})
	}
} 