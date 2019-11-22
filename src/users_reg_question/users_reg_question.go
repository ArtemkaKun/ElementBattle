package users_reg_question

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log_writer"
	"structs"
	"time"
)

func RegUser(my_db *sql.DB, user_info structs.UserInfo) {
	stmtIns, err := my_db.Prepare("INSERT INTO users_reg_questions VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.User_id, 0, 0, 0, 0, 0)
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

//if user try to answer one question two times, check, is question already answered
func CheckAnswers(my_db *sql.DB, user_id int, quest string) int {
	question := fmt.Sprintf("SELECT %v FROM users_reg_questions WHERE user_id = ?", quest)
	stmtOut, err := my_db.Prepare(question)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var is_answer int
	err = stmtOut.QueryRow(user_id).Scan(&is_answer)

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return is_answer
}
func CheckAllAnswers(my_db *sql.DB, user_id int) bool {
	stmtOut, err := my_db.Prepare("SELECT * FROM users_reg_questions WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var user_ID int
	var quest int
	var quest2 int
	var quest3 int
	var quest4 int
	var quest5 int

	err = stmtOut.QueryRow(user_id).Scan(&user_ID, &quest, &quest2, &quest3, &quest4, &quest5)

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	if quest != 0 && quest2 != 0 && quest3 != 0 && quest4 != 0 && quest5 != 0 {
		return true
	} else {
		return false
	}
}
//mark question as answered
func WriteAnswers(my_db *sql.DB, user_id int, quest string) {
	question := fmt.Sprintf("UPDATE users_reg_questions SET %v = 1 WHERE user_id = ?", quest)
	stmtIns, err := my_db.Prepare(question)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_id)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, answer the question %v", user_id, quest)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}