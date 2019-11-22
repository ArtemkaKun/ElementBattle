package main

import (
	"buffer_areas"
	"database/sql"
	"db"
	"enemies"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	_ "github.com/go-sql-driver/mysql"
	"keyboards"
	"log"
	"log_writer"
	"math/rand"
	"pve_fight_buffer"
	"structs"
	"time"
	"users_info"
	"users_reg_question"
	"users_stats"
)

func BotStart() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("970898716:AAG4n8sEnLIxdeffziIRs0oy80uj6osHtSE")
	if err != nil {
		log_writer.ErrLogHandler(err.Error())
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
		log_writer.ErrLogHandler(err.Error())
		log.Panic(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			if !users_info.CheckBan(database, update.CallbackQuery.From.ID) {
				switch update.CallbackQuery.Data {
				case "new_reg":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						user_info := structs.UserInfo {
							update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName,
							update.CallbackQuery.From.LastName, update.CallbackQuery.From.FirstName,
							update.CallbackQuery.From.LanguageCode, 0}
						msg.Text = "Ok. I will ask you few questions. This information only for my report, but tell my only the truth\n" +
							"Firstly, I want to know why are you there."
						if !users_info.RegCheck(database, user_info) {
							log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, was registered", user_info.User_nickname, user_info.User_id)}
							log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)
							users_info.RegUser(database, user_info)
							msg.ReplyMarkup = keyboards.Eng_keyboard.First_quest_keyboard
							_, err := my_bot.Send(msg)
							if err != nil{
								log_writer.ErrLogHandler(err.Error())
							}
						} else {
							log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, try to register, but there is user with same id already!", user_info.User_nickname, user_info.User_id)}
							log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)
							msg.Text = "Hej, I am already have you on my list!"
							_, err := my_bot.Send(msg)
							if err != nil{
								log_writer.ErrLogHandler(err.Error())
							}
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "first_quest_str":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{3, 0, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest1", keyboards.Eng_keyboard.Second_quest_keyboard, "What is your race?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "first_quest_agi":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 3, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest1", keyboards.Eng_keyboard.Second_quest_keyboard, "What is your race?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "first_quest_int":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 0, 3}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest1", keyboards.Eng_keyboard.Second_quest_keyboard, "What is your race?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "first_quest_silent":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{1, 1, 1}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest1", keyboards.Eng_keyboard.Second_quest_keyboard, "What is your race?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "second_quest_hum":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{1, 2, 0}
						users_stats.AddRace(database, update.CallbackQuery.From.ID, "Human")
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest2", keyboards.Eng_keyboard.Third_quest_keyboard, "You have a chocolate cake. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}

					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "second_quest_elf":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 2, 1}
						users_stats.AddRace(database, update.CallbackQuery.From.ID, "Elf")
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest2", keyboards.Eng_keyboard.Third_quest_keyboard, "You have a chocolate cake. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "second_quest_dwarf":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{2, 0, 1}
						users_stats.AddRace(database, update.CallbackQuery.From.ID, "Dwarf")
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest2", keyboards.Eng_keyboard.Third_quest_keyboard, "You have a chocolate cake. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "second_quest_orc":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{2, 1, 0}
						users_stats.AddRace(database, update.CallbackQuery.From.ID, "Orc")
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest2", keyboards.Eng_keyboard.Third_quest_keyboard, "You have a chocolate cake. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "second_quest_silent":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{1, 1, 1}
						users_stats.AddRace(database, update.CallbackQuery.From.ID, "Unknown")
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest2", keyboards.Eng_keyboard.Third_quest_keyboard, "You have a chocolate cake. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "third_quest_str":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{3, 0, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest3", keyboards.Eng_keyboard.Forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "third_quest_agi":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 3, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest3", keyboards.Eng_keyboard.Forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "third_quest_int":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 0, 3}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest3", keyboards.Eng_keyboard.Forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "third_quest_silent":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{1, 1, 1}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest3", keyboards.Eng_keyboard.Forth_quest_keyboard, "Your friend has serious injury. What will you do?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "forth_quest_str":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{3, 0, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest4", keyboards.Eng_keyboard.Fifth_quest_keyboard, "What are you looking here?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "forth_quest_agi":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 3, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest4", keyboards.Eng_keyboard.Fifth_quest_keyboard, "What are you looking here?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "forth_quest_int":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 0, 3}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest4", keyboards.Eng_keyboard.Fifth_quest_keyboard, "What are you looking here?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "forth_quest_silent":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{1, 1, 1}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest4", keyboards.Eng_keyboard.Fifth_quest_keyboard, "What are you looking here?"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "fifth_quest_str":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{3, 0, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest5", keyboards.Eng_keyboard.Menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "fifth_quest_agi":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 3, 0}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest5", keyboards.Eng_keyboard.Menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "fifth_quest_int":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{0, 0, 3}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest5", keyboards.Eng_keyboard.Menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "fifth_quest_silent":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := structs.UserCoreStats{1, 1, 1}
						_, err := my_bot.Send(RegQuestion(update.CallbackQuery.Message.Chat.ID, user_stats, database, update.CallbackQuery.From.ID,
							update.CallbackQuery.From.UserName, "quest5", keyboards.Eng_keyboard.Menu_keyboard, "Ok, that's all for now. Do whatever you want and try not to die"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "check_char_stat":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stast := users_stats.TakeFullStats(database, update.CallbackQuery.From.ID)
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						message := fmt.Sprintf("Your Lvl: %v \nExperiense you have: %v \nExperience need to next Lvl: %v \nSkill points you have: %v \nYour energy: %v \nYour race: %v \nYour HP: %v \nYour stamina: %v \nYour MP: %v \n\nAttributes\n\nYour strength: %v \nYour agility: %v \nYour intelligence: %v \nYour armor: %v \nYour magic armor: %v \nYour stun chance: %v%% \nYour crit chance: %v%% \nYour dodge chance: %v%% \nYour magic effect chance: %v%% \nYour meele miss chance: %v%% \nYour range miss chance: %v%% \n\nMagic elements\n\nFire element: %v \nWater element: %v \nEarth element: %v \nWind element: %v \n", user_stast.Lvl, user_stast.Ex_now, user_stast.Ex_next_lvl, user_stast.Skill_point, user_stast.Energy, user_stast.Race, user_stast.Hp, user_stast.Stamina, user_stast.Mp, user_stast.Str, user_stast.Agi, user_stast.Int, user_stast.Armor, user_stast.Magic_armor, user_stast.Stun_chance, user_stast.Crit_chance, user_stast.Dodge_chance, user_stast.Effect_chance, user_stast.Meele_miss_chance, user_stast.Range_miss_chance, user_stast.Fire, user_stast.Water, user_stast.Earth, user_stast.Wind)
						msg.Text = message
						msg.ReplyMarkup = keyboards.Eng_keyboard.Menu_keyboard
						_, err := my_bot.Send(msg)
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "go_adventure":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						msg.Text = "Choose the location"
						log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, started adventure", update.CallbackQuery.From.UserName, update.CallbackQuery.From.ID)}
						log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)
						msg.ReplyMarkup = keyboards.Eng_keyboard.Area_keyboard
						_, err := my_bot.Send(msg)
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "back":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Back to menu")
						msg.ReplyMarkup = keyboards.Eng_keyboard.Menu_keyboard
						_, err := my_bot.Send(msg)
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "1":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						area_act := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						area_act.Text = "You in forest"
						area_act.ReplyMarkup = keyboards.Eng_keyboard.Area_action_keyboard
						buffer_areas.SetArea(database, update.CallbackQuery.From.ID, 1)
						log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, go to forest", update.CallbackQuery.From.UserName, update.CallbackQuery.From.ID)}
						log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)
						_, err := my_bot.Send(area_act)
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "2":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, go to mountains", update.CallbackQuery.From.UserName, update.CallbackQuery.From.ID)}
						log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "enemy":
					if !pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						_, err := my_bot.Send(EnemySearcher(update.CallbackQuery.Message.Chat.ID, database, update.CallbackQuery.From.ID))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are fighting now!"))
						if err != nil{
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "attack":
					if pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := users_stats.TakeFullStats(database, update.CallbackQuery.From.ID)
						if user_stats.Hp >= 10 {
							_, err := my_bot.Send(CalcDamage(update.CallbackQuery.Message.Chat.ID, database, update.CallbackQuery.From.ID))
							if err != nil {
								log_writer.ErrLogHandler(err.Error())
							}
							_, err = my_bot.Send(CalcBotDamage(update.CallbackQuery.Message.Chat.ID, database, update.CallbackQuery.From.ID))
							if err != nil {
								log_writer.ErrLogHandler(err.Error())
							}
						} else {
							_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are injured! Try to rest"))
							if err != nil {
								log_writer.ErrLogHandler(err.Error())
							}
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are not fighting now!"))
						if err != nil {
							log_writer.ErrLogHandler(err.Error())
						}
					}
				case "defence":
					if pve_fight_buffer.CheckBattle(database, update.CallbackQuery.From.ID) {
						user_stats := users_stats.TakeFullStats(database, update.CallbackQuery.From.ID)
						if user_stats.Stamina >= 5 {
							pve_fight_buffer.SetBlock(database, update.CallbackQuery.From.ID, 1)
							log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, set block", update.CallbackQuery.From.ID)}
							log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
							_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You set a block"))
							if err != nil {
								log_writer.ErrLogHandler(err.Error())
							}
							_, err = my_bot.Send(CalcBotDamage(update.CallbackQuery.Message.Chat.ID, database, update.CallbackQuery.From.ID))
							if err != nil {
								log_writer.ErrLogHandler(err.Error())
							}
						} else {
							_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You have not stamina!"))
							if err != nil {
								log_writer.ErrLogHandler(err.Error())
							}
						}
					} else {
						_, err := my_bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You are not fighting now!"))
						if err != nil {
							log_writer.ErrLogHandler(err.Error())
						}
					}
				}
			}
			my_bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID,""))
		}

		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		chat_id := update.Message.Chat.ID
		if !users_info.CheckBan(database, update.Message.From.ID) {
			msg := tgbotapi.NewMessage(chat_id, "")
			user_info := structs.UserInfo{update.Message.From.ID, update.Message.From.UserName, update.Message.From.LastName, update.Message.From.FirstName, update.Message.From.LanguageCode, 0}

			switch update.Message.Command() {
			case "start":
				if !users_info.RegCheck(database, user_info) {
					log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, start bot", user_info.User_nickname, user_info.User_id)}
					log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)
					msg.Text = "Hello, newcomer. This is a dangerous world, full of monsters, bandits, demons and other evil, that want to kill you. If you want to survive - follow my instructions!\n"
					msg.ReplyMarkup = keyboards.Eng_keyboard.Reg_keyboard
				} else {
					log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, start bot, but this user already registered!", user_info.User_nickname, user_info.User_id)}
					log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)
					mes_for_registered := fmt.Sprintf("Hello, %v. What are you doing here? Or you just lost, little chicken?\n", user_info.User_nickname)
					msg.Text = mes_for_registered
				}
			case "menu":
				if users_info.RegCheck(database, user_info) {
					if users_reg_question.CheckAllAnswers(database, update.Message.From.ID) {
						msg.Text = "You are in menu"
						msg.ReplyMarkup = keyboards.Eng_keyboard.Menu_keyboard
					} else {
						msg.Text = "Complete all answers"
					}
				} else {
					msg.Text = "You need to register first!"
				}
			//case "registration":
			//	msg.Text = "Registration"
			//	if !users_info.RegCheck(database, user_info) {
			//		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, was registered", user_info.User_nickname, user_info.User_id)}
			//		log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)
			//		users_info.RegUser(database, user_info)
			//	} else {
			//		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, try to register, but there is user with same id already!", user_info.User_nickname, user_info.User_id)}
			//		log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)
			//		msg.Text = "Already registered"
			//	}
			//case "start_adventure":
			//	msg.Text = "Adventure"
			//	user_info := structs.UserInfo{update.Message.From.ID, update.Message.From.UserName, update.Message.From.LastName, update.Message.From.FirstName, update.Message.From.LanguageCode, 0}
			//	if !users_info.RegCheck(database, user_info) {
			//		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, try to start adventure, but he is not registered jet!", user_info.User_nickname, user_info.User_id)}
			//		log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)
			//		msg.Text = "Please, register first"
			//	} else {
			//		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, started adventure", user_info.User_nickname, user_info.User_id)}
			//		log_writer.LogWrite(log_insert, log_writer.Log_files.Adventure_log)
			//		msg.ReplyMarkup = keyboards.Eng_keyboard.Area_keyboard
			//	}
			default:
				msg.Text = "Bad command"
			}
			_, err := my_bot.Send(msg)
			if err != nil {
				log_writer.ErrLogHandler(err.Error())
			}
		} else {
			_, err := my_bot.Send(tgbotapi.NewMessage(chat_id, "You are in prison"))
			if err != nil {
				log_writer.ErrLogHandler(err.Error())
			}
		}
	}
}
func RegQuestion(chat_string int64, stats structs.UserCoreStats, my_db *sql.DB, user_id int, user_Name string, question string, keys tgbotapi.InlineKeyboardMarkup, next_quest string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chat_string, "")
	if users_reg_question.CheckAnswers(my_db, user_id, question) == 0 {
		users_stats.AddAttrib(my_db, user_id, stats)
		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" User %v, ID is %v, answer on %v [%v, %v, %v]", user_Name, user_id, question, stats.Str, stats.Agi, stats.Int)}
		log_writer.LogWrite(log_insert, log_writer.Log_files.Reg_log)
		msg.Text = next_quest
		msg.ReplyMarkup = keys
		users_reg_question.WriteAnswers(my_db, user_id, question)
		if question == "quest5" {
			users_stats.CalcStatsAfterReg(my_db, user_id)
		}
	} else {
		msg.Text = "You are already answer!"
	}
	return msg
}
func EnemySearcher(chat_string int64, my_db *sql.DB, user_id int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chat_string, "")
	user_stats := users_stats.TakeFullStats(my_db, user_id)
	if user_stats.Energy > 0 {
		enemy_higher_edge := user_stats.Lvl + 2
		rand.Seed(time.Now().UnixNano())
		enemy_Lvl := rand.Intn(enemy_higher_edge - user_stats.Lvl + 1) + user_stats.Lvl
		area_now := buffer_areas.GetArea(my_db, user_id)
		enemy_stats := enemies.GetEnemy(my_db, area_now, enemy_Lvl)
		pve_fight_buffer.AddBattle(my_db, user_id,enemy_stats)
		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" ID is %v, was find enemy %v", user_id, enemy_stats.Name)}
		log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
		msg.Text = fmt.Sprintf("You was attacked by %v \n", enemy_stats.Name) + fmt.Sprintf("Enemy Lvl: %v \nEnemy HP: %v \nEnemy stamina: %v \nEnemy MP: %v \n", enemy_stats.Lvl, enemy_stats.Hp, enemy_stats.Stamina, enemy_stats.Mp)
		msg.ReplyMarkup = keyboards.Eng_keyboard.Fight_keyboard
	} else {
		msg.Text = "You are tired and no have energy to fight."
	}

	return msg
}
func CalcDamage(chat_string int64, my_db *sql.DB, user_id int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chat_string, "")
	user_stats := users_stats.TakeFullStats(my_db, user_id)
	enemy_stats := pve_fight_buffer.GetFightEnemyStats(my_db, user_id)
	user_fight_stats := pve_fight_buffer.GetBattleUserStats(my_db, user_id)
	if user_stats.Hp > 0 {
		if user_fight_stats.Is_stunned == 0 {
			if user_stats.Stamina >= 5 {
				pve_fight_buffer.SetBlock(my_db, user_id, 0)
				rand.Seed(time.Now().UnixNano())
				is_miss := rand.Float32() * 100
				if is_miss > user_stats.Meele_miss_chance {
					rand.Seed(time.Now().UnixNano())
					is_miss := rand.Float32() * 100
					if is_miss > enemy_stats.Dodge_chance {
						user_stats.Stamina = user_stats.Stamina - 5
						dmg := float32(user_stats.Str) + float32(user_stats.Agi) - enemy_stats.Armor
						msg.Text = fmt.Sprintf("%v was attacked by You on %v", enemy_stats.Name, dmg)
						rand.Seed(time.Now().UnixNano())
						is_crit := rand.Float32() * 100
						if is_crit <= user_stats.Crit_chance {
							dmg *= 2
							msg.Text = fmt.Sprintf("Critical! %v was attacked by You on %v", enemy_stats.Name, dmg)
						}
						rand.Seed(time.Now().UnixNano())
						is_stunned := rand.Float32() * 100
						if is_stunned <= user_stats.Stun_chance {
							enemy_stats.Is_stuned = 1
							msg.Text += " You have stunned an enemy!"
						}

						enemy_stats.Hp = enemy_stats.Hp - dmg
						if enemy_stats.Hp <= 0 {
							user_stats.Energy -= 1
							user_stats.Ex_now += (user_stats.Lvl + enemy_stats.Lvl) * 5
							if user_stats.Ex_now >= user_stats.Ex_next_lvl {
								user_stats.Lvl += 1
								user_stats.Skill_point += 1
								user_stats.Ex_next_lvl = user_stats.Lvl * 100
								msg.Text = fmt.Sprintf("%v was defeated! You are reach a Lvl %v", enemy_stats.Name, user_stats.Lvl)
							} else {
								msg.Text = fmt.Sprintf("%v was defeated!", enemy_stats.Name)
							}

							pve_fight_buffer.SetNewUserStats(my_db, user_id, user_stats)
							//pve_fight_buffer.DeleteFight(my_db, user_id)
							msg.ReplyMarkup = keyboards.Eng_keyboard.Area_action_keyboard
							log_insert := structs.LogRequest{time.Now(), fmt.Sprintf(" ID is %v, attack. ", user_id) + msg.Text}
							log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
							return msg
						} else {
							pve_fight_buffer.SetNewUserStats(my_db, user_id, user_stats)
							pve_fight_buffer.SetNewBotEnemyFightStats(my_db, user_id, enemy_stats)
							pve_fight_buffer.SetNewUserFightStats(my_db, user_id, user_fight_stats)
							log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, attack", user_id)}
							log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
							return msg
						}
					} else {
						msg.Text = "Enemy has dodged!"
						log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, enemy dodged", user_id)}
						log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
						return msg
					}
				} else {
					msg.Text = "You have miss!"
					log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, has miss", user_id)}
					log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
					return msg
				}
			} else {
				msg.Text = "You have not stamina!"
				pve_fight_buffer.SetBlock(my_db, user_id, 0)
				log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, no stamina", user_id)}
				log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
				return msg
			}
		} else {
			msg.Text = "You are stunned!"
			log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, stunned", user_id)}
			log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
			return msg
		}
	} else {
		msg.Text = "You are seriously injured!"
		user_stats.Energy -= 1
		pve_fight_buffer.SetNewUserStats(my_db, user_id, user_stats)
		pve_fight_buffer.SetNewBotEnemyFightStats(my_db, user_id, enemy_stats)
		pve_fight_buffer.SetNewUserFightStats(my_db, user_id, user_fight_stats)
		//pve_fight_buffer.DeleteFight(my_db, user_id)
		msg.ReplyMarkup = keyboards.Eng_keyboard.Area_action_keyboard
		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, has serious injury", user_id)}
		log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
		return msg
	}
}
func CalcBotDamage(chat_string int64, my_db *sql.DB, user_id int) tgbotapi.MessageConfig {
	bot_attack_mess := tgbotapi.NewMessage(chat_string, "")
	user_stats := users_stats.TakeFullStats(my_db, user_id)
	enemy_stats := pve_fight_buffer.GetFightEnemyStats(my_db, user_id)
	user_fight_stats := pve_fight_buffer.GetBattleUserStats(my_db, user_id)

	if enemy_stats.Hp > 0 {
		if enemy_stats.Is_stuned == 0 {
			if enemy_stats.Stamina >= 5 {
				rand.Seed(time.Now().UnixNano())
				is_miss := rand.Float32() * 100
				if is_miss > enemy_stats.Meele_miss_chance {
					rand.Seed(time.Now().UnixNano())
					is_miss := rand.Float32() * 100
					if is_miss > user_stats.Dodge_chance {
						enemy_stats.Stamina = enemy_stats.Stamina - 5
						dmg := float32(enemy_stats.Str) + float32(enemy_stats.Agi) - user_stats.Armor
						if user_fight_stats.Is_block == 1 {
							dmg = float32(enemy_stats.Str) + float32(enemy_stats.Agi) - (user_stats.Armor * 2)
						}
						bot_attack_mess.Text = fmt.Sprintf("You was attacked by %v on %v", enemy_stats.Name, dmg)
						rand.Seed(time.Now().UnixNano())
						is_crit := rand.Float32() * 100
						if is_crit <= user_stats.Crit_chance {
							dmg *= 2
							bot_attack_mess.Text = fmt.Sprintf("Critical! You was attacked by %v on %v", enemy_stats.Name, dmg)
						}
						rand.Seed(time.Now().UnixNano())
						Is_stunned := rand.Float32() * 100
						if Is_stunned <= user_stats.Stun_chance {
							user_fight_stats.Is_stunned = 1
							bot_attack_mess.Text += " You was stunned by enemy!"
						}

						user_stats.Hp = user_stats.Hp - dmg
						if user_stats.Hp <= 0 {
							bot_attack_mess.Text = fmt.Sprintf("You was defeated by %v", enemy_stats.Name)
							user_stats.Energy -= 1
							pve_fight_buffer.SetNewUserStats(my_db, user_id, user_stats)
							pve_fight_buffer.SetNewUserFightStats(my_db, user_id, user_fight_stats)
							pve_fight_buffer.DeleteFight(my_db, user_id)
							bot_attack_mess.ReplyMarkup = keyboards.Eng_keyboard.Area_action_keyboard
							log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, was defeated by %v", user_id, enemy_stats.Name)}
							log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
							return bot_attack_mess
						} else {
							pve_fight_buffer.SetNewUserStats(my_db, user_id, user_stats)
							pve_fight_buffer.SetNewBotEnemyFightStats(my_db, user_id, enemy_stats)
							pve_fight_buffer.SetNewUserFightStats(my_db, user_id, user_fight_stats)
						}
						bot_attack_mess.ReplyMarkup = keyboards.Eng_keyboard.Fight_keyboard
						log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v was attacked by %v", user_id, enemy_stats.Name)}
						log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
						return bot_attack_mess
					} else {
						bot_attack_mess.Text = "You have dodged!"
						bot_attack_mess.ReplyMarkup = keyboards.Eng_keyboard.Fight_keyboard
						log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, dodged", user_id)}
						log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
						return bot_attack_mess
					}
				} else {
					bot_attack_mess.Text = fmt.Sprintf("%v has miss!", enemy_stats.Name)
					bot_attack_mess.ReplyMarkup = keyboards.Eng_keyboard.Fight_keyboard
					log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, %v miss", user_id, enemy_stats.Name)}
					log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
					return bot_attack_mess
				}
			} else {
				bot_attack_mess.Text = fmt.Sprintf("%v has not stamina!", enemy_stats.Name)
				bot_attack_mess.ReplyMarkup = keyboards.Eng_keyboard.Fight_keyboard
				log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, %v without stamina", user_id, enemy_stats.Name)}
				log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
				return bot_attack_mess
			}
		} else {
			bot_attack_mess.Text = fmt.Sprintf("%v is stunned!", enemy_stats.Name)
			bot_attack_mess.ReplyMarkup = keyboards.Eng_keyboard.Fight_keyboard
			log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, %v stunned", user_id, enemy_stats.Name)}
			log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
			return bot_attack_mess
		}
	} else {
		bot_attack_mess.Text = fmt.Sprintf("%v is dead!", enemy_stats.Name)
		pve_fight_buffer.DeleteFight(my_db, user_id)
		bot_attack_mess.ReplyMarkup = keyboards.Eng_keyboard.Area_action_keyboard
		log_insert := structs.LogRequest{time.Now(), fmt.Sprintf("ID is %v, %v dead", user_id, enemy_stats.Name)}
		log_writer.LogWrite(log_insert, log_writer.Log_files.Battle_log)
		return bot_attack_mess
	}
}

func main() {

	bot := BotStart()
	my_db := db.DBStart()
	defer my_db.Close()

	BotUpdateLoop(bot, my_db)

}
