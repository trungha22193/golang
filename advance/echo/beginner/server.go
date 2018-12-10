package main 
import (
	"github.com/labstack/echo"
	//_ "github.com/labstack/echo/engine/stardard"
	_ "fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"			// --> Must be here. If not, show ERROR
	"my-echo-vue/handlers"
	"go/build"
	)

func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	// ############# Level 1 #############
	// e.GET("/tasks", func(c echo.Context) error {
	// 	return c.JSON(200, "GET Tasks")
	// 	})
	// e.PUT("/tasks", func(c echo.Context) error {
	// 	return c.JSON(201, "PUT Tasks")
	// 	})
	// e.DELETE("/tasks/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	return c.JSON(200, "DELETE task id:=" + id)
	// 	})

	// ############# Level 2 #############
	gopath := build.Default.GOPATH
	e.File("/", gopath + "/src/my-echo-vue/public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
    e.PUT("/tasks", handlers.PutTask(db))
    e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// e.Run(standard.New(":8000"))
	e.Logger.Fatal(e.Start(":1234"))
}


func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("DB nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );`
    _, err := db.Exec(sql)
    if err != nil {
    	panic(err)
    }
}