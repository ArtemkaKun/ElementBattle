package rests

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log_writer"
	"time"
)

func StartRest(my_db *sql.DB, user_id int, time_need time.Time) {
	stmtIns, err := my_db.Prepare("INSERT INTO users_rest VALUES (?, ?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_id, time_need, nil)
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
func CheckRest(my_db *sql.DB, time_now time.Time) int64 {
	UpdateTime(my_db, time_now)
	stmtOut, err := my_db.Prepare("SELECT user_id FROM users_rest WHERE resting_time <= now_time")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var id int64
	err = stmtOut.QueryRow().Scan(&id)
	if err != nil {
		err = stmtOut.Close()
		if err != nil {
			log_writer.ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return 0
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return id
}
func UpdateTime(my_db *sql.DB, time_now time.Time) {
	stmtIns, err := my_db.Prepare("UPDATE users_rest SET now_time = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}


	_, err = stmtIns.Exec(time_now)
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
func DeleteRest(my_db *sql.DB, user_id int) {
	stmtIns, err := my_db.Prepare("DELETE FROM users_rest WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_id)
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
func IsRest(my_db *sql.DB, user_id int) bool {
	stmtOut, err := my_db.Prepare("SELECT user_id FROM users_rest WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var is_reg int
	err = stmtOut.QueryRow(user_id).Scan(&is_reg)
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

