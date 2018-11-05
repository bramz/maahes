package drivers

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var q []string
var this Driver
var db *DataBase

func TestDataBase_Insert(t *testing.T) {
	name := "Insert"
	want := "Statement executed"
	q = "insert into quotes (quote) values("test quote")"

	t.Run(name, func(t *testing.T)) {
		statement := db.Prepare(q)
		if got := this.Insert(statement); got != want {
			t.Errorf("Driver.Insert() = %v, want %v", got, want)
		}
	}
}

func TestDataBase_Select(t *testing.T) {
	name := "Select"
	q = "select * from quotes"

	t.Run(name, func(t *testing.T)) {
		statement := db.Prepare(q)
		this.Select(statement)
	}
}