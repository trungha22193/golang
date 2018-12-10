package models
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	} 
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		errRow := rows.Scan(&task.ID, &task.Name)			// --> must use "&", Pointer, to fetch data like  ##rows.Scan(&task.ID, &task.Name)##
		if errRow != nil {
			panic(errRow)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, errExec := stmt.Exec(name)
	if errExec != nil {
		panic(errExec)
	}
	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, errExec := stmt.Exec(id)
	if errExec != nil {
		panic(errExec)
	}
	return result.RowsAffected()
}