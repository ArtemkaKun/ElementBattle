package users_stats

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log_writer"
	"structs"
	"time"
)

func RegUser(my_db *sql.DB, user_info structs.UserInfo) {
	stmtIns, err := my_db.Prepare("INSERT INTO users_stats VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_info.User_id,"","", 15, 1, 0, 100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
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

func GetStrAttrib(my_db *sql.DB, user_id int) int {
	stmtOut, err := my_db.Prepare("SELECT strength FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var str int
	err = stmtOut.QueryRow(user_id).Scan(&str)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return str
}
func GetAgiAttrib(my_db *sql.DB, user_id int) int {
	stmtOut, err := my_db.Prepare("SELECT agility FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var agi int
	err = stmtOut.QueryRow(user_id).Scan(&agi)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return agi
}
func GetIntAttrib(my_db *sql.DB, user_id int) int {
	stmtOut, err := my_db.Prepare("SELECT intelligence FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var int int
	err = stmtOut.QueryRow(user_id).Scan(&int)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return int
}
func GetCoreAttribs(my_db *sql.DB, user_id int) structs.UserCoreStats {
	user_core_stats := structs.UserCoreStats{}

	stmtOut, err := my_db.Prepare("SELECT strength, agility, intelligence FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(user_id).Scan(&user_core_stats.Str, &user_core_stats.Agi, &user_core_stats.Int)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return user_core_stats
}

func AddAttrib(my_db *sql.DB, user_id int, new_stats structs.UserCoreStats) {
	old_stats := GetCoreAttribs(my_db, user_id)

	stmtIns, err := my_db.Prepare("UPDATE users_stats SET strength = ?, agility = ?, intelligence = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(old_stats.Str + new_stats.Str, old_stats.Agi + new_stats.Agi, old_stats.Int + new_stats.Int, user_id)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, add atributes. Old attributes: %v, %v, %v. New attributes: %v, %v, %v", user_id, old_stats.Str, old_stats.Agi, old_stats.Int, new_stats.Str, new_stats.Agi, new_stats.Int)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func AddRace(my_db *sql.DB, user_id int, race string) {
	stmtIns, err := my_db.Prepare("UPDATE users_stats SET race = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(race, user_id)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, choose race %v", user_id, race)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func GetRace(my_db *sql.DB, user_id int) string {
	stmtOut, err := my_db.Prepare("SELECT race FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var race string
	err = stmtOut.QueryRow(user_id).Scan(&race)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return race
}
func CalcStatsAfterReg(my_db *sql.DB, user_id int) {
	const CHANCE_COF = 0.05
	const HP_MIN = 20
	const MP_ST_MIN = 10
	const HP_MP_COF = 8
	const ARMOR_COF = 0.5
	const ST_COF = 2.5
	const START_ELEM = 2
	const MISS_COF = 30

	completed_user_stats := GetCoreAttribs(my_db, user_id)
	hp := HP_MIN + float32(completed_user_stats.Str) * HP_MP_COF
	stamina := MP_ST_MIN + float32(completed_user_stats.Str) * ST_COF + float32(completed_user_stats.Agi) * ST_COF
	mana := MP_ST_MIN + float32(completed_user_stats.Int) * HP_MP_COF
	armor := float32(completed_user_stats.Str) * ARMOR_COF
	stun_chance := float32(completed_user_stats.Str) * CHANCE_COF
	dodge_chance := float32(completed_user_stats.Agi) * CHANCE_COF
	crit_chance := float32(completed_user_stats.Agi) * CHANCE_COF
	effect_chance := float32(completed_user_stats.Int) * CHANCE_COF
	magic_armor := float32(completed_user_stats.Int) * ARMOR_COF
	meele_miss_chance := float32(MISS_COF)
	range_miss_chance := float32(MISS_COF)
	fire_elem := 0
	water_elem := 0
	earth_elem := 0
	wind_elem := 0
	if completed_user_stats.Int >= 10 {
		fire_elem = START_ELEM
		water_elem = START_ELEM
		earth_elem = START_ELEM
		wind_elem = START_ELEM
	}

	stmtIns, err := my_db.Prepare("UPDATE users_stats SET HP = ?, MANA = ?, STAMINA = ?, armor = ?, stun_chance = ?, dodge_chance = ?, crit_chance = ?, fire_element = ?, water_element = ?, earth_element = ?, wind_element = ?, effect_chance = ?, magic_armor = ?, meele_miss_chance = ?, range_miss_chance = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	user_stats := structs.FullUserStats{
		user_id, "", GetRace(my_db, user_id), 15, 1,
		0, 100, 1, hp,mana, stamina,completed_user_stats.Str,
		completed_user_stats.Agi, completed_user_stats.Int, armor, magic_armor, stun_chance,
		dodge_chance, crit_chance, effect_chance, meele_miss_chance,
		range_miss_chance, fire_elem, water_elem, earth_elem, wind_elem}
	_, err = stmtIns.Exec(
		user_stats.Hp, user_stats.Mp, user_stats.Stamina, user_stats.Armor, user_stats.Stun_chance,
		user_stats.Dodge_chance, user_stats.Crit_chance, user_stats.Fire, user_stats.Water, user_stats.Earth, user_stats.Wind,
		user_stats.Effect_chance, user_stats.Magic_armor, user_stats.Meele_miss_chance, user_stats.Range_miss_chance, user_id)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, complete the registration", user_id)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func TakeFullStats(my_db *sql.DB, user_id int) structs.FullUserStats {
	user_stats := structs.FullUserStats{}
	stmtOut, err := my_db.Prepare("SELECT * FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(user_id).Scan(
		&user_stats.Id, &user_stats.Title, &user_stats.Race, &user_stats.Energy, &user_stats.Lvl, &user_stats.Ex_now,
		&user_stats.Ex_next_lvl, &user_stats.Skill_point, &user_stats.Hp, &user_stats.Mp, &user_stats.Stamina,
		&user_stats.Str, &user_stats.Agi, &user_stats.Int, &user_stats.Armor, &user_stats.Magic_armor, &user_stats.Stun_chance, &user_stats.Dodge_chance,
		&user_stats.Crit_chance, &user_stats.Effect_chance, &user_stats.Meele_miss_chance, &user_stats.Range_miss_chance, &user_stats.Fire, &user_stats.Water, &user_stats.Earth, &user_stats.Wind)

	if (user_stats.Race == "") {
		user_stats.Race = "Unknown"
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return user_stats
}
func TakeMeeleMissChance(my_db *sql.DB, user_id int) float32 {
	stmtOut, err := my_db.Prepare("SELECT meele_miss_chance FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var chance float32
	err = stmtOut.QueryRow(user_id).Scan(&chance)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return chance
}
func TakeRangeMissChance(my_db *sql.DB, user_id int) float32 {
	stmtOut, err := my_db.Prepare("SELECT range_miss_chance FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var chance float32
	err = stmtOut.QueryRow(user_id).Scan(&chance)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return chance
}
func SetMeeleMiss(my_db *sql.DB, user_id int, new_chance float32) {
	stmtIns, err := my_db.Prepare("UPDATE users_stats SET  meele_miss_chance = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(new_chance, user_id)

	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func SetRangeMiss(my_db *sql.DB, user_id int, new_chance float32) {
	stmtIns, err := my_db.Prepare("UPDATE users_stats SET  range_miss_chance = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(new_chance, user_id)

	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func RecalcStats(my_db *sql.DB, user_id int) {
	const CHANCE_COF = 0.05
	const ARMOR_COF = 0.5

	completed_user_stats := GetCoreAttribs(my_db, user_id)
	armor := float32(completed_user_stats.Str) * ARMOR_COF
	stun_chance := float32(completed_user_stats.Str) * CHANCE_COF
	dodge_chance := float32(completed_user_stats.Agi) * CHANCE_COF
	crit_chance := float32(completed_user_stats.Agi) * CHANCE_COF
	effect_chance := float32(completed_user_stats.Int) * CHANCE_COF
	magic_armor := float32(completed_user_stats.Int) * ARMOR_COF


	stmtIns, err := my_db.Prepare("UPDATE users_stats SET armor = ?, stun_chance = ?, dodge_chance = ?, crit_chance = ?, effect_chance = ?, magic_armor = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	user_stats := structs.FullUserStats{
		user_id, "", GetRace(my_db, user_id), 15, 1,
		0, 100, 1, 0,0, 0,completed_user_stats.Str,
		completed_user_stats.Agi, completed_user_stats.Int, armor, magic_armor, stun_chance,
		dodge_chance, crit_chance, effect_chance, 0,
		0, 0, 0, 0, 0}
	_, err = stmtIns.Exec(
		user_stats.Armor, user_stats.Stun_chance,
		user_stats.Dodge_chance, user_stats.Crit_chance, user_stats.Effect_chance, user_stats.Magic_armor,
		user_id)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, recalc stats", user_id)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func RecalcRest(my_db *sql.DB, user_id int) {
	const HP_MIN = 20
	const MP_ST_MIN = 10
	const HP_MP_COF = 8
	const ST_COF = 2.5

	completed_user_stats := GetCoreAttribs(my_db, user_id)
	hp := HP_MIN + float32(completed_user_stats.Str) * HP_MP_COF
	stamina := MP_ST_MIN + float32(completed_user_stats.Str) * ST_COF + float32(completed_user_stats.Agi) * ST_COF
	mana := MP_ST_MIN + float32(completed_user_stats.Int) * HP_MP_COF


	stmtIns, err := my_db.Prepare("UPDATE users_stats SET HP = ?, MANA = ?, STAMINA = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	user_stats := structs.FullUserStats{
		user_id, "", GetRace(my_db, user_id), 15, 1,
		0, 100, 1, hp,mana, stamina,completed_user_stats.Str,
		completed_user_stats.Agi, completed_user_stats.Int, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0, 0}
	_, err = stmtIns.Exec(
		user_stats.Hp, user_stats.Mp, user_stats.Stamina, user_id)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, recalc rest stats", user_id)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func GetElementsStats(my_db *sql.DB, user_id int) structs.UserElementsStats {
	stmtOut, err := my_db.Prepare("SELECT fire_element, water_element, earth_element, wind_element FROM users_stats WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	user_elements := structs.UserElementsStats{}
	err = stmtOut.QueryRow(user_id).Scan(&user_elements.Fire, &user_elements.Water, &user_elements.Earth, &user_elements.Wind)
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return user_elements
}
func SetElements(my_db *sql.DB, user_id int, new_stats structs.UserElementsStats) {
	stmtIns, err := my_db.Prepare("UPDATE users_stats SET fire_element = ?, water_element = ?, earth_element = ?, wind_element = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(new_stats.Fire, new_stats.Water, new_stats.Earth, new_stats.Wind, user_id)

	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}