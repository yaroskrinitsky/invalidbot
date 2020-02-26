package handle

import (
	"fmt"
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var (
	leagues = []League{
		League{ID: "1", Name: "Invalid cup 2019", URL: "https://www.sports.ru/fantasy/football/league/143767.html"},
		League{ID: "2", Name: "Invalid Champions League Qualification", URL: "https://www.sports.ru/fantasy/football/league/149896.html"},
	}
)

//ComposeInitialMenu returns the initial inline keyboard markup containing buttons to choose a league
func ComposeInitialMenu() tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	for _, l := range leagues {
		callbackData := fmt.Sprintf(`{"cmd":"stats","id":"%v"}`, l.ID)

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(l.Name, callbackData))
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)

	return keyboard
}

//ComposeParticipantsListMenu returns inline keyboard mark up to select a participant
func ComposeParticipantsListMenu(ps []Participant, leagueID string) tgbotapi.InlineKeyboardMarkup {
	var buttons [][]tgbotapi.InlineKeyboardButton
	for i, p := range ps {
		if i%8 == 0 {
			buttons = append(buttons, []tgbotapi.InlineKeyboardButton{})
		}
		callbackData := fmt.Sprintf(`{"cmd":"squad","t":"%v","l":"%v"}`, p.Pos, leagueID)
		log.Println(callbackData)
		buttons[len(buttons)-1] = append(buttons[len(buttons)-1], tgbotapi.NewInlineKeyboardButtonData(p.Pos, callbackData))
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		buttons...,
	)

	return keyboard
}
