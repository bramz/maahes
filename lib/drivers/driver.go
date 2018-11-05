package drivers

import "database/sql"

type Driver interface {
	Init(string) (*DataBase, error)
	Execute(string)
	Insert(string, []string)
	Select(string) string
	//	Query(q string, args ...interface{}) (*sql.Rows, error)
}

type DataBase struct {
	*sql.DB
}
