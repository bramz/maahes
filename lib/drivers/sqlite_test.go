package drivers

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var q string
var elements []string
var this Driver
var db *DataBase

func TestDataBase_Insert(t *testing.T) {
	db = this.Init("data/maahes.db")
	name := "Insert"
	q = "insert into quotes(id, quote) values(?, ?)"
	elements = []string{"1", "testing"}
	t.Run(name, func(t *testing.T) {
		this.Insert(q, elements)
	})
}

func TestDataBase_Select(t *testing.T) {
	name := "Select"
	q = "select * from quotes"

	t.Run(name, func(t *testing.T) {
		this.Select(q)
	})
}
