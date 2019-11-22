package structs

import "time"

type LogRequest struct {
	Log_time    time.Time
	Log_message string
}
type LogTypes struct {
	Reg_log       string
	Err_log       string
	Battle_log    string
	Skill_log     string
	Invertory_log string
	Adventure_log string
}

type UserInfo struct {
	User_id       int
	User_nickname string
	User_lastn    string
	User_firstn   string
	User_contry   string
	Is_baned      int
}
type UserCoreStats struct {
	Str int
	Agi int
	Int int
}
type FullUserStats struct {
	Id 				  int
	Title 			  string
	Race 			  string
	Energy 			  int
	Lvl 			  int
	Ex_now 			  int
	Ex_next_lvl 	  int
	Skill_point 	  int
	Hp 				  float32
	Mp                float32
	Stamina           float32
	Str               int
	Agi               int
	Int               int
	Armor             float32
	Magic_armor       float32
	Stun_chance       float32
	Dodge_chance      float32
	Crit_chance       float32
	Effect_chance     float32
	Meele_miss_chance float32
	Range_miss_chance float32
	Fire 			  int
	Water 			  int
	Earth 			  int
	Wind 			  int
}
type FullBotEnemyStats struct {
	Id 				  int
	Name 			  string
	Area_id 		  int
	Lvl 			  int
	Hp 				  float32
	Mp                float32
	Stamina           float32
	Str               int
	Agi               int
	Int               int
	Armor             float32
	Magic_armor       float32
	Stun_chance       float32
	Dodge_chance      float32
	Crit_chance       float32
	Effect_chance     float32
	Meele_miss_chance float32
	Range_miss_chance float32
	Fire 			  int
	Water 			  int
	Earth 			  int
	Wind 			  int
}

type FullFightBotEnemyStats struct {
	Id 				  int
	Name 			  string
	Area_id 		  int
	Lvl 			  int
	Hp 				  float32
	Mp                float32
	Stamina           float32
	Str               int
	Agi               int
	Int               int
	Armor             float32
	Magic_armor       float32
	Stun_chance       float32
	Dodge_chance      float32
	Crit_chance       float32
	Effect_chance     float32
	Meele_miss_chance float32
	Range_miss_chance float32
	Fire 			  int
	Water 			  int
	Earth 			  int
	Wind 			  int
	Is_stuned 		  int
	Is_under_effects  string
}
type FightUserStats struct {
	Is_under_effect string
	Is_stunned   int
	Is_block  int
}
