package main

import (
	"database/sql"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
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

type UserStatsReg struct {
	str int
	agi int
	int int
}

type FullUserStats struct {
	LVL int
	EX_now int
	EX_next_lvl int
	HP float32
	MANA float32
	STAMINA float32
	str int
	agi int
	int int
	armor float32
	stun float32
	dodge float32
	crit float32
	fire int
	water int
	earth int
	wind int
	effect float32
	magic_arm float32
	meele_miss float32
	range_miss float32
	race string
	energy int
	skill_point int
}

type FullEnemyStats struct {
	ID int
	name string
	area_id int
	LVL int
	HP float32
	MANA float32
	STAMINA float32
	str int
	agi int
	int int
	armor float32
	stun float32
	dodge float32
	crit float32
	fire int
	water int
	earth int
	wind int
	effect float32
	magic_arm float32
	meele_miss float32
	range_miss float32
}

var log_files = LogTypes{"reg_log_bot.txt", "error_log_bot.txt", "battle_log_bot.txt", "skill_log_bot.txt", "invertory_log_bot.txt", "adventure_log_bot.txt"}

var reg_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Nod", "new_reg"),
	),
)

var first_quest_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am an exile", "first_quest_str"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am a thief", "first_quest_agi"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am a warlock", "first_quest_int"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Keep silent", "first_quest_silent"),
	),
)

var second_quest_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am a human", "second_quest_hum"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am an elf", "second_quest_elf"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am a dwarf", "second_quest_dwarf"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am an orc", "second_quest_orc"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Keep silent", "second_quest_silent"),
	),
)

var third_quest_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Give it and experimental elixir to injured soldier", "third_quest_int"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Eat it", "third_quest_str"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Trade it for good knife", "third_quest_agi"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Keep silent", "third_quest_silent"),
	),
)

var forth_quest_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Take all that I need and \n left his die", "forth_quest_agi"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Try to heal him with poisoned herbals", "forth_quest_int"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Kill and bury him", "forth_quest_str"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Keep silent", "forth_quest_silent"),
	),
)

var fifth_quest_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("New enemies", "fifth_quest_str"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Treasures", "fifth_quest_agi"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Secret knowledge", "fifth_quest_int"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Keep silent", "fifth_quest_silent"),
	),
)

var menu_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("You", "check_char_stat"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Go adventure", "go_adventure"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Invertory", "check_invertory"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Skills", "check_skilltree"),
	),
)

var area_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Forest", "1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Mountains", "2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Back to menu", "back"),
	),
)

var area_action_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Find enemy", "enemy"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Training", "train"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Meditate", "meditate"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Rest", "rest"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Back to areas menu", "back_areas"),
	),
)

var fight_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Attack", "attack"),
		tgbotapi.NewInlineKeyboardButtonData("Defence", "defence"),
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
	stmtIns, err := my_db.Prepare("INSERT INTO users_info VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.user_id, user_info.user_nickname, user_info.user_lastn, user_info.user_firstn, user_info.user_contry, 0)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	stmtIns, err = my_db.Prepare("INSERT INTO users_stats VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.user_id, 1, 0, 100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, " ", 15, 1)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	stmtIns, err = my_db.Prepare("INSERT INTO users_reg_questions VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.user_id, 0, 0, 0, 0, 0)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	stmtIns, err = my_db.Prepare("INSERT INTO buffer_areas VALUES(?, ?)")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.user_id, 0)
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

	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	if is_reg != 0 {
		return true
	} else {
		return false
	}

}
//take only str, agi, int
func TakeAttrib(my_db *sql.DB, user_id int) UserStatsReg {
	user_stats := UserStatsReg{}
	stmtOut, err := my_db.Prepare("SELECT strength, agility, intelligence FROM users_stats WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(user_id).Scan(&user_stats.str, &user_stats.agi, &user_stats.int)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return user_stats
}
//this func was used only for registration (when player answer the question)
func AddAttrib(my_db *sql.DB, user_id int, new_stats UserStatsReg, old_stats UserStatsReg) {
	stmtIns, err := my_db.Prepare("UPDATE users_stats SET strength = ?, agility = ?, intelligence = ? WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(old_stats.str + new_stats.str, old_stats.agi + new_stats.agi, old_stats.int + new_stats.int, user_id)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_writer := LogReq{time.Now(), fmt.Sprintf(" User ID is %v, add atributes. Old attributes: %v, %v, %v. New attributes: %v, %v, %v", user_id, old_stats.str, old_stats.agi, old_stats.int, new_stats.str, new_stats.agi, new_stats.int)}
	LogWrite(log_writer, log_files.reg_log)

	err = stmtIns.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
//if user try to answer one question two times, check, is question already answered
func CheckAnswers(my_db *sql.DB, user_id int, quest string) int {
	question := fmt.Sprintf("SELECT %v FROM users_reg_questions WHERE user_id = ?", quest)
	stmtOut, err := my_db.Prepare(question)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var is_answer int
	err = stmtOut.QueryRow(user_id).Scan(&is_answer)
	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return is_answer
}
//mark question as answered
func WriteAnswers(my_db *sql.DB, user_id int, quest string) {
	question := fmt.Sprintf("UPDATE users_reg_questions SET %v = 1 WHERE user_id = ?", quest)
	stmtIns, err := my_db.Prepare(question)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_id)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_writer := LogReq{time.Now(), fmt.Sprintf(" User ID is %v, answer the question %v", user_id, quest)}
	LogWrite(log_writer, log_files.reg_log)

	err = stmtIns.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func AddRace(my_db *sql.DB, user_id int, race string) {
	stmtIns, err := my_db.Prepare("UPDATE users_stats SET race = ? WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(race, user_id)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_writer := LogReq{time.Now(), fmt.Sprintf(" User ID is %v, choose race %v", user_id, race)}
	LogWrite(log_writer, log_files.reg_log)

	err = stmtIns.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func CalcStatsAfterReg(my_db *sql.DB, user_id int) {
	question := fmt.Sprintf("UPDATE users_stats SET HP = ?, MANA = ?, STAMINA = ?, armor = ?, stun_chance = ?, dodge_chance = ?, crit_chance = ?, fire_element = ?, water_element = ?, earth_element = ?, wind_element = ?, effect_chance = ?, magic_armor = ?, meele_miss_chance = ?, range_miss_chance = ? WHERE user_id = ?")
	stmtIns, err := my_db.Prepare(question)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	completed_user_stats := TakeAttrib(my_db, user_id)
	hp := 20 + float32(completed_user_stats.str) * 10 * 0.8
	stamina := 10 + 5 * float32(completed_user_stats.str) * 0.5 + 5 * float32(completed_user_stats.agi) * 0.5
	mana := 10 + 10 * float32(completed_user_stats.int) * 0.8
	armor := 5 * float32(completed_user_stats.str) * 0.1
	stun_chance := float32(completed_user_stats.str) * 0.05
	dodge_chance := float32(completed_user_stats.agi) * 0.05
	crit_chance := float32(completed_user_stats.agi) * 0.05
	effect_chance := float32(completed_user_stats.int) * 0.05
	magic_armor := float32(completed_user_stats.int) * 0.1 * 5
	meele_miss_chance := float32(30)
	range_miss_chance := float32(30)
	fire_elem := 0
	water_elem := 0
	earth_elem := 0
	wind_elem := 0
	if completed_user_stats.int >= 10 {
		fire_elem = 2
		water_elem = 2
		earth_elem = 2
		wind_elem = 2
	}

	user_stats := FullUserStats{1, 0, 100, hp, mana, stamina, completed_user_stats.str, completed_user_stats.agi, completed_user_stats.int,armor, stun_chance,dodge_chance, crit_chance, fire_elem, water_elem, earth_elem, wind_elem, effect_chance, magic_armor, meele_miss_chance, range_miss_chance, " ", 15, 1}
	_, err = stmtIns.Exec(user_stats.HP, user_stats.MANA, user_stats.STAMINA, user_stats.armor, user_stats.stun, user_stats.dodge, user_stats.crit, user_stats.fire, user_stats.water, user_stats.earth, user_stats.wind, user_stats.effect, user_stats.magic_arm, user_stats.meele_miss, user_stats.range_miss, user_id)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_writer := LogReq{time.Now(), fmt.Sprintf(" User ID is %v, complete the registration", user_id)}
	LogWrite(log_writer, log_files.reg_log)

	err = stmtIns.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func TakeFullStats(my_db *sql.DB, user_id int) FullUserStats {
	user_stats := FullUserStats{}
	stmtOut, err := my_db.Prepare("SELECT LVL, EX_now, EX_next_lvl, HP, MANA, STAMINA, strength, agility, intelligence, armor, stun_chance, dodge_chance, crit_chance, fire_element, water_element, earth_element, wind_element, effect_chance, magic_armor, meele_miss_chance, range_miss_chance, energy, skill_point FROM users_stats WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(user_id).Scan(&user_stats.LVL, &user_stats.EX_now, &user_stats.EX_next_lvl, &user_stats.HP, &user_stats.MANA, &user_stats.STAMINA, &user_stats.str, &user_stats.agi, &user_stats.int, &user_stats.armor, &user_stats.stun, &user_stats.dodge, &user_stats.crit, &user_stats.fire, &user_stats.water, &user_stats.earth, &user_stats.wind, &user_stats.effect, &user_stats.magic_arm, &user_stats.meele_miss, &user_stats.range_miss, &user_stats.energy, &user_stats.skill_point)

	stmtOut, err = my_db.Prepare("SELECT race FROM users_stats WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
	var string_stats string
	err = stmtOut.QueryRow(user_id).Scan(&string_stats)
	if (string_stats == " ") {
		user_stats.race = "Unknown"
	} else {
		user_stats.race = string_stats
	}

	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return user_stats
}
func SetArea(my_db *sql.DB, user_id int, area_id int) {
	stmtIns, err := my_db.Prepare("UPDATE buffer_areas SET area_id = ? WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(area_id, user_id)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_writer := LogReq{time.Now(), fmt.Sprintf(" User ID is %v, change area to %v", user_id, area_id)}
	LogWrite(log_writer, log_files.adventure_log)

	err = stmtIns.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func GetArea(my_db *sql.DB, user_id int) int {
	stmtOut, err := my_db.Prepare("SELECT area_id FROM buffer_areas WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var area_id int
	err = stmtOut.QueryRow(user_id).Scan(&area_id)

	log_writer := LogReq{time.Now(), fmt.Sprintf(" User ID is %v, change area to %v", user_id, area_id)}
	LogWrite(log_writer, log_files.adventure_log)

	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return area_id
}
func GetEnemy(my_db *sql.DB, area_id int, lvl int) FullEnemyStats{
	enemy_stats := FullEnemyStats{}
	stmtOut, err := my_db.Prepare("SELECT * FROM enemies WHERE enemy_area = ? AND enemy_lvl = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(area_id, lvl).Scan(&enemy_stats.ID, &enemy_stats.name, &enemy_stats.area_id, &enemy_stats.LVL, &enemy_stats.HP, &enemy_stats.MANA, &enemy_stats.STAMINA, &enemy_stats.str, &enemy_stats.agi, &enemy_stats.int, &enemy_stats.armor, &enemy_stats.stun, &enemy_stats.dodge, &enemy_stats.crit, &enemy_stats.fire, &enemy_stats.water, &enemy_stats.earth, &enemy_stats.wind, &enemy_stats.effect, &enemy_stats.magic_arm, &enemy_stats.meele_miss, &enemy_stats.range_miss)
	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return enemy_stats
}
func AddBattle(my_db *sql.DB, user_id int, enemy_stats FullEnemyStats) {
	stmtIns, err := my_db.Prepare("INSERT INTO fight_buffer VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_id, enemy_stats.ID, enemy_stats.name, enemy_stats.area_id, enemy_stats.LVL, enemy_stats.HP, enemy_stats.MANA, enemy_stats.STAMINA, enemy_stats.str, enemy_stats.agi, enemy_stats.int, enemy_stats.armor, enemy_stats.stun, enemy_stats.dodge, enemy_stats.crit, enemy_stats.fire, enemy_stats.water, enemy_stats.earth, enemy_stats.wind, enemy_stats.effect, enemy_stats.magic_arm, enemy_stats.meele_miss, enemy_stats.range_miss)
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_writer := LogReq{time.Now(), fmt.Sprintf(" User ID is %v, start battle with %v", user_id, enemy_stats.name)}
	LogWrite(log_writer, log_files.adventure_log)

	err = stmtIns.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func CheckBattle(my_db *sql.DB, user_id int) bool {
	stmtOut, err := my_db.Prepare("SELECT user_id FROM fight_buffer WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var fight_user int
	err = stmtOut.QueryRow(user_id).Scan(&fight_user)
	if err != nil {
		err = stmtOut.Close()
		if err != nil {
			ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return false
	}

	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	if fight_user != 0 {
		return true
	} else {
		return false
	}
}
func CheckBan(my_db *sql.DB, user_id int) bool {
	stmtOut, err := my_db.Prepare("SELECT is_ban FROM users_info WHERE user_id = ?")
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var ban int
	err = stmtOut.QueryRow(user_id).Scan(&ban)
	if err != nil {
		err = stmtOut.Close()
		if err != nil {
			ErrLogHandler(err.Error())
			panic(err.Error())
		}
		return false
	}

	err = stmtOut.Close()
	if err != nil {
		ErrLogHandler(err.Error())
		panic(err.Error())
	}

	if ban != 0 {
		return true
	} else {
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
			if !CheckBan(database, update.CallbackQuery.From.ID) {
				switch update.CallbackQuery.Data {
				case "new_reg":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						user_info := UserInfo{update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, update.CallbackQuery.From.LastName, update.CallbackQuery.From.FirstName, update.CallbackQuery.From.LanguageCode}
						fmt.Print(user_info.user_id)
						msg.Text = "Ok. I will ask you few questions. This information only for my report, but tell my only the truth\n" +
							"Firstly, I want to know why are you there."
						if !RegCheck(database, user_info) {
							log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, was registered", user_info.user_nickname, user_info.user_id)}
							LogWrite(log_writer, log_files.reg_log)
							RegUser(database, user_info)
							msg.ReplyMarkup = first_quest_keyboard
							my_bot.Send(msg)
						} else {
							log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, try to register, but there is user with same id already!", user_info.user_nickname, user_info.user_id)}
							LogWrite(log_writer, log_files.reg_log)
							msg.Text = "Hej, I am already have you on my list!"
							my_bot.Send(msg)
						}
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "first_quest_str":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{3, 0, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest1", second_quest_keyboard, "What is your race?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "first_quest_agi":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 3, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest1", second_quest_keyboard, "What is your race?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "first_quest_int":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 0, 3}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest1", second_quest_keyboard, "What is your race?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "first_quest_silent":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{1, 1, 1}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest1", second_quest_keyboard, "What is your race?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "second_quest_hum":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{1, 2, 0}
						AddRace(database, update.CallbackQuery.From.ID, "Human")
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest2", third_quest_keyboard, "You have a chocolate cake. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "second_quest_elf":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 2, 1}
						AddRace(database, update.CallbackQuery.From.ID, "Elf")
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest2", third_quest_keyboard, "You have a chocolate cake. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "second_quest_dwarf":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{2, 0, 1}
						AddRace(database, update.CallbackQuery.From.ID, "Dwarf")
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest2", third_quest_keyboard, "You have a chocolate cake. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "second_quest_orc":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{2, 1, 0}
						AddRace(database, update.CallbackQuery.From.ID, "Orc")
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest2", third_quest_keyboard, "You have a chocolate cake. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "second_quest_silent":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{1, 1, 1}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest2", third_quest_keyboard, "You have a chocolate cake. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "third_quest_str":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{3, 0, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest3", forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "third_quest_agi":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 3, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest3", forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "third_quest_int":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 0, 3}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest3", forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "third_quest_silent":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{1, 1, 1}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest3", forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "forth_quest_str":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{3, 0, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest4", fifth_quest_keyboard, "What are you looking here?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "forth_quest_agi":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 3, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest4", fifth_quest_keyboard, "What are you looking here?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "forth_quest_int":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 0, 3}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest4", fifth_quest_keyboard, "What are you looking here?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "forth_quest_silent":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{1, 1, 1}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest4", fifth_quest_keyboard, "What are you looking here?"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "fifth_quest_str":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{3, 0, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest5", menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "fifth_quest_agi":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 3, 0}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest5", menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "fifth_quest_int":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{0, 0, 3}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest5", menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "fifth_quest_silent":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := UserStatsReg{1, 1, 1}
						my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName, "quest5", menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "check_char_stat":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stast := TakeFullStats(database, update.CallbackQuery.From.ID)
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						message := fmt.Sprintf("Your lvl: %v \nExperiense you have: %v \nExperience need to next lvl: %v \nSkill points you have: %v \nYour energy: %v \nYour race: %v \nYour HP: %v \nYour stamina: %v \nYour MP: %v \n\nAttributes\n\nYour strength: %v \nYour agility: %v \nYour intelligence: %v \nYour armor: %v \nYour magic armor: %v \nYour stun chance: %v%% \nYour crit chance: %v%% \nYour dodge chance: %v%% \nYour magic effect chance: %v%% \nYour meele miss chance: %v%% \nYour range miss chance: %v%% \n\nMagic elements\n\nFire element: %v \nWater element: %v \nEarth element: %v \nWind element: %v \n", user_stast.LVL, user_stast.EX_now, user_stast.EX_next_lvl, user_stast.skill_point, user_stast.energy, user_stast.race, user_stast.HP, user_stast.STAMINA, user_stast.MANA, user_stast.str, user_stast.agi, user_stast.int, user_stast.armor, user_stast.magic_arm, user_stast.stun, user_stast.crit, user_stast.dodge, user_stast.effect, user_stast.meele_miss, user_stast.range_miss, user_stast.fire, user_stast.water, user_stast.earth, user_stast.wind)
						msg.Text = message
						msg.ReplyMarkup = menu_keyboard
						my_bot.Send(msg)
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "go_adventure":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						msg.Text = "Choose the location"
						msg.ReplyMarkup = area_keyboard
						my_bot.Send(msg)
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "back":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Back to menu")
						//msg.Text = "Choose the location"
						msg.ReplyMarkup = menu_keyboard
						my_bot.Send(msg)
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "1":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						area_act := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						area_act.Text = "You in forest"
						area_act.ReplyMarkup = area_action_keyboard
						SetArea(database, update.CallbackQuery.From.ID, 1)
						my_bot.Send(area_act)
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "2":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						fmt.Print(2)
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "enemy":
					if !CheckBattle(database, update.CallbackQuery.From.ID) {
						my_bot.Send(EnemySearcher(update.CallbackQuery.Message.Chat.ID, database, update.CallbackQuery.From.ID))
					} else {
						my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
					}
				case "attack":
					my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Wolf was attacked by You. Wold is dead"))
					area_act := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
					area_act.Text = "You in forest"
					area_act.ReplyMarkup = area_action_keyboard
					my_bot.Send(area_act)
				case "defence":
					my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You was attacked by wolf. You is dead"))
				}
			} else {
				my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are in prison!"))
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
				msg.Text = "Hello, newcomer. This is a dangerous world, full of monsters, bandits, demons and other evil, that want to kill you. If you want to survive - follow my instructions!\n"
				msg.ReplyMarkup = reg_keyboard
			} else {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, start bot, but this user already registered!", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer, log_files.reg_log)
				mes_for_registered := fmt.Sprintf("Hello, %v. What are you doing here? Or you just lost, little chicken?\n", user_info.user_nickname)
				msg.Text = mes_for_registered
			}
		case "menu":
			msg.ReplyMarkup = menu_keyboard
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
				LogWrite(log_writer, log_files.adventure_log)
				msg.Text = "Please, register first"
			} else {
				log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, started adventure", user_info.user_nickname, user_info.user_id)}
				LogWrite(log_writer, log_files.adventure_log)
				msg.ReplyMarkup = area_keyboard
			}
		default:
			msg.Text = "Bad command"
		}

		my_bot.Send(msg)
	}
}
func RegQuestion(chat_string int64, stats UserStatsReg, my_db *sql.DB, user_id int, user_name string, question string, keys tgbotapi.InlineKeyboardMarkup, next_quest string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chat_string, "")
	if CheckAnswers(my_db, user_id, question) == 0 {
		user_stats := TakeAttrib(my_db, user_id)
		AddAttrib(my_db, user_id, stats, user_stats)
		log_writer := LogReq{time.Now(), fmt.Sprintf(" User %v, ID is %v, answer on %v [%v, %v, %v]", user_name, user_id, question, stats.str, stats.agi, stats.int)}
		LogWrite(log_writer, log_files.reg_log)
		msg.Text = next_quest
		msg.ReplyMarkup = keys
		WriteAnswers(my_db, user_id, question)
		if question == "quest5" {
			CalcStatsAfterReg(my_db, user_id)
		}
	} else {
		msg.Text = "You are already answer!"
	}
	return msg
}
func EnemySearcher(chat_string int64, my_db *sql.DB, user_id int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chat_string, "")
	user_stats := TakeFullStats(my_db, user_id)
	if user_stats.energy > 0 {
		enemy_higher_edge := user_stats.LVL + 2
		rand.Seed(time.Now().UnixNano())
		enemy_lvl := rand.Intn(enemy_higher_edge - user_stats.LVL + 1) + user_stats.LVL
		area_now := GetArea(my_db, user_id)
		enemy_stats := GetEnemy(my_db, area_now, enemy_lvl)
		AddBattle(my_db, user_id,enemy_stats)
		msg.Text = fmt.Sprintf("You was attacked by %v \n", enemy_stats.name) + fmt.Sprintf("Enemy lvl: %v \nEnemy HP: %v \nEnemy stamina: %v \nEnemy MP: %v \n", enemy_stats.LVL, enemy_stats.HP, enemy_stats.STAMINA, enemy_stats.MANA)
		msg.ReplyMarkup = fight_keyboard
	} else {
		msg.Text = "You are tired and no have energy to fight."
	}

	return msg
}

//------------------------------------------------------ BOT SECTION END

func main() {

	bot := BotStart()
	db := DBStart()
	defer db.Close()

	BotUpdateLoop(bot, db)

}
