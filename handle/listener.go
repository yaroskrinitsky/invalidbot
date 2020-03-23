package handle

import (
	"encoding/json"
	"log"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

//UpdateHandler ...
type UpdateHandler func(bot *tgbotapi.BotAPI, u *tgbotapi.Update)

//UpdateListener ...
type UpdateListener struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]UpdateHandler
}

//NewUpdateListener ...
func NewUpdateListener(bot *tgbotapi.BotAPI) *UpdateListener {
	u := &UpdateListener{bot: bot}
	u.handlers = make(map[string]UpdateHandler)
	u.Handle("start", Start)
	u.Handle("stats", Stats)
	u.Handle("squad", Squad)
	return u
}

//Handle ...
func (ul *UpdateListener) Handle(cmd string, handler UpdateHandler) {
	ul.handlers[cmd] = handler
}

//ListenAndServe ...
func (ul *UpdateListener) ListenAndServe() {
	updConf := tgbotapi.NewUpdate(0)
	updConf.Timeout = 45

	updates, err := ul.bot.GetUpdatesChan(updConf)

	if err != nil {
		panic(err)
	}

	for update := range updates {
		ul.Serve(&update)
	}
}

//Serve takes an update and calls a corresponding handler for it
func (ul *UpdateListener) Serve(u *tgbotapi.Update) {
	var c string
	if u.CallbackQuery != nil {
		var callbackData map[string]string
		json.Unmarshal([]byte(u.CallbackQuery.Data), &callbackData)
		c = callbackData["cmd"]
	} else if u.Message != nil {
		c = strings.Split(u.Message.Text, " ")[0]
	}

	if h := ul.handlers[c]; h != nil {
		h(ul.bot, u)
	}
}
