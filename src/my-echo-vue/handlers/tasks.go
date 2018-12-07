package handlers
import (
	"database/sql"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"my-echo-vue/models"
	)

type H map[string]interface{}

// They all take a db connection as an argument
// Proper handler used by the Echo router, the function needs to implement the Echo.HandlerFunc interface
// We accomplish this by returning an anonymous function that matches the interface signature


func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
} 

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ############# Level 1 #############
		// return c.JSON(http.StatusCreated, H{
		// 	"created": 123,
		// 	})
		// ############# Level 1 ending #############

		var task models.Task
		c.Bind(&task)
		id, err := models. PutTask(db, task.Name)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusCreated, H{
				"created":id,
				})
		}

	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ############# Level 1 #############
		// id, _ := strconv.Atoi(c.Param("id"))
		// return c.JSON(http.StatusOK, H{
		// 	"deleted": id,
		// 	})
		// ############# Level 1 ending #############

		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
				})
		}
	}
} 