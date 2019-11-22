package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log_writer"
)

func DBStart() *sql.DB {
	my_db, err := sql.Open("mysql", "root:1337@/elementbattles")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
	} else {
		err = my_db.Ping()
		if err != nil {
			log_writer.ErrLogHandler(err.Error())
		}
	}
	return my_db
}
