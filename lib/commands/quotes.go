package commands

import (
	"database/sql"
	"log"
	"strings"
)

type QdbCommand struct {
	*sql.DB
}

func (q QdbCommand) Handle(content []string) string {
	var arg, quote, output string

	if len(content) == 1 {
		output = q.RandQuote()
		return output
	}

	arg = content[1]
	quote = strings.Join(content[2:], " ")

	if arg == "add" {
		q.AddQuote(quote)
		output = "quote added to database"
	}

	return output
}

func (q QdbCommand) AddQuote(quote string) {
	stmt, err := q.DB.Prepare("INSERT INTO quotes (quote) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(quote)
	if err != nil {
		log.Fatal(err)
	}
}

func (q QdbCommand) RandQuote() string {
	var oquote string
	rows, err := q.DB.Query("SELECT quote FROM quotes ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&oquote)
		if err != nil {
			log.Fatal(err)
		}
	}
	return oquote
}
