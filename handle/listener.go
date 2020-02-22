package handle

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

//UpdateHandler ...
type UpdateHandler func(u *tgbotapi.Update)

//UpdateListener ...
type UpdateListener struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]UpdateHandler
}

//NewUpdateListener ...
func NewUpdateListener(bot *tgbotapi.BotAPI) *UpdateListener {
	u := &UpdateListener{bot: bot}
	u.handlers = make(map[string]UpdateHandler)
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
		if update.Message == nil {
			continue
		}
	}
}
