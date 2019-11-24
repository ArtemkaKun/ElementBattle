package trains

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log_writer"
	"time"
)

func StartTrain(my_db *sql.DB, user_id int, chat_id int64, time_need time.Time) {
	stmtIns, err := my_db.Prepare("INSERT INTO users_trains VALUES (?, ?, ?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	//data_form := time_need.Format("01-02-2006 15:04:05")

	_, err = stmtIns.Exec(user_id, chat_id, time_need, nil)
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
func CheckTrain(my_db *sql.DB, time_now time.Time) int64 {
	UpdateTime(my_db, time_now)
	stmtOut, err := my_db.Prepare("SELECT chat_id FROM users_trains WHERE training_time <= now_time")
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
	stmtIns, err := my_db.Prepare("UPDATE users_trains SET now_time = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	//data_form := time_need.Format("01-02-2006 15:04:05")

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
func GetUserId(my_db *sql.DB, chat_id int64) int {
	stmtOut, err := my_db.Prepare("SELECT user_id FROM users_trains WHERE chat_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var id int
	err = stmtOut.QueryRow(chat_id).Scan(&id)
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
func DeleteTrain(my_db *sql.DB, user_id int) {
	stmtIns, err := my_db.Prepare("DELETE FROM users_trains WHERE user_id = ?")
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
func IsTraining(my_db *sql.DB, user_id int) bool {
	stmtOut, err := my_db.Prepare("SELECT user_id FROM users_trains WHERE user_id = ?")
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