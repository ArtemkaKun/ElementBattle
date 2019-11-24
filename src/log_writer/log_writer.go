package log_writer

import (
	"log"
	"os"
	"structs"
	"time"
)

var Log_files = structs.LogTypes {
	"reg_log_bot.txt",
	"error_log_bot.txt",
	"battle_log_bot.txt",
	"skill_log_bot.txt",
	"invertory_log_bot.txt",
	"adventure_log_bot.txt",
	"train_log_bot.txt"}

func LogWrite(log_request structs.LogRequest, log_type string) {
	f, err := os.OpenFile(log_type, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		ErrLogHandler(err.Error())
		return
	}

	data_form := log_request.Log_time.Format("01-02-2006 15:04:05")

	_, err = f.WriteString(data_form + log_request.Log_message +"\n")
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		return
	}
}
func ErrLogHandler(error string) {
	log_insert := structs.LogRequest{time.Now(), error}
	LogWrite(log_insert, Log_files.Err_log)
}
