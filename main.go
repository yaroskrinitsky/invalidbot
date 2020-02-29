package main

import (
	"github.com/yaroskrinitsky/invalidbot/config"
	"github.com/yaroskrinitsky/invalidbot/handle"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var (
	//Env ...
	Env string
	//Cfg ...
	Cfg config.Config
)

func init() {
	if Env == "" {
		Env = "DEV"
	}
	var err error
	if Cfg, err = config.Builders[Env](); err != nil {
		panic(err)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(Cfg.BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	handle.NewUpdateListener(bot).ListenAndServe()

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	// updConf := tgbotapi.NewUpdate(0)
	// updConf.Timeout = 45

	//updates, err := bot.GetUpdatesChan(updConf)

	// for update := range updates {
	// 	if update.Message == nil {
	// 		continue
	// 	}

	// 	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 	if update.Message.IsCommand() {

	// 		switch update.Message.Command() {
	// 		case "list":
	// 			text, err := cmd.GetList(Cfg.LeagueURL)
	// 			if err != nil {
	// 				log.Println(err)
	// 				continue
	// 			}
	// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	// 			//tgbotapi.NewCallback()
	// 			msg.ParseMode = "HTML"
	// 			msg.DisableWebPagePreview = true
	// 			bot.Send(msg)
	// 		case "table":
	// 			img, err := cmd.GetTableImg(Cfg.LeagueURL)
	// 			if err != nil {
	// 				log.Println(err)
	// 				continue
	// 			}
	// 			photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, img)
	// 			bot.Send(photo)
	// 		case "squad":
	// 			team := update.Message.CommandArguments()
	// 			img, err := cmd.GetSquad(Cfg.LeagueURL, team)
	// 			if err != nil {
	// 				log.Println(err)
	// 				continue
	// 			}
	// 			photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, img)

	// 			bot.Send(photo)
	// 		}
	// 	}
	// }
}

// package main

// import (
// 	"fmt"
// 	"log"

// 	tgbotapi "gopkg.in/telegram-bot-api.v4"
// )

// var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
// 		tgbotapi.NewInlineKeyboardButtonSwitch("2sw", "open 2"),
// 		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
// 	),
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
// 		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
// 		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
// 	),
// )

// func main() {
// 	bot, err := tgbotapi.NewBotAPI("1023081918:AAHl1K2Tr88m5hj_ccTmrPZXg-ypL1_jRpQ")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates, err := bot.GetUpdatesChan(u)

// 	fmt.Print(".")
// 	for update := range updates {
// 		if update.CallbackQuery != nil {
// 			//fmt.Print(update)
// 			fmt.Println(update.CallbackQuery)
// 			//bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))

// 			//bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
// 			msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, "EDITED MESSAGE "+update.CallbackQuery.Data+" "+update.CallbackQuery.From.String())
// 			bot.Send(msg)
// 		}
// 		if update.Message != nil {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

// 			switch update.Message.Text {
// 			case "open":
// 				msg.ReplyMarkup = numericKeyboard
// 			case "close":
// 				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

// 			}
// 			msg.Text = "Hi, this is reply markup, press me!"
// 			bot.Send(msg)
// 		}
// 	}
// }
