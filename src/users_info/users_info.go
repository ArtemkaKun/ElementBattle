package users_info

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log_writer"
	"structs"
)

func RegUser(my_db *sql.DB, user_info structs.UserInfo) {
	stmtIns, err := my_db.Prepare("INSERT INTO users_info VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.User_id, user_info.User_nickname, user_info.User_lastn, user_info.User_firstn, user_info.User_contry, 0)
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
func RegCheck(my_db *sql.DB, user_info structs.UserInfo) bool {
	stmtOut, err := my_db.Prepare("SELECT user_id FROM users_info WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var is_reg int
	err = stmtOut.QueryRow(user_info.User_id).Scan(&is_reg)
	if err != nil {
		err = stmtOut.Close()
		if err != nil {
			log_writer.ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return false
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	if is_reg != 0 {
		return true
	} else {
		return false
	}

}
func CheckBan(my_db *sql.DB, user_id int) bool {
	stmtOut, err := my_db.Prepare("SELECT is_ban FROM users_info WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var ban int
	err = stmtOut.QueryRow(user_id).Scan(&ban)
	if err != nil {
		err = stmtOut.Close()
		if err != nil {
			log_writer.ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return false
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	if ban != 0 {
		return true
	} else {
		return false
	}
}
