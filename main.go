package main

import (
	"database/sql"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

type UserInfo struct {
	user_id       int
	user_nickname string
	user_lastn    string
	user_firstn   string
	user_contry   string
}
type LogReq struct {
	log_time    time.Time
	log_message string
}
type LogTypes struct {
	reg_log       string
	err_log       string
	battle_log    string
	skill_log     string
	invertory_log string
	adventure_log string
}

var log_files = LogTypes{"reg_log_bot.txt", "error_log_bot.txt", "battle_log_bot.txt", "skill_log_bot.txt", "invertory_log_bot.txt", "adventure_log_bot.txt"}

var reg_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Nod", "new_reg"),
	),
)

var area_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Forest", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Mountains", "2"),
	),
)

var area_action_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Find enemy", "enemy"),
	),
)

var fight_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ujebat`", "attack"),
		tgbotapi.NewInlineKeyboardButtonData("Zdat`", "defence"),
	),
)

//------------------------------------------------------ LOG SECTION

func LogWrite(log_request LogReq, log_type string) {
	f, err := os.OpenFile(log_type, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		ErrLogHandler(err.Error())
		return
	}

	data_form := log_request.log_time.Format("01-02-2006 15:04:05")

	_, err = f.WriteString(data_form + log_request.log_message +"\n")
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
	log_writer := LogReq{time.Now(), error}
	LogWrite(log_writer, log_files.err_log)
}

//------------------------------------------------------ LOG SECTION END

//------------------------------------------------------ DB SECTION

func DBStart() *sql.DB {
	db, err := sql.Open("mysql", "root:1337@/elementbattles")
	if err != nil {
		ErrLogHandler(err.Error())
	} else {
		err = db.Ping()
		if err != nil {
			ErrLogHandler(err.Error())
		}
	}
	return db
}
func RegUser(my_db *sql.DB, user_info UserInfo) {
	stmtIns, err := my_db.Prepare("INSERT INTO users_info VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.user_id, user_info.user_nickname, user_info.user_lastn, user_info.user_firstn, user_info.user_contry)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	stmtIns, err = my_db.Prepare("INSERT INTO users_stats VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.user_id, 1, 1, 1, 100, 100, 100)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtIns.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func RegCheck(my_db *sql.DB, user_info UserInfo) bool {
	stmtOut, err := my_db.Prepare("SELECT user_id FROM users_info WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var is_reg int
	err = stmtOut.QueryRow(user_info.user_id).Scan(&is_reg)
	if err != nil {
		err = stmtOut.Close()
		if err != nil {
			ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return false
	}

	if is_reg != 0 {
		err = stmtOut.Close()
		if err != nil {
			ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return true
	} else {
		err = stmtOut.Close()
		if err != nil {
			ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return false
	}

}

//------------------------------------------------------ DB SECTION END

//------------------------------------------------------ BOT SECTION

func BotStart() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("970898716:AAG4n8sEnLIxdeffziIRs0oy80uj6osHtSE")
	if err != nil {
		ErrLogHandler(err.Error())
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Autorised on account %s", bot.Self.UserName)

	return bot
}
func BotUpdateLoop(my_bot *tgbotapi.BotAPI, database *sql.DB) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := my_bot.GetUpdatesChan(u)
	if err != nil{
		ErrLogHandler(err.Error())
		log.Panic(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
				case "new_reg":
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
					user_info := UserInfo{update.CallbackQuery.Message.From.ID, update.CallbackQuery.Message.From.UserName, update.CallbackQuery.Message.From.LastName, update.CallbackQuery.Message.From.FirstName, update.CallbackQuery.Message.From.LanguageCode}
					msg.Text = "Ok. I will ask you few questions. This information only for my raport, but tell my only the truth"
					if !RegCheck(database, user_info) {
						log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, was registered", user_info.user_nickname, user_info.user_id)}
						LogWrite(log_writer, log_files.reg_log)
						RegUser(database, user_info)
					} else {
						log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, try to register, but there is user with same id already!", user_info.user_nickname, user_info.user_id)}
						LogWrite(log_writer, log_files.reg_log)
						msg.Text = "Already registered"
					}
				case "1":
					area_act := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
					area_act.Text = "You in forest"
					area_act.ReplyMarkup = area_action_keyboard
					my_bot.Send(area_act)
				case "2":
					fmt.Print(2)
				case "enemy":
					fight_act := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
					fight_act.Text = "You was attacked by wolf"
					fight_act.ReplyMarkup = fight_keyboard
					my_bot.Send(fight_act)
				case "attack":
					my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Wolf was attacked by You. Wold is dead"))
					area_act := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
					area_act.Text = "You in forest"
					area_act.ReplyMarkup = area_action_keyboard
					my_bot.Send(area_act)
				case "defence":
					my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You was attacked by wolf. You is dead"))
			}

			//my_bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID,update.CallbackQuery.Data))
		}

		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		chat_id := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chat_id, "")
		user_info := UserInfo{update.Message.From.ID, update.Message.From.UserName, update.Message.From.LastName, update.Message.From.FirstName, update.Message.From.LanguageCode}

		switch update.Message.Command() {
		case "start":
			if !RegCheck(database, user_info) {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, start bot", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer, log_files.reg_log)
				msg.Text = "Hello, newcomer. This is a dangerous world, full of monsters, bandits, demons and other evil, that want to kill you. If you want to survive - follow my instructions!\n" +
					"Firstly, I want to know why are you there."
				msg.ReplyMarkup = reg_keyboard
			} else {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, start bot, but this user already registered!", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer, log_files.reg_log)
				mes_for_registered := fmt.Sprintf("Hello, %v. What are you doing here? Or you just lost, little chicken?\n", user_info.user_nickname)
				msg.Text = mes_for_registered
			}
		case "registration":
			msg.Text = "Registration"
			if !RegCheck(database, user_info) {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, was registered", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer, log_files.reg_log)
				RegUser(database, user_info)
			} else {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, try to register, but there is user with same id already!", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer, log_files.reg_log)
				msg.Text = "Already registered"
			}
		case "start_adventure":
			msg.Text = "Adventure"
			user_info := UserInfo{update.Message.From.ID, update.Message.From.UserName, update.Message.From.LastName, update.Message.From.FirstName, update.Message.From.LanguageCode}
			if !RegCheck(database, user_info) {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, try to start adventure, but he is not registered jet!", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer)
				msg.Text = "Please, register first"
			} else {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, started adventure", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer)
				msg.ReplyMarkup = area_keyboard
			}
		default:
			msg.Text = "Bad command"
		}

		my_bot.Send(msg)
	}
}

//------------------------------------------------------ BOT SECTION END

func main() {

	bot := BotStart()
	db := DBStart()
	defer db.Close()

	BotUpdateLoop(bot, db)

}
