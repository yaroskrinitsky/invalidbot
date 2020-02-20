package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//Commands type to call bot commands
type Commands struct {
	Cache CommandsCache
}

//GetList returns list of participants as text, including links (sports.ru) to profiles/teams
func (c *Commands) GetList(leagueURL string) (string, error) {
	if val, hasValue := c.Cache.Get("GetList" + leagueURL); hasValue {
		return val, nil
	}
	participants, err := GetParticipants(leagueURL)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var res string
	for _, p := range participants {
		res += fmt.Sprintf(`%v. <a href="%v">%v</a> | <a href="%v">%v</a> | %v | %v`, p.Pos, p.TeamURL, p.Team, p.ProfileURL, p.Name, p.TourPoints, p.TotalPoints)
		res += "\n"
	}

	c.Cache.Add("GetList"+leagueURL, res)

	return res, err
}

//GetTableImg takes a snapshot (using phantomjs) of a league table, returns a file name of it to be retrieved as a static file afterwards
func (c *Commands) GetTableImg(leagueURL string) (string, error) {
	if val, hasValue := c.Cache.Get("GetTableImg" + leagueURL); hasValue {
		return val, nil
	}
	f := "table.png"
	cmd := exec.Command("phantomjs", "snapshot.js", leagueURL, "table.stat-table", f)
	cmd.Run()

	if _, err := os.Stat(f); err != nil {
		return "", err
	}

	c.Cache.Add("GetTableImg"+leagueURL, f)
	return f, nil
}

//GetSquad takes a snapshot (using phantomjs) of a squad, returns a file name of it to be retrieved as a static file afterwards
func (c *Commands) GetSquad(leagueURL string, team string) (string, error) {
	if val, hasValue := c.Cache.Get("GetSquad" + leagueURL + team); hasValue {
		return val, nil
	}
	participants, err := GetParticipants(leagueURL)
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

	f := p.Team + ".png"
	cmd := exec.Command("phantomjs", "snapshot.js", p.TeamURL, "div.grace.full-field", f)
	cmd.Run()
	if _, err := os.Stat(f); err != nil {
		return "", err
	}

	c.Cache.Add("GetSquad"+leagueURL+team, f)

	return f, err
}
