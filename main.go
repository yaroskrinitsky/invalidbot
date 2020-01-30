package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

var (
	Env string
	Cfg Config
)

func init() {
	if Env == "" {
		Env = "DEV"
	}
	var err error
	if Cfg, err = CfgBuilders[Env](); err != nil {
		panic(err)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(Cfg.BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updConf := tgbotapi.NewUpdate(0)
	updConf.Timeout = 45

	updates, err := bot.GetUpdatesChan(updConf)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {

			switch update.Message.Command() {
			case "list":
				text, err := GetList(Cfg.LeagueURL)
				if err != nil {
					log.Println(err)
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				msg.ParseMode = "HTML"
				msg.DisableWebPagePreview = true
				bot.Send(msg)
			case "table":
				img, err := GetTableImg(Cfg.LeagueURL)
				if err != nil {
					log.Println(err)
					continue
				}
				photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, img)
				bot.Send(photo)
			case "squad":
				team := update.Message.CommandArguments()
				img, err := GetSquad(Cfg.LeagueURL, team)
				if err != nil {
					log.Println(err)
					continue
				}
				photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, img)

				bot.Send(photo)
			}
		}
	}
}