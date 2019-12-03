package last_message

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log_writer"
)

func GetLastMessage(my_db *sql.DB, user_id int) int {
	stmtOut, err := my_db.Prepare("SELECT message_id FROM last_message WHERE user_id = ?")
	if err != nil {
		go log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var mes_id int
	err = stmtOut.QueryRow(user_id).Scan(&mes_id)
	if err != nil {
		err = stmtOut.Close()
		if err != nil {
			go log_writer.ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return 0
	}

	err = stmtOut.Close()
	if err != nil {
		go log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return mes_id
}
func SetNewMessage(my_db *sql.DB, user_id int, mes_id int) {
	stmtIns, err := my_db.Prepare("UPDATE last_message SET message_id = ? WHERE user_id = ?")
	if err != nil {
		go log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(mes_id, user_id)
	if err != nil {
		go log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtIns.Close()
	if err != nil {
		go log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func AddNewUser(my_db *sql.DB, user_id int) {
	stmtIns, err := my_db.Prepare("INSERT INTO last_message VALUES(?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_id, 0)
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