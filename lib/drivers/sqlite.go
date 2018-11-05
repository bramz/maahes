package drivers

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Init(dbname string) *DataBase {
	connect, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}
	return &DataBase{connect}
}

func (db *DataBase) Execute(q string) {
	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) Insert(q string, elements []string) {
	statement, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(elements)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) Select(q string) string {
	results, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}

	var data string
	for results.Next() {
		err := results.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
	}
	return data
}

/*
func (db *DataBase) Query(q []string) string {
}
*/
