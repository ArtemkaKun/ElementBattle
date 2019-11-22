package buffer_areas

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log_writer"
	"structs"
	"time"
)

func RegUser(my_db *sql.DB, user_info structs.UserInfo) {
	stmtIns, err := my_db.Prepare("INSERT INTO buffer_areas VALUES(?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.User_id, 0)
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
func SetArea(my_db *sql.DB, user_id int, area_id int) {
	stmtIns, err := my_db.Prepare("UPDATE buffer_areas SET area_id = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(area_id, user_id)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, change area to %v", user_id, area_id)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func GetArea(my_db *sql.DB, user_id int) int {
	stmtOut, err := my_db.Prepare("SELECT area_id FROM buffer_areas WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var area_id int
	err = stmtOut.QueryRow(user_id).Scan(&area_id)

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, change area to %v", user_id, area_id)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return area_id
}