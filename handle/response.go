package handle

import (
	"fmt"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var (
	leagues = []League{
		League{Name: "Invalid cup 2019", URL: "https://www.sports.ru/fantasy/football/league/143767.html"},
		League{Name: "Invalid Champions League Qualification", URL: "https://www.sports.ru/fantasy/football/league/149896.html"},
	}
)

//ComposeInitialMenu returns the initial inline keyboard markup containing buttons to choose a league
func ComposeInitialMenu() tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	for _, l := range leagues {
		// 'league' -- command
		callbackData := fmt.Sprintf(`
		{
			"command": "league",
			"name": "%v",
			"url": "%v"
		}`, l.Name, l.URL)

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(l.Name, callbackData))
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)

	return keyboard
}
