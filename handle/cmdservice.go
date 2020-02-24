package handle

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//League represents a fantasy league
type League struct {
	Name string
	URL  string
}

var (
	c = NewCache(30)
)

//GetList returns list of participants as text, including links (sports.ru) to profiles/teams
func GetList(league League) (string, error) {
	if val, hasValue := c.Get("GetList" + league.URL); hasValue {
		return val, nil
	}
	participants, err := GetParticipants(league.URL)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var res string
	for _, p := range participants {
		res += fmt.Sprintf(`%v. <a href="%v">%v</a> | <a href="%v">%v</a> | %v | %v`, p.Pos, p.TeamURL, p.Team, p.ProfileURL, p.Name, p.TourPoints, p.TotalPoints)
		res += "\n"
	}

	c.Add("GetList"+league.URL, res)

	return res, err
}

//GetTableImg takes a snapshot (using phantomjs) of a league table, returns a file name of it to be retrieved as a static file afterwards
func GetTableImg(league League) (string, error) {
	if val, hasValue := c.Get("GetTableImg" + league.URL); hasValue {
		return val, nil
	}
	f := "table_" + league.Name + ".png"
	cmd := exec.Command("phantomjs", "snapshot.js", league.URL, "table.stat-table", f)
	cmd.Run()

	if _, err := os.Stat(f); err != nil {
		return "", err
	}

	c.Add("GetTableImg"+league.URL, f)
	return f, nil
}

//GetSquad takes a snapshot (using phantomjs) of a squad, returns a file name of it to be retrieved as a static file afterwards
func GetSquad(league League, team string) (string, error) {
	if val, hasValue := c.Get("GetSquad" + league.URL + team); hasValue {
		return val, nil
	}
	participants, err := GetParticipants(league.URL)
	var p *Participant
	for _, pp := range participants {
		if strings.Trim(pp.Team, " ") == team {
			p = &pp
			break
		}
	}
	if p == nil {
		return "", fmt.Errorf("No such team in the league: %v", team)
	}

	f := league.Name + "_" + p.Team + ".png"
	cmd := exec.Command("phantomjs", "snapshot.js", p.TeamURL, "div.grace.full-field", f)
	cmd.Run()
	if _, err := os.Stat(f); err != nil {
		return "", err
	}

	c.Add("GetSquad"+league.URL+team, f)

	return f, err
}
