package handle

import (
	"encoding/json"
	"fmt"
	"html"
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

//Start sends the initial bot page, which consists of buttons to select a league
// [/start]
func Start(bot *tgbotapi.BotAPI, u *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Лиги:")
	msg.ReplyMarkup = ComposeInitialMenu()
	msg.ParseMode = "HTML"
	bot.Send(msg)
}

//Stats sends a league stats page and keyboard to select a player
// [/stats]
func Stats(bot *tgbotapi.BotAPI, u *tgbotapi.Update) {
	var cd struct {
		Cmd string `json:"cmd"`
		ID  string `json:"id"`
	}

	json.Unmarshal([]byte(html.UnescapeString(u.CallbackQuery.Data)), &cd)

	var league League
	for _, l := range leagues {
		log.Println(l)
		if l.ID == cd.ID {
			league = l
			break
		}
	}

	msgText, err := GetList(league.URL)
	if err != nil {
		log.Println(u, cd)
		return
	}
	msgText = league.Name + "\n" + msgText

	msg := tgbotapi.NewEditMessageText(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID, msgText)
	//bot.Send(msg1)
	msg.ParseMode = "HTML"
	msg.DisableWebPagePreview = true
	ps, err := GetParticipants(league.URL)
	if err != nil {
		log.Println(u, cd)
	}
	kb := ComposeParticipantsListMenu(ps, league.ID)
	msg.ReplyMarkup = &kb

	bot.Send(msg)

}

//Squad sends a participant squad page and buttons to be able to select another participant
// [/squad]
func Squad(bot *tgbotapi.BotAPI, u *tgbotapi.Update) {
	var cd struct {
		Cmd      string `json:"cmd"`
		Team     string `json:"t"`
		LeagueID string `json:"l"`
	}

	json.Unmarshal([]byte(html.UnescapeString(u.CallbackQuery.Data)), &cd)
	var league League
	for _, l := range leagues {
		if l.ID == cd.LeagueID {
			league = l
			break
		}
	}
	log.Println(cd)
	participants, err := GetParticipants(league.URL)
	if err != nil {
		log.Println(err)
	}
	var p Participant
	for _, pp := range participants {
		if cd.Team == pp.Pos {
			p = pp
			break
		}
	}

	file, err := GetSquad(league.Name, league.URL, p.Team)
	if err != nil {
		log.Println(err)
	}
	photo := tgbotapi.NewPhotoUpload(u.CallbackQuery.Message.Chat.ID, file)
	photo.Caption = fmt.Sprintf("%v | %v | %v | %v | %v", p.Pos, p.Name, p.Team, p.TourPoints, p.TotalPoints)
	bot.Send(photo)
}
