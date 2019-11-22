package pve_fight_buffer

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log_writer"
	"structs"
	"time"
)

func AddBattle(my_db *sql.DB, user_id int, enemy_stats structs.FullBotEnemyStats) {
	stmtIns, err := my_db.Prepare("INSERT INTO pve_fight_buffer VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(
		user_id, 0, "", 0, enemy_stats.Id, enemy_stats.Name, enemy_stats.Area_id, enemy_stats.Lvl, enemy_stats.Hp, enemy_stats.Mp, enemy_stats.Stamina,
		enemy_stats.Str, enemy_stats.Agi, enemy_stats.Int, enemy_stats.Armor, enemy_stats.Magic_armor, enemy_stats.Stun_chance, enemy_stats.Dodge_chance,
		enemy_stats.Crit_chance, enemy_stats.Effect_chance, enemy_stats.Meele_miss_chance, enemy_stats.Range_miss_chance, enemy_stats.Fire,
		enemy_stats.Water, enemy_stats.Earth, enemy_stats.Wind, 0, "")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User ID is %v, start battle with %v", user_id, enemy_stats.Name)}
	log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)

	err = stmtIns.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}
}
func CheckBattle(my_db *sql.DB, user_id int) bool {
	stmtOut, err := my_db.Prepare("SELECT user_id FROM pve_fight_buffer WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	var fight_user int
	err = stmtOut.QueryRow(user_id).Scan(&fight_user)
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

	if fight_user != 0 {
		return true
	} else {
		return false
	}
}
func GetFightEnemyStats(my_db *sql.DB, user_id int) structs.FullFightBotEnemyStats {
	enemy_stats := structs.FullFightBotEnemyStats{}
	stmtOut, err := my_db.Prepare("SELECT enemy_id, enemy_name, enemy_area, enemy_lvl, enemy_hp, enemy_mana, enemy_stamina, enemy_str, enemy_agi, enemy_int, enemy_armor, enemy_stun_chance, enemy_dodge_chance, enemy_crit_chance, enemy_fire, enemy_water, enemy_earth, enemy_wind, enemy_effect_chance, enemy_magic_armor, enemy_meele_miss_chance, enemy_range_miss_chance, enemy_under_effect, enemy_under_stun FROM pve_fight_buffer WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(user_id).Scan(
		&enemy_stats.Id, &enemy_stats.Name, &enemy_stats.Area_id, &enemy_stats.Lvl, &enemy_stats.Hp,
		&enemy_stats.Mp, &enemy_stats.Stamina, &enemy_stats.Str, &enemy_stats.Agi, &enemy_stats.Int,
		&enemy_stats.Armor, &enemy_stats.Stun_chance, &enemy_stats.Dodge_chance,
		&enemy_stats.Crit_chance, &enemy_stats.Fire, &enemy_stats.Water, &enemy_stats.Earth, &enemy_stats.Wind, &enemy_stats.Effect_chance,
		&enemy_stats.Magic_armor, &enemy_stats.Meele_miss_chance, &enemy_stats.Range_miss_chance, &enemy_stats.Is_under_effects,
		&enemy_stats.Is_stuned)

	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return enemy_stats
}
func GetBattleUserStats(my_db *sql.DB, user_id int) structs.FightUserStats {
	user_stats :=  structs.FightUserStats{}
	stmtOut, err := my_db.Prepare("SELECT user_under_stun, user_under_effect, user_block FROM pve_fight_buffer WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(user_id).Scan(&user_stats.Is_stunned, &user_stats.Is_under_effect, &user_stats.Is_block)
	err = stmtOut.Close()
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return user_stats
}
func SetNewBotEnemyFightStats(my_db *sql.DB, user_id int, enemy_stats structs.FullFightBotEnemyStats) {
	stmtIns, err := my_db.Prepare("UPDATE pve_fight_buffer SET enemy_hp = ?, enemy_mana = ?, enemy_stamina = ?, enemy_str = ?, enemy_agi = ?, enemy_int = ?, enemy_armor = ?, enemy_stun_chance = ?, enemy_dodge_chance = ?, enemy_crit_chance = ?, enemy_fire = ?, enemy_water = ?, enemy_earth = ?, enemy_wind = ?, enemy_effect_chance = ?, enemy_magic_armor = ?, enemy_meele_miss_chance = ?, enemy_range_miss_chance = ?, enemy_under_effect = ?, enemy_under_stun = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(
		enemy_stats.Hp, enemy_stats.Mp, enemy_stats.Stamina, enemy_stats.Str, enemy_stats.Agi, enemy_stats.Int,
		enemy_stats.Armor, enemy_stats.Stun_chance, enemy_stats.Dodge_chance, enemy_stats.Crit_chance,
		enemy_stats.Fire, enemy_stats.Water, enemy_stats.Earth, enemy_stats.Wind, enemy_stats.Effect_chance,
		enemy_stats.Magic_armor, enemy_stats.Meele_miss_chance, enemy_stats.Range_miss_chance, enemy_stats.Is_under_effects,
		enemy_stats.Is_stuned, user_id)

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
func SetNewUserFightStats(my_db *sql.DB, user_id int, user_fight_stats structs.FightUserStats) {
	stmtIns, err := my_db.Prepare("UPDATE pve_fight_buffer SET user_under_stun = ? , user_under_effect = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(user_fight_stats.Is_stunned, user_fight_stats.Is_under_effect, user_id)
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
func SetNewUserStats(my_db *sql.DB, user_id int, full_user_stats structs.FullUserStats) {
	stmtIns, err := my_db.Prepare("UPDATE users_stats SET LVL = ?, EX_now = ?, EX_next_lvl = ?,HP = ?, MANA = ?, STAMINA = ?, strength = ?, agility = ?, intelligence = ?, armor = ?, stun_chance = ?, dodge_chance = ?, crit_chance = ?, fire_element = ?, water_element = ?, earth_element = ?, wind_element = ?, effect_chance = ?, magic_armor = ?, meele_miss_chance = ?, range_miss_chance = ?, energy = ?, skill_point = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(
		full_user_stats.Lvl, full_user_stats.Ex_now, full_user_stats.Ex_next_lvl, full_user_stats.Hp, full_user_stats.Mp,
		full_user_stats.Stamina, full_user_stats.Str, full_user_stats.Agi, full_user_stats.Int, full_user_stats.Armor,
		full_user_stats.Stun_chance, full_user_stats.Dodge_chance, full_user_stats.Crit_chance, full_user_stats.Fire,
		full_user_stats.Water, full_user_stats.Earth, full_user_stats.Wind, full_user_stats.Effect_chance,
		full_user_stats.Magic_armor, full_user_stats.Meele_miss_chance, full_user_stats.Range_miss_chance,
		full_user_stats.Energy, full_user_stats.Skill_point, user_id)
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
func DeleteFight(my_db *sql.DB, user_id int) {
	stmtIns, err := my_db.Prepare("DELETE FROM pve_fight_buffer WHERE user_id = ?")
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
func SetBlock(my_db *sql.DB, user_id int, block int) {
	stmtIns, err := my_db.Prepare("UPDATE pve_fight_buffer SET user_block = ? WHERE user_id = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	_, err = stmtIns.Exec(block, user_id)
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