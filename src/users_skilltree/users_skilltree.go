package users_skilltree

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log_writer"
	"structs"
)

func RegUser(my_db *sql.DB, user_info structs.UserInfo) {
	stmtIns, err := my_db.Prepare("INSERT INTO users_skilltree VALUES(?, ?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.User_id, 0, 0)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
