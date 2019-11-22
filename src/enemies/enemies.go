package enemies

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log_writer"
	"structs"
)

func GetEnemy(my_db *sql.DB, area_id int, lvl int) structs.FullBotEnemyStats{
	enemy_stats := structs.FullBotEnemyStats{}
	stmtOut, err := my_db.Prepare("SELECT * FROM enemies WHERE enemy_area = ? AND enemy_lvl = ?")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	err = stmtOut.QueryRow(area_id, lvl).Scan(
		&enemy_stats.Id, &enemy_stats.Name, &enemy_stats.Area_id, &enemy_stats.Lvl, &enemy_stats.Hp, &enemy_stats.Mp,
		&enemy_stats.Stamina, &enemy_stats.Str, &enemy_stats.Agi, &enemy_stats.Int, &enemy_stats.Armor,
		&enemy_stats.Magic_armor, &enemy_stats.Stun_chance, &enemy_stats.Dodge_chance, &enemy_stats.Crit_chance,
		&enemy_stats.Effect_chance, &enemy_stats.Meele_miss_chance, &enemy_stats.Range_miss_chance, &enemy_stats.Fire, &enemy_stats.Water, &enemy_stats.Earth, &enemy_stats.Wind)

	err = stmtOut.Close()

	if err != nil {
		log_writer.ErrLogHandler(err.Error())
		panic(err.Error())
	}

	return enemy_stats
}
