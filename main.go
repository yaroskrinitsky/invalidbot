package main

import (
	"flag"
	"fmt"

	"github.com/yaroskrinitsky/invalidbot/config"
	"github.com/yaroskrinitsky/invalidbot/handle"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var (
	//cfg ...
	cfg config.Config
)

func main() {
	env := flag.String("env", "dev", "Current environment")
	flag.Parse()
	var err error
	if cfg, err = config.Builders[*env](); err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	var leagues []handle.League
	for i, l := range cfg.Leagues {
		leagues = append(leagues, handle.League{ID: fmt.Sprint(i), Name: l.Name, URL: l.URL})
	}

	handle.NewUpdateListener(bot, leagues).ListenAndServe()
}

