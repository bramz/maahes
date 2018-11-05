package drivers

import "database/sql"

type Driver interface {
	Insert(string) string
	Select(string) string
	//	Query(q string, args ...interface{}) (*sql.Rows, error)
}

type DataBase struct {
	*sql.DB
}
