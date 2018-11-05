package drivers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Init(dbname string) (*DataBase, error) {
	connect, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	return &DataBase{connect}, nil
}

func (db *DataBase) Insert(q []string) string {
	statement, err := db.Prepare(q)
	if err != nil {
		return fmt.Print(err)
	}
	_, err := statement.Exec(statement)
	if err != nil {
		return fmt.Print(err)
	}
	return fmt.Print("Statement executed")
}

func (db *DataBase) Select(q []string) string {
	statement, err := db.Prepare(q)
	if err != nil {
		return fmt.Print("Query failed")
	}

	results := db.Query(statement)
	return results
}

/*
func (db *DataBase) Query(q []string) string {
}
*/
